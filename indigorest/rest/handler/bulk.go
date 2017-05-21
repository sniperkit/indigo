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
	"strconv"
)

type BulkHandler struct {
	client proto.IndigoClient
}

func NewBulkHandler(client proto.IndigoClient) *BulkHandler {
	return &BulkHandler{
		client: client,
	}
}

type BulkRequest struct {
	Method string      `json:"method,omitempty"`
	Id     string      `json:"id,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

type BulkResource struct {
	BatchSize    int32         `json:"batch_size,omitempty"`
	BulkRequests []BulkRequest `json:"bulk_requests,omitempty"`
}

func (h *BulkHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	bulkResource := BulkResource{}
	err = json.Unmarshal(resourceBytes, &bulkResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bulkRequestsBytes, err := json.Marshal(bulkResource.BulkRequests)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk requests")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoBulkRequest := &proto.BulkRequest{
		Index:        index,
		BatchSize:    bulkResource.BatchSize,
		BulkRequests: bulkRequestsBytes,
	}

	if req.URL.Query().Get("batchSize") != "" {
		i, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to convert batch size")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		protoBulkRequest.BatchSize = int32(i)
	}

	if protoBulkRequest.BatchSize == 0 {
		protoBulkRequest.BatchSize = DefaultBatchSize
	}

	resp, err := h.client.Bulk(context.Background(), protoBulkRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to index documents in bulk")

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
