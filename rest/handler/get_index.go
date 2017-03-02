package handler

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type GetIndexHandler struct {
	client proto.IndigoClient
}

func NewGetIndexHandler(client proto.IndigoClient) *GetIndexHandler {
	return &GetIndexHandler{
		client: client,
	}
}

func (h *GetIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)
	indexName := vars["indexName"]

	resp, err := h.client.GetIndex(context.Background(), &proto.GetIndexRequest{IndexName: indexName})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	indexStats := make(map[string]interface{})
	if err := json.Unmarshal(resp.IndexStats, &indexStats); err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	indexMapping := bleve.NewIndexMapping()
	if err := json.Unmarshal(resp.IndexMapping, &indexMapping); err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	r := struct {
		DocumentCount uint64                    `json:"document_count"`
		IndexStats    map[string]interface{}    `json:"index_stats"`
		IndexMapping  *mapping.IndexMappingImpl `json:"index_mapping"`
	}{
		DocumentCount: resp.DocumentCount,
		IndexStats:    indexStats,
		IndexMapping:  indexMapping,
	}

	output, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in creating response JSON\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

	return
}
