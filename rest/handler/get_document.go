package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

type GetDocumentHandler struct {
	client proto.IndigoClient
}

func NewGetDocumentHandler(client proto.IndigoClient) *GetDocumentHandler {
	return &GetDocumentHandler{
		client: client,
	}
}

func (h *GetDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	indexName := vars["indexName"]
	id := vars["id"]

	resp, err := h.client.GetDocument(context.Background(), &proto.GetDocumentRequest{IndexName: indexName, Id: id})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to get document")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	fields := make(map[string]interface{})
	if err := json.Unmarshal(resp.Fields, &fields); err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	r := struct {
		ID     string                 `json:"id"`
		Fields map[string]interface{} `json:"fields"`
	}{
		ID:     resp.Id,
		Fields: fields,
	}

	output, err := json.MarshalIndent(r, "", "  ")
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
