package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
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
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)
	indexName := vars["indexName"]

	response := make(map[string]interface{})

	resp, err := h.client.GetMapping(context.Background(), &proto.GetMappingRequest{Name: indexName})
	if err == nil {
		log.Print("info: request to Indigo gRPC Server was successful\n")

		indexMapping := bleve.NewIndexMapping()

		err = json.Unmarshal(resp.Mapping, indexMapping)
		if err == nil {
			log.Print("info: index mapping created\n")

			w.WriteHeader(http.StatusOK)
			response["mapping"] = indexMapping
		} else {
			log.Printf("error: failed to create index mapping (%s)\n", err.Error())

			w.WriteHeader(http.StatusServiceUnavailable)
			response["error"] = err.Error()
		}
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
