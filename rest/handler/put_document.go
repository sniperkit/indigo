package handler

import (
	"bytes"
	"encoding/json"
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

	fields, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: indexName, Id: id, Fields: fields})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in creating response JSON\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

	return
}
