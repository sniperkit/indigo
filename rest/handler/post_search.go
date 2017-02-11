package handler

import (
	"fmt"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
)

type PostSearchHandler struct {
	client proto.IndigoClient
}

func NewPostSearchHandler(client proto.IndigoClient) *PostSearchHandler {
	return &PostSearchHandler{
		client: client,
	}
}

func (h *PostSearchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: request_uri=%s user_agent=%s\n", req.RequestURI, req.UserAgent())

	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error: %s", err.Error())
		return
	}

	resp, err := h.client.Search(context.Background(), &proto.SearchRequest{Request: string(requestBody)})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", resp.Result)
}
