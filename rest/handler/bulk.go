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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	batchSize := constant.DefaultBatchSize
	bs, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
	if err == nil {
		if bs > 0 {
			batchSize = int32(bs)
		} else {
			log.Printf("warn: unexpected batchSize (%d)\n", bs)
		}
	} else {
		log.Printf("warn: failed to convert batchSize to int (%s)\n", err.Error())
	}

	response := make(map[string]interface{})

	resp, err := h.client.Bulk(context.Background(), &proto.BulkRequest{Name: indexName, BulkRequest: bulkRequest, BatchSize: batchSize})
	if err == nil {
		log.Print("info: request to Indigo gRPC Server was successful\n")

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
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("error: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(bytesResponse))

	fmt.Fprintf(w, "%s\n", buf.String())

	return
}
