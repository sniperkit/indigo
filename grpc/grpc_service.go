package grpc

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	_ "github.com/mosuka/indigo/config"
	"github.com/mosuka/indigo/proto"
	"github.com/pkg/errors"
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
	indices := make(map[string]bleve.Index)
	mutexes := make(map[string]*sync.RWMutex)

	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err == nil {
			log.Printf("debug: succeeded in creating data directory dataDir=\"%s\"\n", dataDir)
		} else {
			log.Printf("error: failed to create data directory dataDir=\"%s\" error=\"%s\"\n", dataDir, err.Error())
		}
	} else {
		err = fmt.Errorf("%s already exists", dataDir)
		log.Printf("debug: data directory already exists dataDir=\"%s\" error=\"%s\"\n", dataDir, err.Error())
	}

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
		mutexes: mutexes,
	}
}

func (igs *indigoGRPCService) lockIndex(indexName string) {
	if _, existed := igs.mutexes[indexName]; !existed {
		igs.mutexes[indexName] = new(sync.RWMutex)
	}

	igs.mutexes[indexName].Lock()
	log.Printf("debug: lock index indexName=\"%s\"\n", indexName)
}

func (igs *indigoGRPCService) unlockIndex(indexName string) {
	if _, existed := igs.mutexes[indexName]; !existed {
		igs.mutexes[indexName] = new(sync.RWMutex)
	}

	igs.mutexes[indexName].Unlock()
	log.Printf("debug: unlock index indexName=\"%s\"\n", indexName)
}

func (igs *indigoGRPCService) OpenIndices() {
	if fiList, err := ioutil.ReadDir(igs.dataDir); err == nil {
		for _, fi := range fiList {
			if fi.IsDir() {
				indexName := fi.Name()
				indexDir := path.Join(igs.dataDir, indexName)
				index, err := bleve.Open(indexDir)
				if err == nil {
					log.Printf("info: succeeded in opening index indexName=\"%s\"\n", indexName)
					igs.indices[indexName] = index
				} else {
					log.Printf("warn: %s indexName=\"%s\"\n", err.Error(), igs.dataDir)
				}
			}
		}
	} else {
		log.Printf("warn: %s dataDir=\"%s\"\n", err.Error(), igs.dataDir)
	}

	return
}

