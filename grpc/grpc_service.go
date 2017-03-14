package grpc

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	_ "github.com/mosuka/indigo/dependency"
	"github.com/mosuka/indigo/proto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

type indigoGRPCService struct {
	dataDir string
	indices map[string]bleve.Index
	mutexes map[string]*sync.RWMutex
}

func NewIndigoGRPCService(dataDir string) *indigoGRPCService {
	indices := make(map[string]bleve.Index)
	mutexes := make(map[string]*sync.RWMutex)

	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err == nil {
			log.WithFields(log.Fields{
				"dataDir": dataDir,
			}).Debug("succeeded in creating data directory")
		} else {
			log.WithFields(log.Fields{
				"dataDir": dataDir,
				"err":     err,
			}).Error("failed to create data directory")
		}
	} else {
		log.WithFields(log.Fields{
			"dataDir": dataDir,
		}).Warn("data directory already exists")
	}

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
		mutexes: mutexes,
	}
}

func (igs *indigoGRPCService) lockIndex(index string) {
	if _, existed := igs.mutexes[index]; !existed {
		igs.mutexes[index] = new(sync.RWMutex)
	}

	igs.mutexes[index].Lock()

	log.WithFields(log.Fields{
		"index": index,
	}).Debug("index was locked")
}

func (igs *indigoGRPCService) unlockIndex(index string) {
	if _, existed := igs.mutexes[index]; !existed {
		igs.mutexes[index] = new(sync.RWMutex)
	}

	igs.mutexes[index].Unlock()

	log.WithFields(log.Fields{
		"index": index,
	}).Debug("index was unlocked")
}

func (igs *indigoGRPCService) OpenIndices() {
	if fiList, err := ioutil.ReadDir(igs.dataDir); err == nil {
		for _, fi := range fiList {
			if fi.IsDir() {
				index := fi.Name()
				indexDir := path.Join(igs.dataDir, index)
				idx, err := bleve.Open(indexDir)
				if err == nil {
					log.WithFields(log.Fields{
						"index":    index,
						"indexDir": indexDir,
					}).Info("succeeded in opening index")

					igs.indices[index] = idx
				} else {
					log.WithFields(log.Fields{
						"index":    index,
						"indexDir": indexDir,
						"err":      err,
					}).Warn("failed to open index")
				}
			}
		}
	} else {
		log.WithFields(log.Fields{
			"dataDir": igs.dataDir,
			"err":     err,
		}).Warn("failed to open data directory")
	}

	return
}

func (igs *indigoGRPCService) CloseIndices() {
	for index, idx := range igs.indices {
		if err := idx.Close(); err == nil {
			log.WithFields(log.Fields{
				"index": index,
			}).Info("succeeded in closing index")
		} else {
			log.WithFields(log.Fields{
				"index": index,
				"err":   err,
			}).Warn("failed to close index")
		}
	}

	return
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, req *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to create index")

		return &proto.CreateIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	indexMapping := bleve.NewIndexMapping()
	if req.IndexMapping != nil {
		if err := json.Unmarshal(req.IndexMapping, &indexMapping); err == nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
			}).Debug("succeeded in creating index mapping")
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to creat index mapping")

			return &proto.CreateIndexResponse{}, err
		}
	}

	kvConfig := make(map[string]interface{})
	if req.Kvconfig != nil {
		if err := json.Unmarshal(req.Kvconfig, &kvConfig); err == nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
			}).Debug("succeeded in creating kv config")
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to create kv config")

			return &proto.CreateIndexResponse{}, err
		}
	}

	_, err := os.Stat(indexDir)
	if os.IsNotExist(err) {
		var index bleve.Index = nil
		index, err = bleve.NewUsing(indexDir, indexMapping, req.IndexType, req.Kvstore, kvConfig)
		if err == nil {
			log.WithFields(log.Fields{
				"index":     req.IndexName,
				"indexDir":  indexDir,
				"indexType": req.IndexType,
				"kvStore":   req.Kvstore,
			}).Info("succeeded in creating index")

			igs.indices[req.IndexName] = index
		} else {
			log.WithFields(log.Fields{
				"index":     req.IndexName,
				"indexDir":  indexDir,
				"indexType": req.IndexType,
				"kvStore":   req.Kvstore,
				"err":       err,
			}).Error("failed to create index")
		}
	} else {
		log.WithFields(log.Fields{
			"index":     req.IndexName,
			"indexDir":  indexDir,
			"indexType": req.IndexType,
			"kvStore":   req.Kvstore,
			"err":       err,
		}).Error("failed to create index")
	}

	return &proto.CreateIndexResponse{
		IndexName: req.IndexName,
		IndexDir:  indexDir,
	}, err
}

