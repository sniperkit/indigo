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

type SearchHandler struct {
	client proto.IndigoClient
}

func NewSearchHandler(client proto.IndigoClient) *SearchHandler {
	return &SearchHandler{
		client: client,
	}
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("info: host=\"%s\" request_uri=\"%s\" method=\"%s\" remote_addr=\"%s\" user_agent=\"%s\"\n", req.Host, req.RequestURI, req.Method, req.RemoteAddr, req.UserAgent())

	vars := mux.Vars(req)

	indexName := vars["indexName"]

	searchRequest, err := ioutil.ReadAll(req.Body)
	if err == nil {
		log.Print("debug: read request body")
	} else {
		log.Printf("error: failed to read request body (%s)\n", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.client.Search(context.Background(), &proto.SearchRequest{IndexName: indexName, SearchRequest: searchRequest})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Print("debug: succeeded in requesting to the Indigo gRPC Server\n")

	result := make(map[string]interface{})

	searchResult := make(map[string]interface{})
	if err := json.Unmarshal(resp.SearchResult, &searchResult); err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	result["searchResult"] = searchResult

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
