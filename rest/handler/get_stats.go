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

	resp, err := h.client.GetStats(context.Background(), &proto.GetStatsRequest{IndexName: indexName})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	result := make(map[string]interface{})

	indexStats := make(map[string]interface{})
	if err := json.Unmarshal(resp.IndexStats, &indexStats); err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	result["indexStats"] = indexStats

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in creating response JSON\n")

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", buf.String())

	return
}
