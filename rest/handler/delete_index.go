package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type DeleteIndexHandler struct {
	client proto.IndigoClient
}

func NewDeleteIndexHandler(client proto.IndigoClient) *DeleteIndexHandler {
	return &DeleteIndexHandler{
		client: client,
	}
}

func (h *DeleteIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)

	indexName := vars["indexName"]

	response := make(map[string]interface{})

	resp, err := h.client.DeleteIndex(context.Background(), &proto.DeleteIndexRequest{IndexName: indexName})
	if err == nil {
		log.Print("debug: request to Indigo gRPC Server\n")

		w.WriteHeader(http.StatusOK)
		response["index_name"] = resp.IndexName
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
