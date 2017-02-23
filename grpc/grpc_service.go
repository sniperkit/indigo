package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	_ "github.com/mosuka/indigo/config"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
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
	var err error = nil

	indices := make(map[string]bleve.Index)
	mutexes := make(map[string]*sync.RWMutex)

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err == nil {
			log.Printf("debug: make data directory dataDir=\"%s\"\n", dataDir)
		} else {
			log.Printf("error: failed to make data directory dataDir=\"%s\" error=\"%s\"\n", dataDir, err.Error())
		}
	} else {
		err = errors.New("directory exists")
		log.Printf("debug: data directory already exists dataDir=\"%s\" error=\"%s\"\n", dataDir, err.Error())
	}

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
		mutexes: mutexes,
	}
}

func (igs *indigoGRPCService) lockIndex(indexName string) {
	_, mutexExisted := igs.mutexes[indexName]
	if mutexExisted == false {
		igs.mutexes[indexName] = new(sync.RWMutex)
	}

	igs.mutexes[indexName].Lock()
	log.Printf("debug: lock index indexName=\"%s\"\n", indexName)
}

func (igs *indigoGRPCService) unlockIndex(indexName string) {
	_, mutexExisted := igs.mutexes[indexName]
	if mutexExisted == false {
		igs.mutexes[indexName] = new(sync.RWMutex)
	}

	igs.mutexes[indexName].Unlock()
	log.Printf("debug: unlock index indexName=\"%s\"\n", indexName)
}

func (igs *indigoGRPCService) OpenIndices() error {
	var err error = nil

	fiList, err := ioutil.ReadDir(igs.dataDir)
	if err == nil {
		for _, fi := range fiList {
			if fi.IsDir() {
				indexName := fi.Name()
				indexDir := path.Join(igs.dataDir, indexName)
				index, err := bleve.Open(indexDir)
				if err == nil {
					log.Printf("info: open index indexName=\"%s\"\n", indexName)
					igs.indices[indexName] = index
				} else {
					log.Printf("error: failed to open index indexName=\"%s\"\n", indexDir)
				}
			}
		}
	} else {
		log.Printf("error: failed to read data directory dataDir=\"%s\" error=\"%s\"\n", igs.dataDir, err.Error())
	}

	return err
}

