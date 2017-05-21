package handler

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve/mapping"
	"github.com/gorilla/mux"
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

type CreateIndexResource struct {
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     map[string]interface{}    `json:"kvconfig,omitempty"`
}

func (h *CreateIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)
	index := vars["index"]

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createIndexResource := CreateIndexResource{}
	err = json.Unmarshal(resourceBytes, &createIndexResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexMappingBytes, err := json.Marshal(createIndexResource.IndexMapping)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create index mapping")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kvconfigBytes, err := json.Marshal(createIndexResource.Kvconfig)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create kvconfig")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoCreateIndexRequest := &proto.CreateIndexRequest{
		Index:        index,
		IndexMapping: indexMappingBytes,
		IndexType:    createIndexResource.IndexType,
		Kvstore:      createIndexResource.Kvstore,
		Kvconfig:     kvconfigBytes,
	}

	if req.URL.Query().Get("indexType") == "" {
		protoCreateIndexRequest.IndexType = req.URL.Query().Get("indexType")
	}

	if protoCreateIndexRequest.IndexType == "" {
		protoCreateIndexRequest.IndexType = DefaultIndexType
	}

	if req.URL.Query().Get("kvstore") == "" {
		protoCreateIndexRequest.Kvstore = req.URL.Query().Get("kvstore")
	}

	if protoCreateIndexRequest.Kvstore == "" {
		protoCreateIndexRequest.Kvstore = DefaultKvstore
	}

	resp, err := h.client.CreateIndex(context.Background(), protoCreateIndexRequest)
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
