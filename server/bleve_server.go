package server

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve"
	_ "github.com/mosuka/bleve-server/config"
	"github.com/mosuka/bleve-server/proto"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

type bleveServer struct {
	mappingConfig *viper.Viper
	index         bleve.Index
}

func NewBleveServer(indexDir string, indexMappingFile string, indexType string, indexStore string) *bleveServer {
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

	return &bleveServer{
		mappingConfig: mappingConfig,
		index:         index,
	}
}

func (bs *bleveServer) Mapping(ctx context.Context, r *proto.MappingRequest) (*proto.MappingResponse, error) {
	log.Print("info: Mapping")

	/*
	 * create response
	 */
	bytesResponse, err := json.Marshal(bs.index.Mapping())
	if err != nil {
		log.Printf("error: %s", err.Error())
		return &proto.MappingResponse{Mapping: ""}, nil
	}

	/*
	 * return response
	 */
	return &proto.MappingResponse{Mapping: string(bytesResponse)}, nil
}

func (bs *bleveServer) Index(ctx context.Context, r *proto.IndexRequest) (*proto.IndexResponse, error) {
	log.Print("info: Index")

	documentCount := 0

	/*
	 * create documents
	 */
	var documents interface{}
	err := json.Unmarshal([]byte(r.Documents), &documents)
	if err != nil {
		log.Printf("error: a %s", err.Error())
		return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
	}

	/*
	 * create batch
	 */
	batch := bs.index.NewBatch()
	for ID, doc := range documents.(map[string]interface{}) {
		err = batch.Index(ID, doc)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
		}

		documentCount++

		if documentCount%int(r.BatchSize) == 0 {
			err = bs.index.Batch(batch)
			if err != nil {
				log.Printf("error: %s", err.Error())
				return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
			}

			/*
			 * recreate batch
			 */
			batch = bs.index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		/*
		 * index document
		 */
		err = bs.index.Batch(batch)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
		}
	}

	/*
	 * return response
	 */
	log.Printf("info: %d document(s) indexed", documentCount)
	return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
}

func (bs *bleveServer) Delete(ctx context.Context, r *proto.DeleteRequest) (*proto.IndexResponse, error) {
	log.Print("info: Delete")

	documentCount := 0

	/*
	 * create document id list
	 */
	var ids []string
	err := json.Unmarshal([]byte(r.Ids), &ids)
	if err != nil {
		log.Printf("error: a %s", err.Error())
		return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
	}

	/*
	 * create batch
	 */
	batch := bs.index.NewBatch()
	for i := range ids {
		batch.Delete(ids[i])
		if err != nil {
			log.Printf("error: %s", err.Error())
			return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
		}

		documentCount++

		if documentCount%int(r.BatchSize) == 0 {
			/*
			 * delete document
			 */
			err = bs.index.Batch(batch)
			if err != nil {
				log.Printf("error: %s", err.Error())
				return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
			}

			/*
			 * recreate batch
			 */
			batch = bs.index.NewBatch()
		}
	}

	if batch.Size() > 0 {
		/*
		 * delete document
		 */
		err = bs.index.Batch(batch)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
		}
	}

	/*
	 * return response
	 */
	log.Printf("info: %d document(s) deleted", documentCount)
	return &proto.IndexResponse{DocumentCount: int32(documentCount)}, nil
}

func (bs *bleveServer) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
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
	searchResult, err := bs.index.Search(searchRequest)
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
