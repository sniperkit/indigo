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
	mutex   sync.RWMutex
	dataDir string
	indices map[string]bleve.Index
}

func NewIndigoGRPCService(dataDir string) *indigoGRPCService {
	indices := make(map[string]bleve.Index)
	var err error

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err == nil {
			log.Printf("info: make data directory dir=\"%s\"\n", dataDir)
		} else {
			log.Printf("error: failed to make data directory (%s) dir=\"%s\"\n", err.Error(), dataDir)
		}
	} else {
		log.Printf("info: data directory already exists dir=\"%s\"\n", dataDir)
	}

	fiList, err := ioutil.ReadDir(dataDir)
	if err == nil {
		for _, fi := range fiList {
			if fi.IsDir() {
				indexDir := path.Join(dataDir, fi.Name())
				index, err := bleve.Open(indexDir)
				if err == nil {
					log.Printf("info: open existing index name=\"%s\"\n", fi.Name())
					indices[fi.Name()] = index
				} else {
					log.Printf("error: failed to open index dir=\"%s\"\n", indexDir)
				}
			}
		}
	} else {
		log.Printf("error: failed to read data directory dir=\"%s\"\n", dataDir)
	}

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
	}
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, req *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	igs.mutex.Lock()
	defer igs.mutex.Unlock()

	indexPath := path.Join(igs.dataDir, req.Name)
	indexMapping := bleve.NewIndexMapping()
	var index bleve.Index
	var err error

	_, ok := igs.indices[req.Name]
	if ok == false {
		_, err = os.Stat(indexPath)
		if os.IsNotExist(err) {
			err = json.Unmarshal(req.Mapping, indexMapping)
			if err == nil {
				log.Printf("info: create index mapping name=\"%s\"\n", req.Name)

				index, err = bleve.NewUsing(indexPath, indexMapping, req.Type, req.Store, nil)
				if err == nil {
					log.Printf("info: create index name=\"%s\"\n", req.Name)
				} else {
					log.Printf("error: faild to create index (%s) name=\"%s\"\n", err.Error(), req.Name)
				}
			} else {
				log.Printf("error: faild to create index mapping (%s) name=\"%s\"\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: index directory exists (%s) name=\"%s\"\n", err.Error(), req.Name)
		}

		igs.indices[req.Name] = index
	} else {
		err = errors.New(fmt.Sprintf("%s already exists", req.Name))
		log.Printf("error: index exists (%s) name=\"%s\"\n", err.Error(), req.Name)
	}

	return &proto.CreateIndexResponse{Name: req.Name}, err
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	igs.mutex.Lock()
	defer igs.mutex.Unlock()

	indexPath := path.Join(igs.dataDir, req.Name)
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		_, err = os.Stat(indexPath)
		if err == nil {
			err = index.Close()
			if err == nil {
				log.Printf("info: close index name=\"%s\"\n", req.Name)
			} else {
				log.Printf("error: failed to close index (%s) name=\"%s\"\n", err.Error(), req.Name)
			}

			err = os.RemoveAll(path.Join(igs.dataDir, req.Name))
			if err == nil {
				log.Printf("info: delete index name=\"%s\"\n", req.Name)
			} else {
				log.Printf("error: failed to delete index (%s) name=\"%s\"\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: index directory does not exist (%s) name=\"%s\"\n", err.Error(), req.Name)
		}

		delete(igs.indices, req.Name)
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index does not exist (%s) name=\"%s\"\n", err.Error(), req.Name)
	}

	return &proto.DeleteIndexResponse{Name: req.Name}, err
}

func (igs *indigoGRPCService) GetStats(ctx context.Context, req *proto.GetStatsRequest) (*proto.GetStatsResponse, error) {
	var indexStat []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		indexStat, err = index.Stats().MarshalJSON()
		if err == nil {
			log.Printf("info: create index stats name=\"%s\"\n", req.Name)
		} else {
			log.Printf("error: faild to create index stats (%s) name=\"%s\"\n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index does not exist (%s) name=\"%s\"\n", err.Error(), req.Name)
	}

	return &proto.GetStatsResponse{Stats: indexStat}, err
}

