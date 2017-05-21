package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

type OpenIndexHandler struct {
	client proto.IndigoClient
}

func NewOpenIndexHandler(client proto.IndigoClient) *OpenIndexHandler {
	return &OpenIndexHandler{
		client: client,
	}
}

type OpenIndexResource struct {
	RuntimeConfig map[string]interface{} `json:"runtime_config,omitempty"`
}

func (h *OpenIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	openIndexResource := OpenIndexResource{}
	err = json.Unmarshal(resourceBytes, &openIndexResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create open index resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	runtimeConfigBytes, err := json.Marshal(openIndexResource.RuntimeConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create runtime config")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoOpenIndexRequest := &proto.OpenIndexRequest{
		Index:         index,
		RuntimeConfig: runtimeConfigBytes,
	}

	resp, err := h.client.OpenIndex(context.Background(), protoOpenIndexRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("faild to open index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("faild to create response")

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
