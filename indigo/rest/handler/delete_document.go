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

type DeleteDocumentHandler struct {
	client proto.IndigoClient
}

func NewDeleteDocumentHandler(client proto.IndigoClient) *DeleteDocumentHandler {
	return &DeleteDocumentHandler{
		client: client,
	}
}

func (h *DeleteDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]
	id := vars["id"]

	resp, err := h.client.DeleteDocument(context.Background(), &proto.DeleteDocumentRequest{Index: index, Id: id})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to delete document")

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
