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

type PutDocumentHandler struct {
	client proto.IndigoClient
}

func NewPutDocumentHandler(client proto.IndigoClient) *PutDocumentHandler {
	return &PutDocumentHandler{
		client: client,
	}
}

type PutDocumentResource struct {
	Fields interface{} `json:"fields,omitempty"`
}

func (h *PutDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)
	index := vars["index"]
	id := vars["id"]

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	putDocumentResource := PutDocumentResource{}
	err = json.Unmarshal(resourceBytes, &putDocumentResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create put document resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fieldsBytes, err := json.Marshal(putDocumentResource.Fields)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create fields")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoPutDocumentRequest := &proto.PutDocumentRequest{
		Index:  index,
		Id:     id,
		Fields: fieldsBytes,
	}

	resp, err := h.client.PutDocument(context.Background(), protoPutDocumentRequest)
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
