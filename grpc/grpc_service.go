package grpc

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve"
	_ "github.com/mosuka/indigo/config"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

type indigoGRPCService struct {
	mappingConfig *viper.Viper
	index         bleve.Index
}

func NewIndigoGRPCService(indexDir string, indexMappingFile string, indexType string, indexStore string) *indigoGRPCService {
	var err error
	var mappingConfig = viper.New()
	var mappingBytes []byte
	var index bleve.Index

	/*
	 * load index mapping from mapping.json
	 */
	mappingConfig.SetConfigName("mapping")
	mappingConfig.SetConfigType("json")
	mappingBytes, err = ioutil.ReadFile(indexMappingFile)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	mappingConfig.ReadConfig(bytes.NewBuffer(mappingBytes))

	/*
	 * create index mapping
	 */
	var indexMapping = bleve.NewIndexMapping()
	err = json.Unmarshal(mappingBytes, indexMapping)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}

	/*
	 * open or create index
	 */
	_, err = os.Stat(indexDir)
	if err == nil {
		index, err = bleve.OpenUsing(indexDir, map[string]interface{}{})
	} else {
		index, err = bleve.NewUsing(indexDir, indexMapping, indexType, indexStore, nil)
	}
	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	return &indigoGRPCService{
		mappingConfig: mappingConfig,
		index:         index,
	}
}

func (igs *indigoGRPCService) Mapping(ctx context.Context, r *proto.MappingRequest) (*proto.MappingResponse, error) {
	log.Print("info: Mapping")

	/*
	 * create response
	 */
	bytesResponse, err := json.Marshal(igs.index.Mapping())
	if err != nil {
		log.Printf("error: %s", err.Error())
		return &proto.MappingResponse{Mapping: ""}, nil
	}

	/*
	 * return response
	 */
	return &proto.MappingResponse{Mapping: string(bytesResponse)}, nil
}

type IndexResult struct {
	DocumentCount int `json:"document_count"`
}

func (igs *indigoGRPCService) Index(ctx context.Context, r *proto.IndexRequest) (*proto.IndexResponse, error) {
	log.Print("info: Index")

	indexResult := IndexResult{DocumentCount: 0}

	/*
	 * create documents
	 */
	var documents interface{}
	err := json.Unmarshal([]byte(r.Documents), &documents)
	if err != nil {
		log.Printf("error: a %s", err.Error())

		bytesResponse, err := json.Marshal(&indexResult)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}

		return &proto.IndexResponse{Result: string(bytesResponse)}, nil
	}

	/*
	 * create batch
	 */
	batch := igs.index.NewBatch()
	for ID, doc := range documents.(map[string]interface{}) {
		err = batch.Index(ID, doc)
		if err != nil {
			log.Printf("error: %s", err.Error())

			bytesResponse, err := json.Marshal(&indexResult)
			if err != nil {
				log.Printf("error: %s", err.Error())
			}

			return &proto.IndexResponse{Result: string(bytesResponse)}, nil
		}

		indexResult.DocumentCount++

		if indexResult.DocumentCount%int(r.BatchSize) == 0 {
			err = igs.index.Batch(batch)
			if err != nil {
				log.Printf("error: %s", err.Error())

				bytesResponse, err := json.Marshal(&indexResult)
				if err != nil {
					log.Printf("error: %s", err.Error())
				}

				return &proto.IndexResponse{Result: string(bytesResponse)}, nil
			}

			/*
			 * recreate batch
			 */
			batch = igs.index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		/*
		 * index document
		 */
		err = igs.index.Batch(batch)
		if err != nil {
			log.Printf("error: %s", err.Error())

			bytesResponse, err := json.Marshal(&indexResult)
			if err != nil {
				log.Printf("error: %s", err.Error())
			}

			return &proto.IndexResponse{Result: string(bytesResponse)}, nil
		}
	}

	/*
	 * return response
	 */
	log.Printf("info: %d document(s) indexed", indexResult.DocumentCount)

	bytesResponse, err := json.Marshal(indexResult)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	return &proto.IndexResponse{Result: string(bytesResponse)}, nil
}

func (igs *indigoGRPCService) Delete(ctx context.Context, r *proto.DeleteRequest) (*proto.IndexResponse, error) {
	log.Print("info: Delete")

	indexResult := IndexResult{DocumentCount: 0}

	/*
	 * create document id list
	 */
	var ids []string
	err := json.Unmarshal([]byte(r.Ids), &ids)
	if err != nil {
		log.Printf("error: a %s", err.Error())

		bytesResponse, err := json.Marshal(&indexResult)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}

		return &proto.IndexResponse{Result: string(bytesResponse)}, nil
	}

	/*
	 * create batch
	 */
	batch := igs.index.NewBatch()
	for i := range ids {
		batch.Delete(ids[i])
		if err != nil {
			log.Printf("error: %s", err.Error())

			bytesResponse, err := json.Marshal(&indexResult)
			if err != nil {
				log.Printf("error: %s", err.Error())
			}

			return &proto.IndexResponse{Result: string(bytesResponse)}, nil
		}

		indexResult.DocumentCount++

		if indexResult.DocumentCount%int(r.BatchSize) == 0 {
			/*
			 * delete document
			 */
			err = igs.index.Batch(batch)
			if err != nil {
				log.Printf("error: %s", err.Error())

				bytesResponse, err := json.Marshal(&indexResult)
				if err != nil {
					log.Printf("error: %s", err.Error())
				}

				return &proto.IndexResponse{Result: string(bytesResponse)}, nil
			}

			/*
			 * recreate batch
			 */
			batch = igs.index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		/*
		 * delete document
		 */
		err = igs.index.Batch(batch)
		if err != nil {
			log.Printf("error: %s", err.Error())

			bytesResponse, err := json.Marshal(&indexResult)
			if err != nil {
				log.Printf("error: %s", err.Error())
			}

			return &proto.IndexResponse{Result: string(bytesResponse)}, nil
		}
	}

	/*
	 * return response
	 */
	log.Printf("info: %d document(s) deleted", indexResult.DocumentCount)

	bytesResponse, err := json.Marshal(&indexResult)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	return &proto.IndexResponse{Result: string(bytesResponse)}, nil
}

func (igs *indigoGRPCService) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	log.Print("info: Search")

	var bytesResponse []byte
	var err error

	/*
	 * create search request
	 */
	searchRequest := bleve.NewSearchRequest(nil)
	err = json.Unmarshal([]byte(r.Request), searchRequest)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return &proto.SearchResponse{Result: ""}, nil
	}

	/*
	 * create search result
	 */
	searchResult, err := igs.index.Search(searchRequest)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return &proto.SearchResponse{Result: ""}, nil
	}

	/*
	 * create response
	 */
	bytesResponse, err = json.Marshal(&searchResult)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return &proto.SearchResponse{Result: ""}, nil
	}

	/*
	 * return response
	 */
	return &proto.SearchResponse{Result: string(bytesResponse)}, nil
}