func (igs *indigoGRPCService) CloseIndices() error {
	var err error = nil

	for indexName, index := range igs.indices {
		err = index.Close()
		if err == nil {
			log.Printf("info: close index indexName=\"%s\"\n", indexName)
		} else {
			log.Printf("error: failed to close index indexName=\"%s\" error=\"%s\"\n", indexName, err.Error())
		}
	}

	return err
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, req *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	var (
		index bleve.Index
		err   error
	)

	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	indexDir := path.Join(igs.dataDir, req.IndexName)
	indexMapping := bleve.NewIndexMapping()
	kvConfig := new(map[string]interface{})

	_, ok := igs.indices[req.IndexName]
	if ok == false {
		_, err = os.Stat(indexDir)
		if os.IsNotExist(err) {
			if req.IndexMapping != nil {
				err = json.Unmarshal(req.IndexMapping, indexMapping)
				if err == nil {
					log.Printf("debug: create index mapping indexName=\"%s\"\n", req.IndexName)
				} else {
					log.Printf("error: faild to create index mapping indexName=\"%s\" error=\"%s\"\n", req.IndexName, err.Error())
				}
			} else {
				log.Printf("debug: use default index mapping indexName=\"%s\"\n", req.IndexName)
			}

			if req.KvConfig != nil {
				err = json.Unmarshal(req.KvConfig, kvConfig)
				if err == nil {
					log.Printf("debug: create kv config indexName=\"%s\"\n", req.IndexName)
				} else {
					log.Printf("error: faild to create kv config indexName=\"%s\" error=\"%s\"\n", req.IndexName, err.Error())
				}
			} else {
				log.Printf("debug: use default kv config indexName=\"%s\"\n", req.IndexName)
			}

			index, err = bleve.NewUsing(indexDir, indexMapping, req.IndexType, req.KvStore, *kvConfig)
			if err == nil {
				log.Printf("info: create index indexName=\"%s\" indexDir=\"%s\" indexType=\"%s\" kvStore=\"%s\"\n", req.IndexName, indexDir, req.IndexType, req.KvStore)
				log.Printf("info: open index indexName=\"%s\" indexDir=\"%s\" indexType=\"%s\" kvStore=\"%s\"\n", req.IndexName, indexDir, req.IndexType, req.KvStore)
				igs.indices[req.IndexName] = index
			} else {
				log.Printf("error: faild to create index (%s) indexDir=\"%s\" indexDir=\"%s\" indexType=\"%s\" kvStore=\"%s\"\n", err.Error(), req.IndexName, indexDir, req.IndexType, req.KvStore)
			}
		} else {
			err = errors.New("directory exists")
			log.Printf("error: index directory already exists indexDir=\"%s\" error=\"%s\"\n", indexDir, err.Error())
		}
	} else {
		err = errors.New("index opend")
		log.Printf("error: index already opened indexName=\"%s\" error=\"%s\"\n", req.IndexName, err.Error())
	}

	return &proto.CreateIndexResponse{IndexName: req.IndexName}, err
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	var err error

	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	indexDir := path.Join(igs.dataDir, req.IndexName)

	_, ok := igs.indices[req.IndexName]
	if ok == false {
		_, err = os.Stat(indexDir)
		if err == nil {
			err = os.RemoveAll(indexDir)
			if err == nil {
				log.Printf("info: delete index indexDir=\"%s\"\n", indexDir)
			} else {
				log.Printf("error: failed to delete index directory (%s) index_dir=\"%s\"\n", err.Error(), indexDir)
			}
		} else {
			log.Printf("error: index directory does not exist (%s) index_dir=\"%s\"\n", err.Error(), indexDir)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s already exists", req.IndexName))
		log.Printf("error: index has been opened (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.DeleteIndexResponse{IndexName: req.IndexName}, err
}

func (igs *indigoGRPCService) OpenIndex(ctx context.Context, req *proto.OpenIndexRequest) (*proto.OpenIndexResponse, error) {
	var (
		index bleve.Index
		err   error
	)

	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	indexDir := path.Join(igs.dataDir, req.IndexName)
	runtimeConfig := new(map[string]interface{})

	_, ok := igs.indices[req.IndexName]
	if ok == false {
		_, err = os.Stat(indexDir)
		if err == nil {
			if req.RuntimeConfig != nil {
				err = json.Unmarshal(req.RuntimeConfig, runtimeConfig)
				if err == nil {
					log.Printf("debug: create runtime config index_name=\"%s\"\n", req.IndexName)
				} else {
					log.Printf("error: faild to create runtime config (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
				}
			} else {
				log.Printf("debug: use default kv config index_name=\"%s\"\n", req.IndexName)
			}

			index, err = bleve.OpenUsing(indexDir, *runtimeConfig)
			if err == nil {
				log.Printf("info: open index index_name=\"%s\"\n", req.IndexName)

				igs.indices[req.IndexName] = index
			} else {
				log.Printf("error: failed to open index (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
			}
		} else {
			log.Printf("error: index directory does not exist (%s) index_dir=\"%s\"\n", err.Error(), indexDir)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index already opened (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.OpenIndexResponse{IndexName: req.IndexName}, err
}

func (igs *indigoGRPCService) CloseIndex(ctx context.Context, req *proto.CloseIndexRequest) (*proto.CloseIndexResponse, error) {
	var (
		index bleve.Index
		err   error
	)

	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	index, ok := igs.indices[req.IndexName]
	if ok {
		err = index.Close()
		if err == nil {
			log.Printf("info: close index index_name=\"%s\"\n", req.IndexName)
			delete(igs.indices, req.IndexName)
		} else {
			log.Printf("error: failed to close index (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index already closed (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.CloseIndexResponse{IndexName: req.IndexName}, err
}

func (igs *indigoGRPCService) GetStats(ctx context.Context, req *proto.GetStatsRequest) (*proto.GetStatsResponse, error) {
	var (
		indexStat []byte
		err       error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		indexStat, err = index.Stats().MarshalJSON()
		if err == nil {
			log.Printf("debug: create index stats index_name=\"%s\"\n", req.IndexName)
		} else {
			log.Printf("error: faild to create index stats (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.GetStatsResponse{IndexStats: indexStat}, err
}

func (igs *indigoGRPCService) GetMapping(ctx context.Context, req *proto.GetMappingRequest) (*proto.GetMappingResponse, error) {
	var (
		indexMapping []byte
		err          error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		indexMapping, err = json.Marshal(index.Mapping())
		if err == nil {
			log.Printf("debug: create index mapping index_name=\"%s\"\n", req.IndexName)
		} else {
			log.Printf("error: failed to create index mapping (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.GetMappingResponse{IndexMapping: indexMapping}, err
}

func (igs *indigoGRPCService) PutDocument(ctx context.Context, req *proto.PutDocumentRequest) (*proto.PutDocumentResponse, error) {
	var (
		doc           interface{}
		putCount      int32 = 0
		putErrorCount int32 = 0
		err           error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		err = json.Unmarshal(req.Document, &doc)
		if err == nil {
			log.Printf("debug: create document index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)

			err = index.Index(req.DocumentID, doc)
			if err == nil {
				putCount++
				log.Printf("info: put document index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)
			} else {
				putErrorCount++
				log.Printf("error: failed to put document (%s) index_name=\"%s\" document_id=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
			}
		} else {
			log.Printf("error: failed to create document (%s) index_name=\"%s\" document_id=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.PutDocumentResponse{PutCount: putCount, PutErrorCount: putErrorCount}, err
}

func (igs *indigoGRPCService) GetDocument(ctx context.Context, req *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {
	var (
		bytesResp []byte
		err       error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		doc, err := index.Document(req.DocumentID)
		if err == nil {
			fields := make(map[string]interface{})

			if doc != nil {
				log.Printf("ingo: get document index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)

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
				log.Printf("debug: document does not exist index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)
			}

			bytesResp, err = json.Marshal(fields)
			if err == nil {
				log.Printf("debug: create document index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)
			} else {
				log.Printf("error: failed to create document (%s) index_name=\"%s\" document_id=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
			}
		} else {
			log.Printf("error: failed to get document (%s) index_name=\"%s\" document_id=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.GetDocumentResponse{Document: bytesResp}, err
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, req *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	var (
		deleteCount int32 = 0
		err         error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		err = index.Delete(req.DocumentID)
		if err == nil {
			deleteCount++
			log.Printf("info: delete document index_name=\"%s\" document_id=\"%s\"\n", req.IndexName, req.DocumentID)
		} else {
			log.Printf("error: failed to delete document (%s) index_name=\"%s\" document_id=%s\n", err.Error(), req.IndexName, req.DocumentID)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.DeleteDocumentResponse{DeleteCount: deleteCount}, err
}

func (igs *indigoGRPCService) Bulk(ctx context.Context, req *proto.BulkRequest) (*proto.BulkResponse, error) {
	var (
		batchCount    int32 = 0
		putCount      int32 = 0
		putErrorCount int32 = 0
		deleteCount   int32 = 0
		bulkRequest   interface{}
		err           error
	)

	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		err = json.Unmarshal(req.BulkRequest, &bulkRequest)
		if err == nil {
			log.Printf("debug: create bulk request index_name=\"%s\"\n", req.IndexName)

			batch := index.NewBatch()

			if _, ok := bulkRequest.([]interface{}); ok {
				log.Printf("debug: expected bulk request format name=\"%s\"\n", req.IndexName)

				for num, request := range bulkRequest.([]interface{}) {
					if request, ok := request.(map[string]interface{}); ok {
						log.Printf("debug: expected request format index_name=\"%s\" num=%d\n", req.IndexName, num)

						var method string
						var id string

						if _, ok := request["method"]; ok {
							log.Printf("debug: method exists in request index_name=\"%s\" num=%d\n", req.IndexName, num)
							method = request["method"].(string)
						} else {
							log.Printf("error: method does not exist in request (%s) index_name=\"%s\" num=%d\n", err.Error(), req.IndexName, num)
							continue
						}
						if _, ok := request["id"]; ok {
							log.Printf("debug: id exists in request index_name=\"%s\" num=%d\n", req.IndexName, num)
							id = request["id"].(string)
						} else {
							log.Printf("error: id does not exist in request (%s) index_name=\"%s\" num=%d\n", err.Error(), req.IndexName, num)
							continue
						}

						switch method {
						case "put":
							var doc interface{}

							if _, ok := request["document"]; ok {
								log.Printf("debug: document exists in request index_name=\"%s\" num=%d\n", req.IndexName, num)
								doc = request["document"]
							} else {
								log.Printf("error: document does not exist in request (%s) index_name=\"%s\" num=%d\n", err.Error(), req.IndexName, num)
								continue
							}

							err = batch.Index(id, doc)
							if err == nil {
								putCount++
								batchCount++
								log.Printf("info: put document index_name=\"%s\" document_id=\"%s\" num=%d\n", req.IndexName, id, num)
							} else {
								putErrorCount++
								log.Printf("error: failed to put document (%s) index_name=\"%s\" document_id=\"%s\" num=%d\n", err.Error(), req.IndexName, id, num)
							}
						case "delete":
							batch.Delete(id)
							deleteCount++
							batchCount++
							log.Printf("info: delete document index_name=\"%s\" document_id=\"%s\" num=%d\n", req.IndexName, id, num)
						default:
							log.Printf("error: unexpected method method=\"%s\" index_name=\"%s\" document_id=\"%s\"\n", method, req.IndexName, id)
						}
					} else {
						log.Printf("error: unexpected request format index_name=\"%s\"\n", req.IndexName)
					}

					if batchCount%req.BatchSize == 0 {
						err = index.Batch(batch)
						if err == nil {
							log.Printf("info: index documents in bulk index_name=\"%s\" documents=%d\n", req.IndexName, batch.Size())
						} else {
							log.Printf("error: failed to index documents in bulk (%s) index_name=\"%s\" documents=%d\n", err.Error(), req.IndexName, batch.Size())
						}

						batch = index.NewBatch()
					}
				}
			} else {
				log.Printf("error: unexpected bulk request format index_name=\"%s\"\n", req.IndexName)
			}

			if batch.Size() > 0 {
				err = index.Batch(batch)
				if err == nil {
					log.Printf("info: index documents in bulk index_name=\"%s\" documents=%d\n", req.IndexName, batch.Size())
				} else {
					log.Printf("error: failed to index documents in bulk (%s) index_name=\"%s\" documents=%d\n", err.Error(), req.IndexName, batch.Size())
				}
			}
		} else {
			log.Printf("error: failed to create bulk request (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.BulkResponse{PutCount: putCount, PutErrorCount: putErrorCount, DeleteCount: deleteCount}, err
}

func (igs *indigoGRPCService) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	var (
		bytesResp []byte
		err       error
	)

	index, ok := igs.indices[req.IndexName]
	if ok == true {
		searchRequest := bleve.NewSearchRequest(nil)
		err = json.Unmarshal(req.SearchRequest, searchRequest)
		if err == nil {
			log.Printf("debug: create search request index_name=\"%s\"\n", req.IndexName)

			searchResult, err := index.Search(searchRequest)
			if err == nil {
				log.Printf("info: search documents index_name=\"%s\"\n", req.IndexName)

				bytesResp, err = json.Marshal(&searchResult)
				if err == nil {
					log.Printf("debug: create search result index_name=\"%s\"\n", req.IndexName)
				} else {
					log.Printf("error: failed to create search result (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
				}
			} else {
				log.Printf("error: failed to search documents (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
			}
		} else {
			log.Printf("error: failed to create search request (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.IndexName))
		log.Printf("error: index does not exist (%s) index_name=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.SearchResponse{SearchResult: bytesResp}, err
}
