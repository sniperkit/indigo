package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	_ "github.com/mosuka/indigo/config"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
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
	indices := make(map[string]bleve.Index)

	return &indigoGRPCService{
		dataDir: dataDir,
		indices: indices,
	}
}

func (igs *indigoGRPCService) CreateIndex(ctx context.Context, r *proto.CreateIndexRequest) (*proto.CreateIndexResponse, error) {
	log.Printf("info: create index index_name=%s\n", r.IndexName)

	igs.mutex.Lock()
	defer igs.mutex.Unlock()

	var indexPath = path.Join(igs.dataDir, r.IndexName)
	var index bleve.Index
	var indexMapping = bleve.NewIndexMapping()
	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == false {
		_, err = os.Stat(indexPath)
		if os.IsNotExist(err) {
			err = json.Unmarshal([]byte(r.IndexMapping), indexMapping)
			if err == nil {
				index, err = bleve.NewUsing(indexPath, indexMapping, r.IndexType, r.IndexStore, nil)
				if err != nil {
					log.Printf("error: faild to create index (%s) index_name=%s\n", err.Error(), r.IndexName)
				}
			} else {
				log.Printf("error: faild to create index mapping (%s) index_name=%s\n", err.Error(), r.IndexName)
			}
		} else {
			log.Printf("error: index directory exists (%s) index_name=%s\n", err.Error(), r.IndexName)
		}

		igs.indices[r.IndexName] = index
	} else {
		err = errors.New(fmt.Sprintf("%s exists", r.IndexName))
		log.Printf("error: index name exists (%s) index_name=%s\n", err.Error(), r.IndexName)
	}

	return &proto.CreateIndexResponse{Result: indexPath}, nil
}

func (igs *indigoGRPCService) DeleteIndex(ctx context.Context, r *proto.DeleteIndexRequest) (*proto.DeleteIndexResponse, error) {
	log.Printf("info: delete index index_name=%s\n", r.IndexName)

	igs.mutex.Lock()
	defer igs.mutex.Unlock()

	var indexPath = path.Join(igs.dataDir, r.IndexName)
	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		_, err = os.Stat(indexPath)
		if err == nil {
			index := igs.indices[r.IndexName]

			err = index.Close()
			if err != nil {
				log.Printf("error: failed to close index (%s) index_name=%s\n", err.Error(), r.IndexName)
			}
			err = os.RemoveAll(path.Join(igs.dataDir, r.IndexName))
			if err != nil {
				log.Printf("error: failed to delete index (%s) index_name=%s\n", err.Error(), r.IndexName)
			}
		} else {
			log.Printf("error: index directory does not exist (%s) index_name=%s\n", err.Error(), r.IndexName)
		}

		delete(igs.indices, r.IndexName)
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s) index_name=%s\n", err.Error(), r.IndexName)
	}

	return &proto.DeleteIndexResponse{Result: ""}, nil
}

