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

type PutDocumentHandler struct {
	client proto.IndigoClient
}

func NewPutDocumentHandler(client proto.IndigoClient) *PutDocumentHandler {
	return &PutDocumentHandler{
		client: client,
	}
}

func (h *PutDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]
	id := vars["id"]

	fields, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create document fields")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.PutDocument(context.Background(), &proto.PutDocumentRequest{Index: index, Id: id, Fields: fields})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to put document")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
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
