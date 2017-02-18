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

type GetStatsHandler struct {
	client proto.IndigoClient
}

func NewGetStatsHandler(client proto.IndigoClient) *GetStatsHandler {
	return &GetStatsHandler{
		client: client,
	}
}

func (h *GetStatsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)
	indexName := vars["indexName"]

	response := make(map[string]interface{})

	resp, err := h.client.GetStats(context.Background(), &proto.GetStatsRequest{Name: indexName})
	if err == nil {
		log.Print("info: request to Indigo gRPC Server was successful\n")

		indexStats := make(map[string]interface{})

		err = json.Unmarshal(resp.Stats, &indexStats)
		if err == nil {
			log.Print("info: index stats created\n")

			w.WriteHeader(http.StatusOK)
			response["stats"] = indexStats
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