func (igs *indigoGRPCService) CloseIndices() {
	for indexName, index := range igs.indices {
		if err := index.Close(); err == nil {
			log.Printf("info: succeeded in closing index indexName=\"%s\"\n", indexName)
		} else {
			log.Printf("warn: failed to close index indexName=\"%s\" error=\"%s\"\n", indexName, err.Error())
		}
	}

	return
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, req *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.CreateIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	indexMapping := bleve.NewIndexMapping()
	if req.IndexMapping != nil {
		if err := json.Unmarshal(req.IndexMapping, &indexMapping); err == nil {
			log.Printf("debug: succeeded in creating index mapping indexName=\"%s\"\n", req.IndexName)
		} else {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
			return &proto.CreateIndexResponse{}, err
		}
	}

	kvConfig := make(map[string]interface{})
	if req.KvConfig != nil {
		if err := json.Unmarshal(req.KvConfig, &kvConfig); err == nil {
			log.Printf("debug: succeeded in creating kv config indexName=\"%s\"\n", req.IndexName)
		} else {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
			return &proto.CreateIndexResponse{}, err
		}
	}

	_, err := os.Stat(indexDir)
	if os.IsNotExist(err) {
		var index bleve.Index = nil
		index, err = bleve.NewUsing(indexDir, indexMapping, req.IndexType, req.KvStore, kvConfig)
		if err == nil {
			log.Printf("info: succeeded in creating index indexName=\"%s\"\n", req.IndexName)
			igs.indices[req.IndexName] = index
		} else {
			log.Printf("error: %s indexDir=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		err = errors.New("index directory already exists")
		log.Printf("error: %s indexDir=\"%s\"\n", err.Error(), indexDir)
	}

	return &proto.CreateIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.DeleteIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	_, err := os.Stat(indexDir)
	if err == nil {
		err = os.RemoveAll(indexDir)
		if err == nil {
			log.Printf("info: succeeded in deleting index indexDir=\"%s\"\n", indexDir)
		} else {
			log.Printf("error: %s indexDir=\"%s\"\n", err.Error(), indexDir)
		}
	} else {
		log.Printf("error: %s indexDir=\"%s\"\n", err.Error(), indexDir)
	}

	return &proto.DeleteIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) OpenIndex(ctx context.Context, req *proto.OpenIndexRequest) (*proto.OpenIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	if _, open := igs.indices[req.IndexName]; open {
		err := errors.New("index already opened")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.OpenIndexResponse{}, err
	}

	indexDir := path.Join(igs.dataDir, req.IndexName)

	runtimeConfig := make(map[string]interface{})
	if req.RuntimeConfig != nil {
		err := json.Unmarshal(req.RuntimeConfig, runtimeConfig)
		if err == nil {
			log.Printf("debug: succeeded in creating runtime config indexName=\"%s\"\n", req.IndexName)
		} else {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
			return &proto.OpenIndexResponse{}, err
		}
	}

	_, err := os.Stat(indexDir)
	if err == nil {
		var index bleve.Index = nil
		index, err = bleve.OpenUsing(indexDir, runtimeConfig)
		if err == nil {
			log.Printf("info: succeeded in opening index indexName=\"%s\"\n", req.IndexName)
			igs.indices[req.IndexName] = index
		} else {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		}
	} else {
		log.Printf("error: %s indexDir=\"%s\"\n", err.Error(), indexDir)
	}

	return &proto.OpenIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) CloseIndex(ctx context.Context, req *proto.CloseIndexRequest) (*proto.CloseIndexResponse, error) {
	igs.lockIndex(req.IndexName)
	defer igs.unlockIndex(req.IndexName)

	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.CloseIndexResponse{}, err
	}

	err := index.Close()
	if err == nil {
		log.Printf("info: succeeded in closing index indexName=\"%s\"\n", req.IndexName)
		delete(igs.indices, req.IndexName)
	} else {
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.CloseIndexResponse{
		IndexName: req.IndexName,
	}, err
}

func (igs *indigoGRPCService) GetDocumentCount(ctx context.Context, req *proto.GetDocumentCountRequest) (*proto.GetDocumentCountResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.GetDocumentCountResponse{}, err
	}

	count, err := index.DocCount()

	return &proto.GetDocumentCountResponse{DocumentCount: count}, err
}

func (igs *indigoGRPCService) GetStats(ctx context.Context, req *proto.GetStatsRequest) (*proto.GetStatsResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.GetStatsResponse{}, err
	}

	bytesIndexStat, err := index.Stats().MarshalJSON()
	if err == nil {
		log.Printf("info: succeeded in creating index stats indexName=\"%s\"\n", req.IndexName)
	} else {
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.GetStatsResponse{
		IndexStats: bytesIndexStat,
	}, err
}

func (igs *indigoGRPCService) GetMapping(ctx context.Context, req *proto.GetMappingRequest) (*proto.GetMappingResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.GetMappingResponse{}, err
	}

	indexMapping, err := json.Marshal(index.Mapping())
	if err == nil {
		log.Printf("info: succeeded in creating index mapping indexName=\"%s\"\n", req.IndexName)
	} else {
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
	}

	return &proto.GetMappingResponse{
		IndexMapping: indexMapping,
	}, err
}

