package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
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
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)

	indexName := vars["indexName"]

	bulkRequest, err := ioutil.ReadAll(req.Body)
	if err == nil {
		log.Print("debug: read request body")
	} else {
		log.Printf("error: failed to read request body (%s)\n", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	batchSize := constant.DefaultBatchSize
	bs, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
	if err == nil {
		log.Printf("debug: convert string to int batchSize=\"%s\"\n", batchSize)
		if bs > 0 {
			log.Printf("debug: batch size batchSize=%d\n", bs)
			batchSize = int32(bs)
		} else {
			log.Printf("warn: unexpected batch size batchSize=%d\n", bs)
		}
	} else {
		log.Printf("warn: failed to convert string to int (%s) batchSize=\"%s\"\n", err.Error(), batchSize)
	}

	resp, err := h.client.Bulk(context.Background(), &proto.BulkRequest{IndexName: indexName, BulkRequest: bulkRequest, BatchSize: batchSize})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in creating response JSON\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", buf.String())

	return
}
