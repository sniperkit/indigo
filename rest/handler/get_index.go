package handler

import (
	"bytes"
	"encoding/json"
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

	result := make(map[string]interface{})

	result["document_count"] = resp.DocumentCount

	indexStats := make(map[string]interface{})
	if err := json.Unmarshal(resp.IndexStats, &indexStats); err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	result["index_stats"] = indexStats

	indexMapping := make(map[string]interface{})
	if err := json.Unmarshal(resp.IndexMapping, &indexMapping); err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	result["index_mapping"] = indexMapping

	output, err := json.MarshalIndent(result, "", "  ")
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