func (igs *indigoGRPCService) PutDocument(ctx context.Context, req *proto.PutDocumentRequest) (*proto.PutDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.PutDocumentResponse{}, err
	}

	success := false
	var doc interface{}
	err := json.Unmarshal(req.Document, &doc)
	if err == nil {
		log.Printf("debug: succeeded in creating document indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)

		err = index.Index(req.DocumentID, doc)
		if err == nil {
			success = true
			log.Printf("info: succeeded in putting document indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)
		} else {
			log.Printf("error: %s indexName=\"%s\" documentID=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
		}
	} else {
		log.Printf("error: %s indexName=\"%s\" documentID=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
	}

	return &proto.PutDocumentResponse{
		Success: success,
	}, err
}

func (igs *indigoGRPCService) GetDocument(ctx context.Context, req *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.GetDocumentResponse{}, err
	}

	doc := make(map[string]interface{})
	if d, err := index.Document(req.DocumentID); err == nil {
		if d != nil {
			log.Printf("ingo: succeeded in getting document indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)

			for _, field := range d.Fields {
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

				existedField, existed := doc[field.Name()]
				if existed {
					switch existedField := existedField.(type) {
					case []interface{}:
						doc[field.Name()] = append(existedField, value)
					case interface{}:
						arr := make([]interface{}, 2)
						arr[0] = existedField
						arr[1] = value
						doc[field.Name()] = arr
					}
				} else {
					doc[field.Name()] = value
				}
			}
		} else {
			log.Printf("info: document does not exist indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)
		}
	} else {
		log.Printf("error: %s indexName=\"%s\" documentID=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
		return &proto.GetDocumentResponse{}, err
	}

	bytesDoc, err := json.Marshal(doc)
	if err == nil {
		log.Printf("debug: succeeded in creating document indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)
	} else {
		log.Printf("error: %s index_name=\"%s\" document_id=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
	}

	return &proto.GetDocumentResponse{
		Document: bytesDoc,
	}, err
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, req *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	index, open := igs.indices[req.IndexName]
	if !open {
		err := errors.New("index is not open")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.DeleteDocumentResponse{}, err
	}

	success := false
	err := index.Delete(req.DocumentID)
	if err == nil {
		success = true
		log.Printf("info: succeeded in deleting document indexName=\"%s\" documentID=\"%s\"\n", req.IndexName, req.DocumentID)
	} else {
		log.Printf("error: %s indexName=\"%s\" documentID=\"%s\"\n", err.Error(), req.IndexName, req.DocumentID)
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
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.BulkResponse{}, err
	}

	var bulkRequest interface{}
	if req.BulkRequest != nil {
		err := json.Unmarshal(req.BulkRequest, &bulkRequest)
		if err != nil {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
			return &proto.BulkResponse{}, err
		}
	}

	_, ok := bulkRequest.([]interface{})
	if !ok {
		err := errors.New("unexpected bulk request format")
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
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
			log.Printf("error: unexpected request format indexName=\"%s\" num=%d\n", req.IndexName, num)
			continue
		}

		var method string
		if _, ok := request["method"]; ok {
			log.Printf("debug: method exists in request indexName=\"%s\" num=%d\n", req.IndexName, num)
			method = request["method"].(string)
		} else {
			log.Printf("error: method does not exist in request indexName=\"%s\" num=%d\n", req.IndexName, num)
			continue
		}

		var id string
		if _, ok := request["id"]; ok {
			log.Printf("debug: id exists in request indexName=\"%s\" num=%d\n", req.IndexName, num)
			id = request["id"].(string)
		} else {
			log.Printf("error: id does not exist in request indexName=\"%s\" num=%d\n", req.IndexName, num)
			continue
		}

		switch method {
		case "put":
			var doc interface{}

			if _, ok := request["document"]; ok {
				log.Printf("debug: document exists in request indexName=\"%s\" num=%d\n", req.IndexName, num)
				doc = request["document"]
			} else {
				log.Printf("error: document does not exist in request indexName=\"%s\" num=%d error=\"%s\"\n", req.IndexName, num)
				continue
			}

			err := batch.Index(id, doc)
			if err == nil {
				log.Printf("info: succeeded in putting document indexName=\"%s\" documentID=\"%s\" num=%d\n", req.IndexName, id, num)
				putCount++
				batchCount++
			} else {
				log.Printf("error: %s indexName=\"%s\" documentID=\"%s\" num=%d\n", err.Error(), req.IndexName, id, num)
				putErrorCount++
			}
		case "delete":
			batch.Delete(id)
			log.Printf("info: succeeded in deleting document indexName=\"%s\" documentID=\"%s\" num=%d\n", req.IndexName, id, num)
			deleteCount++
			batchCount++
		default:
			log.Printf("error: unexpected method method=\"%s\" indexName=\"%s\" documentID=\"%s\"\n", method, req.IndexName, id)
			continue
		}

		if batchCount%req.BatchSize == 0 {
			err := index.Batch(batch)
			if err == nil {
				log.Printf("info: succeeded in indexing documents in bulk indexName=\"%s\" documents=%d\n", req.IndexName, batch.Size())
			} else {
				log.Printf("error: %s indexName=\"%s\" documents=%d\n", err.Error(), req.IndexName, batch.Size())
			}

			batch = index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		err := index.Batch(batch)
		if err == nil {
			log.Printf("info: succeeded in indexing documents in bulk indexName=\"%s\" documents=%d\n", req.IndexName, batch.Size())
		} else {
			log.Printf("error: %s indexName=\"%s\" documents=%d\n", err.Error(), req.IndexName, batch.Size())
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
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.SearchResponse{}, err
	}

	searchRequest := bleve.NewSearchRequest(nil)
	if req.SearchRequest != nil {
		err := json.Unmarshal(req.SearchRequest, searchRequest)
		if err != nil {
			log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
			return &proto.SearchResponse{}, err
		}
	}

	searchResult, err := index.Search(searchRequest)
	if err == nil {
		log.Printf("info: succeeded in searching documents indexName=\"%s\"\n", req.IndexName)
	} else {
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
		return &proto.SearchResponse{}, err
	}

	bytesSearchResult, err := json.Marshal(&searchResult)
	if err == nil {
		log.Printf("debug: succeeded in creating search result indexName=\"%s\"\n", req.IndexName)
	} else {
		log.Printf("error: %s indexName=\"%s\"\n", err.Error(), req.IndexName)
	}
	return &proto.SearchResponse{
		SearchResult: bytesSearchResult,
	}, err
}
