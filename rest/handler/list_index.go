package handler

import (
	"bytes"
	"encoding/json"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"log"
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
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	resp, err := h.client.ListIndex(context.Background(), &proto.ListIndexRequest{})
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