func (igs *indigoGRPCService) OpenIndex(ctx context.Context, req *proto.OpenIndexRequest) (*proto.OpenIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to open index")

		return &proto.OpenIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	runtimeConfig := make(map[string]interface{})
	if req.RuntimeConfig != nil {
		err := json.Unmarshal(req.RuntimeConfig, &runtimeConfig)
		if err == nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
			}).Debug("succeeded in creating runtime config")
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to create runtime config")

			return &proto.OpenIndexResponse{}, err
		}
	}

	_, err := os.Stat(indexDir)
	if err == nil {
		var index bleve.Index = nil
		index, err = bleve.OpenUsing(indexDir, runtimeConfig)
		if err == nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
			}).Info("succeeded in opening index")

			igs.indices[req.IndexName] = index
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to open index")
		}
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to open index")
	}

	return &proto.OpenIndexResponse{
		IndexName: req.IndexName,
		IndexDir:  indexDir,
	}, err
}

func (igs *indigoGRPCService) GetIndex(ctx context.Context, req *proto.GetIndexRequest) (*proto.GetIndexResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to get index")

		return &proto.GetIndexResponse{}, err
	}

	documentCount, err := index.DocCount()

	indexStats, err := index.Stats().MarshalJSON()
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
		}).Info("succeeded in creating index stats")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to create index stats")
	}

	indexMapping, err := json.Marshal(index.Mapping())
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
		}).Info("succeeded in creating index mapping")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to create index mapping")
	}

	return &proto.GetIndexResponse{
		DocumentCount: documentCount,
		IndexStats:    indexStats,
		IndexMapping:  indexMapping,
	}, err
}

func (igs *indigoGRPCService) CloseIndex(ctx context.Context, req *proto.CloseIndexRequest) (*proto.CloseIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to close index")

		return &proto.CloseIndexResponse{}, err
	}

	err := index.Close()
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
		}).Info("succeeded in closing index")

		delete(igs.indices, req.IndexName)
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to close index")
	}

	return &proto.CloseIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to delete index")

		return &proto.DeleteIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	_, err := os.Stat(indexDir)
	if err == nil {
		err = os.RemoveAll(indexDir)
		if err == nil {
			log.WithFields(log.Fields{
				"index":    req.IndexName,
				"indexDir": indexDir,
			}).Info("succeeded in deleting index")
		} else {
			log.WithFields(log.Fields{
				"index":    req.IndexName,
				"indexDir": indexDir,
				"err":      err,
			}).Error("failed to delete index")
		}
	} else {
		log.WithFields(log.Fields{
			"index":    req.IndexName,
			"indexDir": indexDir,
			"err":      err,
		}).Error("failed to delete index")
	}

	return &proto.DeleteIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) ListIndex(ctx context.Context, req *proto.ListIndexRequest) (*proto.ListIndexResponse, error) {
	indices := make([]string, 0)

	for indexName := range igs.indices {
		indices = append(indices, indexName)
	}

	return &proto.ListIndexResponse{
		Indices: indices,
	}, nil
}

func (igs *indigoGRPCService) PutDocument(ctx context.Context, req *proto.PutDocumentRequest) (*proto.PutDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to put document")

		return &proto.PutDocumentResponse{}, err
	}

	success := false
	var fields interface{}
	err := json.Unmarshal(req.Fields, &fields)
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
		}).Debug("succeeded in creating document")

		err = index.Index(req.Id, fields)
		if err == nil {
			success = true

			log.WithFields(log.Fields{
				"index": req.IndexName,
				"id":    req.Id,
			}).Info("succeeded in putting document")
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"id":    req.Id,
				"err":   err,
			}).Error("failed to put document")
		}
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
			"err":   err,
		}).Error("failed to put document")
	}

	return &proto.PutDocumentResponse{
		Success: success,
	}, err
}

func (igs *indigoGRPCService) GetDocument(ctx context.Context, req *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to get document")

		return &proto.GetDocumentResponse{}, err
	}

	fields := make(map[string]interface{})
	if doc, err := index.Document(req.Id); err == nil {
		if doc != nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"id":    req.Id,
			}).Info("succeeded in getting document")

			for _, field := range doc.Fields {
				var value interface{}

				switch field := field.(type) {
				case *document.TextField:
					value = string(field.Value())
				case *document.NumericField:
					numValue, err := field.Number()
					if err == nil {
						value = numValue
					}
				case *document.DateTimeField:
					dateValue, err := field.DateTime()
					if err == nil {
						dateValue.Format(time.RFC3339Nano)
						value = dateValue
					}
				}

				existedField, existed := fields[field.Name()]
				if existed {
					switch existedField := existedField.(type) {
					case []interface{}:
						fields[field.Name()] = append(existedField, value)
					case interface{}:
						arr := make([]interface{}, 2)
						arr[0] = existedField
						arr[1] = value
						fields[field.Name()] = arr
					}
				} else {
					fields[field.Name()] = value
				}
			}
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"id":    req.Id,
			}).Info("document does not exist")
		}
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
			"err":   err,
		}).Error("failed to get document")

		return &proto.GetDocumentResponse{}, err
	}

	bytesFields, err := json.Marshal(fields)
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
		}).Debug("succeeded in creating document")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
			"err":   err,
		}).Error("failed to get document")
	}

	return &proto.GetDocumentResponse{
		Id:     req.Id,
		Fields: bytesFields,
	}, err
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, req *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to delete document")

		return &proto.DeleteDocumentResponse{}, err
	}

	success := false
	err := index.Delete(req.Id)
	if err == nil {
		success = true
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
		}).Info("succeeded in deleting document")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"id":    req.Id,
			"err":   err,
		}).Error("failed to delete document")
	}

	return &proto.DeleteDocumentResponse{
		Success: success,
	}, err
}

