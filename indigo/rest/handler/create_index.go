package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
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
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]

	indexMapping, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("faild to create index mapping")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexType := req.URL.Query().Get("indexType")
	if indexType == "" {
		indexType = defaultvalue.DefaultIndexType
	}

	indexStore := req.URL.Query().Get("indexStore")
	if indexStore == "" {
		indexStore = defaultvalue.DefaultKVStore
	}

	resp, err := h.client.CreateIndex(context.Background(), &proto.CreateIndexRequest{Index: index, IndexMapping: indexMapping, IndexType: indexType, Kvstore: indexStore, Kvconfig: nil})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

	return
}
