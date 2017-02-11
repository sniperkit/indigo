package handler

import (
	"fmt"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PostIndexHandler struct {
	client proto.IndigoClient
}

func NewPostIndexHandler(client proto.IndigoClient) *PostIndexHandler {
	return &PostIndexHandler{
		client: client,
	}
}

func (h *PostIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: request_uri=%s user_agent=%s\n", req.RequestURI, req.UserAgent())

	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	batchSize, err := strconv.Atoi(req.Form.Get("batchSize"))
	if err != nil {
		log.Printf("warn: %s", err.Error())
		batchSize = 1000
	}

	resp, err := h.client.Index(context.Background(), &proto.IndexRequest{Documents: string(requestBody), BatchSize: int32(batchSize)})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", resp.Result)
}
