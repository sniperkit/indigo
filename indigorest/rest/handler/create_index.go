package handler

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve/mapping"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/service"
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

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createIndexRequest := service.CreateIndexRequest{}
	err = json.Unmarshal(resourceBytes, &createIndexRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create CreateIndexRequest")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createIndexRequest.Index = index

	if req.URL.Query().Get("indexMapping") != "" {
		indexMapping := &mapping.IndexMappingImpl{}
		err := json.Unmarshal([]byte(req.URL.Query().Get("indexMapping")), indexMapping)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create indexMapping")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createIndexRequest.IndexMapping = indexMapping
	}

	if req.URL.Query().Get("indexType") != "" {
		createIndexRequest.IndexType = req.URL.Query().Get("indexType")
	}

	if req.URL.Query().Get("kvstore") != "" {
		createIndexRequest.Kvstore = req.URL.Query().Get("kvstore")
	}

	if req.URL.Query().Get("kvconfig") != "" {
		kvconfig := make(map[string]interface{})
		err := json.Unmarshal([]byte(req.URL.Query().Get("kvconfig")), kvconfig)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create kvconfig")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createIndexRequest.Kvconfig = kvconfig
	}

	protoCreateIndexRequest, err := createIndexRequest.ProtoMessage()
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
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
