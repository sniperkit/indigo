package handler

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"net/http"
)

type ListIndexHandler struct {
	client proto.IndigoClient
}

func NewListIndicesHandler(client proto.IndigoClient) *ListIndexHandler {
	return &ListIndexHandler{
		client: client,
	}
}

func (h *ListIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	resp, err := h.client.ListIndex(context.Background(), &proto.ListIndexRequest{})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to list index")

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
