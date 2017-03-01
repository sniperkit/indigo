package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateIndexHandler struct {
	client proto.IndigoClient
}

func NewCreateIndexHandler(client proto.IndigoClient) *CreateIndexHandler {
	return &CreateIndexHandler{
		client: client,
	}
}

func (h *CreateIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)

	indexName := vars["indexName"]

	indexMapping, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexType := req.URL.Query().Get("indexType")
	if indexType == "" {
		indexType = constant.DefaultIndexType
	}

	indexStore := req.URL.Query().Get("indexStore")
	if indexStore == "" {
		indexStore = constant.DefaultKVStore
	}

	resp, err := h.client.CreateIndex(context.Background(), &proto.CreateIndexRequest{IndexName: indexName, IndexMapping: indexMapping, IndexType: indexType, Kvstore: indexStore, Kvconfig: nil})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	output, err := json.MarshalIndent(resp, "", "  ")
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
