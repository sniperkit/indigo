package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/defaultvalue"
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

func (h *BulkHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]

	bulkRequest, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("faild to create bulk request")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	batchSize := defaultvalue.DefaultBatchSize
	bs, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
	if err == nil {
		if bs > 0 {
			batchSize = int32(bs)
		} else {
			log.WithFields(log.Fields{
				"batchSize": bs,
				"err":       err,
			}).Warn("unexpected batch size")
		}
	} else {
		log.WithFields(log.Fields{
			"batchSize": bs,
			"err":       err,
		}).Warn("unexpected batch size")
	}

	resp, err := h.client.Bulk(context.Background(), &proto.BulkRequest{Index: index, BulkRequests: bulkRequest, BatchSize: batchSize})
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
