package handler

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

type GetIndexHandler struct {
	client proto.IndigoClient
}

func NewGetIndexHandler(client proto.IndigoClient) *GetIndexHandler {
	return &GetIndexHandler{
		client: client,
	}
}

func (h *GetIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)
	index := vars["index"]

	resp, err := h.client.GetIndex(context.Background(), &proto.GetIndexRequest{Index: index})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	indexStats := make(map[string]interface{})
	if err := json.Unmarshal(resp.IndexStats, &indexStats); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create index stats")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	indexMapping := bleve.NewIndexMapping()
	if err := json.Unmarshal(resp.IndexMapping, &indexMapping); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create index mapping")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	r := struct {
		DocumentCount uint64                    `json:"document_count"`
		IndexStats    map[string]interface{}    `json:"index_stats"`
		IndexMapping  *mapping.IndexMappingImpl `json:"index_mapping"`
	}{
		DocumentCount: resp.DocumentCount,
		IndexStats:    indexStats,
		IndexMapping:  indexMapping,
	}

	output, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
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
