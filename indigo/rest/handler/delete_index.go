package handler

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"net/http"
)

type DeleteIndexHandler struct {
	client proto.IndigoClient
}

func NewDeleteIndexHandler(client proto.IndigoClient) *DeleteIndexHandler {
	return &DeleteIndexHandler{
		client: client,
	}
}

func (h *DeleteIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]

	resp, err := h.client.DeleteIndex(context.Background(), &proto.DeleteIndexRequest{Index: index})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to delete index")

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