func (igs *indigoGRPCService) GetMapping(ctx context.Context, req *proto.GetMappingRequest) (*proto.GetMappingResponse, error) {
	var indexMapping []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		indexMapping, err = json.Marshal(index.Mapping())
		if err == nil {
			log.Printf("info: create index mapping name=\"%s\"\n", req.Name)
		} else {
			log.Printf("error: failed to create index mapping (%s)\n", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.GetMappingResponse{Mapping: indexMapping}, err
}

func (igs *indigoGRPCService) PutDocument(ctx context.Context, req *proto.PutDocumentRequest) (*proto.PutDocumentResponse, error) {
	var doc interface{}
	var err error
	count := 0

	index, ok := igs.indices[req.Name]
	if ok == true {
		err = json.Unmarshal(req.Document, &doc)
		if err == nil {
			log.Printf("info: create document name=\"%s\" id=\"%s\"\n", req.Name, req.Id)

			err = index.Index(req.Id, doc)
			if err == nil {
				count++
				log.Printf("info: index document name=\"%s\" id=\"%s\"\n", req.Name, req.Id)
			} else {
				log.Printf("error: failed to index document (%s) name=\"%s\" id=\"%s\"\n", err.Error(), req.Name, req.Id)
			}
		} else {
			log.Printf("error: failed to create document (%s) name=\"%s\" id=\"%s\"\n", err.Error(), req.Name, req.Id)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.PutDocumentResponse{Count: int32(count)}, err
}

func (igs *indigoGRPCService) GetDocument(ctx context.Context, req *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {
	var bytesResp []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		doc, err := index.Document(req.Id)
		if err == nil {
			fields := make(map[string]interface{})

			if doc != nil {
				log.Printf("info: document exists name=\"%s\" id=\"%s\"\n", req.Name, req.Id)

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
				log.Printf("info: document does not exist name=\"%s\" id=\"%s\"\n", req.Name, req.Id)
			}

			bytesResp, err = json.Marshal(fields)
			if err == nil {
				log.Printf("info: create document name=\"%s\"\n", req.Name)
			} else {
				log.Printf("error: failed to create document (%s) name=\"%s\"\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: failed to get document (%s) name=\"%s\" id=\"%s\"\n", err.Error(), req.Name, req.Id)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.GetDocumentResponse{Document: bytesResp}, err
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, req *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	var err error
	count := 0

	index, ok := igs.indices[req.Name]
	if ok == true {
		err = index.Delete(req.Id)
		if err == nil {
			count++
			log.Printf("debug: delete document name=\"%s\" id=\"%s\"\n", req.Name, req.Id)
		} else {
			log.Printf("error: failed to delete document (%s) name=\"%s\" id=%s\n", err.Error(), req.Name, req.Id)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.DeleteDocumentResponse{Count: int32(count)}, err
}

func (igs *indigoGRPCService) Bulk(ctx context.Context, req *proto.BulkRequest) (*proto.BulkResponse, error) {
	var batchCount int32 = 0
	var putCount int32 = 0
	var putErrorCount int32 = 0
	var deleteCount int32 = 0
	var bulkRequest interface{}
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		err = json.Unmarshal(req.BulkRequest, &bulkRequest)
		if err == nil {
			log.Printf("debug: create documents name=\"%s\"\n", req.Name)

			batch := index.NewBatch()

			if _, ok := bulkRequest.([]interface{}); ok {
				log.Printf("debug: expected bulk request format name=\"%s\"\n", req.Name)

				for num, request := range bulkRequest.([]interface{}) {
					if request, ok := request.(map[string]interface{}); ok {
						log.Printf("debug: expected request format name=\"%s\" num=%d\n", req.Name, num)

						var method string
						var id string

						if _, ok := request["method"]; ok {
							log.Printf("debug: method exists in request name=\"%s\" num=%d\n", req.Name, num)
							method = request["method"].(string)
						} else {
							log.Printf("error: method does not exist in request (%s) name=\"%s\" num=%d\n", err.Error(), req.Name, num)
							continue
						}
						if _, ok := request["id"]; ok {
							log.Printf("debug: id exists in request name=\"%s\" num=%d\n", req.Name, num)
							id = request["id"].(string)
						} else {
							log.Printf("error: id does not exist in request (%s) name=\"%s\" num=%d\n", err.Error(), req.Name, num)
							continue
						}

						switch method {
						case "put":
							var document interface{}

							if _, ok := request["document"]; ok {
								log.Printf("debug: document exists in request name=\"%s\" num=%d\n", req.Name, num)
								document = request["document"]
							} else {
								log.Printf("error: document does not exist in request (%s) name=\"%s\" num=%d\n", err.Error(), req.Name, num)
								continue
							}

							err = batch.Index(id, document)
							if err == nil {
								putCount++
								log.Printf("debug: index document name=\"%s\" id=\"%s\"\n", req.Name, id)
							} else {
								putErrorCount++
								log.Printf("error: failed to index document (%s) name=\"%s\" id=\"%s\"\n", err.Error(), req.Name, id)
							}
						case "delete":
							batch.Delete(id)
							deleteCount++
							log.Printf("debug: delete document name=\"%s\" id=\"%s\"\n", req.Name, id)
						default:
							log.Printf("error: unexpected method name=\"%s\" method=\"%s\" id=\"%s\"\n", req.Name, method, id)
						}
						batchCount++
					} else {
						log.Printf("error: unexpected request format name=\"%s\"\n", req.Name)
					}

					if batchCount%req.BatchSize == 0 {
						err = index.Batch(batch)
						if err == nil {
							log.Printf("info: index documents in bulk name=\"%s\"\n", req.Name)
						} else {
							log.Printf("error: failed to index documents in bulk (%s) name=\"%s\"\n", err.Error(), req.Name)
						}

						batch = index.NewBatch()
					}
				}
			} else {
				log.Printf("error: unexpected bulk request format name=\"%s\"\n", req.Name)
			}

			if batch.Size() > 0 {
				err = index.Batch(batch)
				if err == nil {
					log.Printf("info: index documents in bulk name=\"%s\"\n", req.Name)
				} else {
					log.Printf("error: failed to index documents in bulk (%s) name=\"%s\"\n", err.Error(), req.Name)
				}
			}
		} else {
			log.Printf("error: failed to create documents (%s) name=\"%s\"\n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.BulkResponse{PutCount: putCount, PutErrorCount: putErrorCount, DeleteCount: deleteCount}, err
}

func (igs *indigoGRPCService) Search(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	var bytesResp []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		searchRequest := bleve.NewSearchRequest(nil)
		err = json.Unmarshal(req.SearchRequest, searchRequest)
		if err == nil {
			log.Printf("info: create search request name=\"%s\"\n", req.Name)

			searchResult, err := index.Search(searchRequest)
			if err == nil {
				log.Printf("info: search documents name=\"%s\"\n", req.Name)

				bytesResp, err = json.Marshal(&searchResult)
				if err == nil {
					log.Printf("info: create search result name=\"%s\"\n", req.Name)
				} else {
					log.Printf("error: failed to create search result (%s) name=\"%s\"\n", err.Error(), req.Name)
				}
			} else {
				log.Printf("error: failed to search documents (%s) name=\"%s\"\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: failed to create search request (%s) name=\"%s\"\n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.SearchResponse{SearchResult: bytesResp}, err
}
