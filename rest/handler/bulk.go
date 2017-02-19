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

	response := make(map[string]interface{})

	resp, err := h.client.Bulk(context.Background(), &proto.BulkRequest{IndexName: indexName, BulkRequest: bulkRequest, BatchSize: batchSize})
	if err == nil {
		log.Print("debug: request to the Indigo gRPC Server\n")

		w.WriteHeader(http.StatusOK)
		response["put_count"] = resp.PutCount
		response["put_error_count"] = resp.PutErrorCount
		response["delete_count"] = resp.DeleteCount
	} else {
		log.Printf("error: failed to request to the Indigo gRPC Server (%s)\n", err.Error())

		w.WriteHeader(http.StatusServiceUnavailable)
		response["error"] = err.Error()
	}

	bytesResponse, err := json.Marshal(response)
	if err == nil {
		log.Print("debug: create response\n")
	} else {
		log.Printf("error: failed to create response (%s)\n", err.Error())

		w.WriteHeader(http.StatusServiceUnavailable)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(bytesResponse))

	fmt.Fprintf(w, "%s\n", buf.String())

	return
}
