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

type DeleteIndexHandler struct {
	client proto.IndigoClient
}

func NewDeleteIndexHandler(client proto.IndigoClient) *DeleteIndexHandler {
	return &DeleteIndexHandler{
		client: client,
	}
}

func (h *DeleteIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	resp, err := h.client.Delete(context.Background(), &proto.DeleteRequest{Ids: string(requestBody), BatchSize: int32(batchSize)})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", resp.Result)
}