func (igs *indigoGRPCService) Bulk(ctx context.Context, req *proto.BulkRequest) (*proto.BulkResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to index documents in bulk")

		return &proto.BulkResponse{}, err
	}

	var bulkRequest interface{}
	if req.BulkRequest != nil {
		err := json.Unmarshal(req.BulkRequest, &bulkRequest)
		if err != nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to index documents in bulk")

			return &proto.BulkResponse{}, err
		}
	}

	_, ok := bulkRequest.([]interface{})
	if !ok {
		err := errors.New("unexpected bulk request format")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to index documents in bulk")

		return &proto.BulkResponse{}, err
	}

	var (
		batchCount    int32
		putCount      int32
		putErrorCount int32
		deleteCount   int32
	)

	batch := index.NewBatch()

	for num, request := range bulkRequest.([]interface{}) {
		request, ok := request.(map[string]interface{})
		if !ok {
			log.WithFields(log.Fields{
				"index":   req.IndexName,
				"num":     num,
				"request": request,
			}).Warn("unexpected request format")

			continue
		}

		var method string
		if _, ok := request["method"]; !ok {
			log.WithFields(log.Fields{
				"index":   req.IndexName,
				"num":     num,
				"request": request,
			}).Warn("method does not exist in request")

			continue
		}
		method = request["method"].(string)

		var id string
		if _, ok := request["id"]; !ok {
			log.WithFields(log.Fields{
				"index":   req.IndexName,
				"num":     num,
				"request": request,
			}).Warn("id does not exist in request")

			continue
		}
		id = request["id"].(string)

		switch method {
		case "put":
			var fields interface{}
			if _, ok := request["fields"]; !ok {
				log.WithFields(log.Fields{
					"index":   req.IndexName,
					"num":     num,
					"request": request,
				}).Warn("fields does not exist in request")

				continue
			}
			fields = request["fields"]

			err := batch.Index(id, fields)
			if err == nil {
				log.WithFields(log.Fields{
					"index":   req.IndexName,
					"num":     num,
					"request": request,
				}).Info("succeeded in putting document")

				putCount++
				batchCount++
			} else {
				log.WithFields(log.Fields{
					"index":   req.IndexName,
					"num":     num,
					"request": request,
					"err":     err,
				}).Warn("failed to put document")

				putErrorCount++
			}
		case "delete":
			batch.Delete(id)

			log.WithFields(log.Fields{
				"index":   req.IndexName,
				"num":     num,
				"request": request,
			}).Info("succeeded in deleting document")

			deleteCount++
			batchCount++
		default:
			log.WithFields(log.Fields{
				"index":   req.IndexName,
				"num":     num,
				"request": request,
			}).Warn("unexpected method")

			continue
		}

		if batchCount%req.BatchSize == 0 {
			err := index.Batch(batch)
			if err == nil {
				log.WithFields(log.Fields{
					"index": req.IndexName,
					"count": batch.Size(),
				}).Info("succeeded in indexing documents in bulk")
			} else {
				log.WithFields(log.Fields{
					"index": req.IndexName,
					"count": batch.Size(),
				}).Warn("failed to index  documents in bulk")
			}

			batch = index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		err := index.Batch(batch)
		if err == nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"count": batch.Size(),
			}).Info("succeeded in indexing documents in bulk")
		} else {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"count": batch.Size(),
			}).Warn("failed to index  documents in bulk")
		}
	}

	return &proto.BulkResponse{
		PutCount:      putCount,
		PutErrorCount: putErrorCount,
		DeleteCount:   deleteCount,
	}, nil
}

func (igs *indigoGRPCService) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")

		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to search documents")

		return &proto.SearchResponse{}, err
	}

	searchRequest := bleve.NewSearchRequest(nil)
	if req.SearchRequest != nil {
		err := json.Unmarshal(req.SearchRequest, searchRequest)
		if err != nil {
			log.WithFields(log.Fields{
				"index": req.IndexName,
				"err":   err,
			}).Error("failed to search documents")

			return &proto.SearchResponse{}, err
		}
	}

	searchResult, err := index.Search(searchRequest)
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
		}).Info("succeeded in searching documents")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to search documents")

		return &proto.SearchResponse{}, err
	}

	bytesSearchResult, err := json.Marshal(&searchResult)
	if err == nil {
		log.WithFields(log.Fields{
			"index": req.IndexName,
		}).Debug("succeeded in creating search result")
	} else {
		log.WithFields(log.Fields{
			"index": req.IndexName,
			"err":   err,
		}).Error("failed to create search result")
	}
	return &proto.SearchResponse{
		SearchResult: bytesSearchResult,
	}, err
}