func (igs *indigoGRPCService) GetMapping(ctx context.Context, r *proto.GetMappingRequest) (*proto.GetMappingResponse, error) {
	log.Printf("info: get index mapping index_name=%s\n", r.IndexName)

	var bytesResp []byte
	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		index := igs.indices[r.IndexName]

		bytesResp, err = json.Marshal(index.Mapping())
		if err != nil {
			log.Printf("error: failed to create index mapping (%s)\n", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.GetMappingResponse{Mapping: string(bytesResp)}, nil
}

func (igs *indigoGRPCService) IndexDocument(ctx context.Context, r *proto.IndexDocumentRequest) (*proto.IndexDocumentResponse, error) {
	log.Printf("info: index document index_name=%s document_id=%s\n", r.IndexName, r.DocumentID)

	var document interface{}
	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		err = json.Unmarshal([]byte(r.Document), &document)
		if err == nil {
			index := igs.indices[r.IndexName]

			err = index.Index(r.DocumentID, document)
			if err != nil {
				log.Printf("error: failed to index document (%s) index_name=%s document_id=%s\n", err.Error(), r.IndexName, r.DocumentID)
			}
		} else {
			log.Printf("error: failed to create document (%s) index_name=%s document_id=%s\n", err.Error(), r.IndexName, r.DocumentID)
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s) index_name=%s document_id=%s\n", err.Error(), r.IndexName, r.DocumentID)
	}

	return &proto.IndexDocumentResponse{Result: ""}, nil
}

func (igs *indigoGRPCService) DeleteDocument(ctx context.Context, r *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {
	log.Printf("info: delete document IndexName=%s DocumentID=%s\n", r.IndexName, r.DocumentID)

	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		index := igs.indices[r.IndexName]

		err = index.Delete(r.DocumentID)
		if err != nil {
			log.Printf("error: failed to delete document (%s)\n", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.DeleteDocumentResponse{Result: ""}, nil
}

type IndexResult struct {
	DocumentCount int `json:"document_count"`
}

func (igs *indigoGRPCService) IndexDocuments(ctx context.Context, r *proto.IndexDocumentsRequest) (*proto.IndexDocumentsResponse, error) {
	log.Printf("info: index documents in bulk IndexName=%s\n", r.IndexName)

	var documents interface{}
	var err error
	var cnt int

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		index := igs.indices[r.IndexName]

		err = json.Unmarshal([]byte(r.Documents), &documents)
		if err == nil {
			/*
			 * create batch
			 */
			batch := index.NewBatch()
			for ID, doc := range documents.(map[string]interface{}) {
				err = batch.Index(ID, doc)
				if err != nil {
					log.Printf("error: failed to index document IndexName=%s DocumentID=%s (%s)\n", r.IndexName, ID, err.Error())
				}

				cnt++

				if cnt%int(r.BatchSize) == 0 {
					err = index.Batch(batch)
					if err != nil {
						log.Printf("error: failed to index documents in bulk IndexName=%s (%s)\n", r.IndexName, err.Error())
					}

					/*
					 * recreate batch
					 */
					batch = index.NewBatch()
				}
			}

			if batch.Size() > 0 {
				/*
				 * index document
				 */
				err = index.Batch(batch)
				if err != nil {
					log.Printf("error: failed to index documents in bulk IndexName=%s (%s)\n", r.IndexName, err.Error())
				}
			}
		} else {
			log.Printf("error: failed to create documents %s", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.IndexDocumentsResponse{Result: ""}, nil
}

func (igs *indigoGRPCService) DeleteDocuments(ctx context.Context, r *proto.DeleteDocumentsRequest) (*proto.DeleteDocumentsResponse, error) {
	log.Printf("info: delete documents in bulk IndexName=%s\n", r.IndexName)

	var ids []string
	var err error
	var cnt int

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		index := igs.indices[r.IndexName]

		err = json.Unmarshal([]byte(r.Ids), &ids)
		if err == nil {
			/*
			 * create batch
			 */
			batch := index.NewBatch()
			for i := range ids {
				batch.Delete(ids[i])

				cnt++

				if cnt%int(r.BatchSize) == 0 {
					err = index.Batch(batch)
					if err != nil {
						log.Printf("error: failed to delete documents in bulk IndexName=%s (%s)\n", r.IndexName, err.Error())
					}

					/*
					 * recreate batch
					 */
					batch = index.NewBatch()
				}
			}

			if batch.Size() > 0 {
				/*
				 * index document
				 */
				err = index.Batch(batch)
				if err != nil {
					log.Printf("error: failed to delete documents in bulk IndexName=%s (%s)\n", r.IndexName, err.Error())
				}
			}
		} else {
			log.Printf("error: failed to create documents %s", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.DeleteDocumentsResponse{Result: ""}, nil
}

func (igs *indigoGRPCService) SearchDocuments(ctx context.Context, r *proto.SearchDocumentsRequest) (*proto.SearchDocumentsResponse, error) {
	log.Printf("info: search documents IndexName=%s\n", r.IndexName)

	var bytesResp []byte
	var err error

	_, ok := igs.indices[r.IndexName]
	if ok == true {
		index := igs.indices[r.IndexName]

		searchRequest := bleve.NewSearchRequest(nil)
		err = json.Unmarshal([]byte(r.SearchRequest), searchRequest)
		if err == nil {
			searchResult, err := index.Search(searchRequest)
			if err == nil {
				/*
				 * create response
				 */
				bytesResp, err = json.Marshal(&searchResult)
				if err != nil {
					log.Printf("error: %s", err.Error())
				}
			} else {
				log.Printf("error: failed to search documents (%s)\n", err.Error())
			}
		} else {
			log.Printf("error: failed to create search request (%s)\n", err.Error())
		}
	} else {
		err = errors.New(fmt.Sprintf("%s does not exist", r.IndexName))
		log.Printf("error: index name does not exist (%s)\n", err.Error())
	}

	return &proto.SearchDocumentsResponse{SearchResult: string(bytesResp)}, nil
}
