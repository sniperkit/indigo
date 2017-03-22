package handler

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"golang.org/x/net/context"
	"io/ioutil"
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
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	index := vars["index"]

	searchRequest, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create search request")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.Search(context.Background(), &proto.SearchRequest{Index: index, SearchRequest: searchRequest})
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to search documents")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	searchResult := make(map[string]interface{})
	if err := json.Unmarshal(resp.SearchResult, &searchResult); err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create search result")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	r := struct {
		SearchResult map[string]interface{} `json:"search_result"`
	}{
		SearchResult: searchResult,
	}

	output, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(bytes.NewReader(output))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

	return
}