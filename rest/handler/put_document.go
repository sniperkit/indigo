package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
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
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)

	indexName := vars["indexName"]
	id := vars["id"]

	document, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	response := make(map[string]interface{})

	resp, err := h.client.PutDocument(context.Background(), &proto.PutDocumentRequest{Name: indexName, Id: id, Document: document})
	if err == nil {
		log.Print("info: request to Indigo gRPC Server was successful\n")

		w.WriteHeader(http.StatusOK)
		response["count"] = resp.Count
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
