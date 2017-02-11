package handler

import (
	"fmt"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type GetMappingHandler struct {
	client proto.IndigoClient
}

func NewGetMappingHandler(client proto.IndigoClient) *GetMappingHandler {
	return &GetMappingHandler{
		client: client,
	}
}

func (h *GetMappingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: request_uri=%s user_agent=%s\n", req.RequestURI, req.UserAgent())

	resp, err := h.client.Mapping(context.Background(), &proto.MappingRequest{})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", resp.Mapping)
}
