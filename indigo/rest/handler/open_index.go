package handler

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
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

func (h *OpenIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]

	runtimeConfig, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("faild to create runtime config")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.OpenIndex(context.Background(), &proto.OpenIndexRequest{Index: index, RuntimeConfig: runtimeConfig})
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
