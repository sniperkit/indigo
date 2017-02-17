package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	_ "github.com/mosuka/indigo/config"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

type indigoGRPCService struct {
	mutex   sync.RWMutex
	dataDir string
	indices map[string]bleve.Index
}

func NewIndigoGRPCService(dataDir string) *indigoGRPCService {
	var indices map[string]bleve.Index = make(map[string]bleve.Index)
	var err error

	_, err = os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err == nil {
			log.Printf("info: make data directory dir=%s\n", dataDir)
		} else {
			log.Printf("error: failed to make data directory (%s) dir=%s\n", err.Error(), dataDir)
		}
	} else {
		log.Printf("info: data directory already exists dir=\"%s\"\n", dataDir)
	}

	list, err := ioutil.ReadDir(dataDir)
	if err == nil {
		for _, f := range list {
			if f.IsDir() {
				var indexDir = path.Join(dataDir, f.Name())
				index, err := bleve.Open(indexDir)
				if err == nil {
					log.Printf("info: open existing index name=\"%s\"\n", f.Name())
					indices[f.Name()] = index
				} else {
					log.Printf("error: failed to open index dir=%s\n", indexDir)
				}
			}
		}
	} else {
		log.Printf("error: failed to read data directory dir=%s\n", dataDir)
	}

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
	}
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, req *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	igs.mutex.Lock()
	defer igs.mutex.Unlock()

	var indexPath = path.Join(igs.dataDir, req.Name)
	var index bleve.Index = nil
	var indexMapping = bleve.NewIndexMapping()
	var err error = nil

	_, ok := igs.indices[req.Name]
	if ok == false {
		_, err = os.Stat(indexPath)
		if os.IsNotExist(err) {
			err = json.Unmarshal(req.Mapping, indexMapping)
			if err == nil {
				log.Printf("info: create index mapping name=%s\n", req.Name)

				index, err = bleve.NewUsing(indexPath, indexMapping, req.Type, req.Store, nil)
				if err == nil {
					log.Printf("info: create index name=%s\n", req.Name)
				} else {
					log.Printf("error: faild to create index (%s) name=%s\n", err.Error(), req.Name)
				}
			} else {
				log.Printf("error: faild to create index mapping (%s) name=%s\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: index directory exists (%s) name=%s\n", err.Error(), req.Name)
		}

		igs.indices[req.Name] = index
	} else {
		err = errors.New(fmt.Sprintf("%s exists", req.Name))
		log.Printf("error: index exists (%s) name=%s\n", err.Error(), req.Name)
	}

	return &proto.CreateIndexResponse{Name: req.Name}, nil
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, req *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	igs.mutex.Lock()
	defer igs.mutex.Unlock()
ぎt
	var indexPath = path.Join(igs.dataDir, req.Name)
	var err error

	_, ok := igs.indices[req.Name]
	if ok == true {
		_, err = os.Stat(indexPath)
		if err == nil {
			index := igs.indices[req.Name]

			err = index.Close()
			if err == nil {
				log.Printf("info: close index name=%s\n", req.Name)
			} else {
				log.Printf("error: failed to close index (%s) name=%s\n", err.Error(), req.Name)
			}

			err = os.RemoveAll(path.Join(igs.dataDir, req.Name))
			if err == nil {
				log.Printf("info: delete index name=%s\n", req.Name)
			} else {
				log.Printf("error: failed to delete index (%s) name=%s\n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: index directory does not exist (%s) name=%s\n", err.Error(), req.Name)
		}

		delete(igs.indices, req.Name)
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index does not exist (%s) name=%s\n", err.Error(), req.Name)
	}

	return &proto.DeleteIndexResponse{Name: req.Name}, nil
}

func (igs *indigoGRPCService) GetStats(ctx context.Context, req *proto.GetStatsRequest) (*proto.GetStatsResponse, error) {
	var indexStat []byte
	var err error

	_, ok := igs.indices[req.Name]
	if ok == true {
		index := igs.indices[req.Name]
		indexStat, err = index.Stats().MarshalJSON()
		if err == nil {
			log.Printf("info: create index stats name=%s\n", req.Name)
		} else {
			log.Printf("error: faild to create index stats (%s) name=%s\n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index does not exist (%s) name=%s\n", err.Error(), req.Name)
	}

	return &proto.GetStatsResponse{Stats: indexStat}, nil
}

func (igs *indigoGRPCService) GetMapping(ctx context.Context, req *proto.GetMappingRequest) (*proto.GetMappingResponse, error) {
	var indexMapping []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		indexMapping, err = json.Marshal(index.Mapping())
		if err == nil {
			log.Printf("info: create index mapping name=%s\n", req.Name)
		} else {
			log.Printf("error: failed to create index mapping (%s)\n", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.GetMappingResponse{Mapping: indexMapping}, nil
}

func (igs *indigoGRPCService) IndexDocument(ctx context.Context, req *proto.IndexDocumentRequest) (*proto.IndexDocumentResponse, error) {
	var document interface{}
	var err error

	var count int32 = 0

	index, ok := igs.indices[req.Name]
	if ok == true {
		err = json.Unmarshal(req.Document, &document)
		if err == nil {
			log.Printf("info: create document name=%s id=%s\n", req.Name, req.Id)

			err = index.Index(req.Id, document)
			if err == nil {
				count++
				log.Printf("info: index document name=%s id=%s\n", req.Name, req.Id)
			} else {
				log.Printf("error: failed to index document (%s) name=%s id=%s\n", err.Error(), req.Name, req.Id)
			}
		} else {
			log.Printf("error: failed to create document (%s) name=%s id=%s\n", err.Error(), req.Name, req.Id)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.IndexDocumentResponse{Count: count}, nil
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, req *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	var err error

	var count int32 = 0

	_, ok := igs.indices[req.Name]
	if ok == true {
		index := igs.indices[req.Name]

		err = index.Delete(req.Id)
		if err == nil {
			count++
			log.Printf("info: delete document name=%s id=%s\n", req.Name, req.Id)
		} else {
			log.Printf("error: failed to delete document (%s) name=%s id=%\n", err.Error(), req.Name, req.Id)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.DeleteDocumentResponse{Count: count}, nil
}

type IndexResult struct {
	DocumentCount int `json:"document_count"`
}

func (igs *indigoGRPCService) IndexDocuments(ctx context.Context, req *proto.IndexDocumentsRequest) (*proto.IndexDocumentsResponse, error) {
	var documents interface{}
	var err error
	var count int32 = 0
	var batchCount int32 = 0

	_, ok := igs.indices[req.Name]
	if ok == true {
		index := igs.indices[req.Name]

		err = json.Unmarshal(req.Documents, &documents)
		if err == nil {
			log.Printf("info: create documents name=%s\n", req.Name)

			batch := index.NewBatch()
			for id, doc := range documents.(map[string]interface{}) {
				err = batch.Index(id, doc)
				if err == nil {
					count++
					log.Printf("info: index document name=%s id=%s\n", req.Name, id)
				} else {
					log.Printf("error: failed to index document (%s) name=%s id=%s\n", err.Error(), req.Name, id)
				}

				batchCount++

				if batchCount%req.BatchSize == 0 {
					err = index.Batch(batch)
					if err == nil {
						log.Printf("info: index documents in bulk name=%s\n", req.Name)
					} else {
						log.Printf("error: failed to index documents in bulk (%s) name=%s\n", err.Error(), req.Name)
					}

					batch = index.NewBatch()
				}
			}

			if batch.Size() > 0 {
				err = index.Batch(batch)
				if err == nil {
					log.Printf("info: index documents in bulk name=%s\n", req.Name)
				} else {
					log.Printf("error: failed to index documents in bulk (%s) name=%s\n", err.Error(), req.Name)
				}
			}
		} else {
			log.Printf("error: failed to create documents (%s) name=%s\n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.IndexDocumentsResponse{Count: count}, nil
}

func (igs *indigoGRPCService) DeleteDocuments(ctx context.Context, req *proto.DeleteDocumentsRequest) (*proto.DeleteDocumentsResponse, error) {
	var ids []string
	var err error

	var count int32 = 0

	index, ok := igs.indices[req.Name]
	if ok == true {
		err = json.Unmarshal(req.Ids, &ids)
		if err == nil {
			log.Printf("info: create document ids name=%s\n", req.Name)

			batch := index.NewBatch()
			for i := range ids {
				id := ids[i]

				batch.Delete(id)
				log.Printf("info: delete document name=%s id=%s\n", req.Name, id)

				count++

				if count%req.BatchSize == 0 {
					err = index.Batch(batch)
					if err == nil {
						log.Printf("info: delete documents in bulk name=%s\n", req.Name)
					} else {
						log.Printf("error: failed to delete documents in bulk (%s) name=%s\n", err.Error(), req.Name)
					}

					batch = index.NewBatch()
				}
			}

			if batch.Size() > 0 {
				err = index.Batch(batch)
				if err == nil {
					log.Printf("info: delete documents in bulk name=%s\n", req.Name)
				} else {
					log.Printf("error: failed to delete documents in bulk (%s) name=%s\n", err.Error(), req.Name)
				}
			}
		} else {
			log.Printf("error: failed to create document ids (%s)", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.DeleteDocumentsResponse{Count: count}, nil
}

func (igs *indigoGRPCService) SearchDocuments(ctx context.Context, req *proto.SearchDocumentsRequest) (*proto.SearchDocumentsResponse, error) {
	log.Printf("info: search documents IndexName=%s\n", req.Name)

	var bytesResp []byte
	var err error

	index, ok := igs.indices[req.Name]
	if ok == true {
		searchRequest := bleve.NewSearchRequest(nil)
		err = json.Unmarshal(req.SearchRequest, searchRequest)
		if err == nil {
			log.Printf("info: create search request name=%s\n", req.Name)

			searchResult, err := index.Search(searchRequest)
			if err == nil {
				log.Printf("info: search documents name=%s\n", req.Name)

				bytesResp, err = json.Marshal(&searchResult)
				if err == nil {
					log.Printf("info: create search result name=%s\n", req.Name)
				} else {
					log.Printf("error: failed to create search result (%s) name=%s\n", err.Error(), req.Name)
				}
			} else {
				log.Printf("error: failed to search documents (%s) name=%s \n", err.Error(), req.Name)
			}
		} else {
			log.Printf("error: failed to create search request (%s) name=%s \n", err.Error(), req.Name)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", req.Name))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.SearchDocumentsResponse{SearchResult: bytesResp}, nil
}
