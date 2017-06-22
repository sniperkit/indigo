//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"bytes"
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type SearchHandler struct {
	client proto.IndigoClient
}

func NewSearchHandler(client proto.IndigoClient) *SearchHandler {
	return &SearchHandler{
		client: client,
	}
}

type SearchResponse struct {
	SearchResult map[string]interface{} `json:"search_result"`
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	searchRequest := bleve.NewSearchRequest(nil)
	if len(resourceBytes) > 0 {
		err := searchRequest.UnmarshalJSON(resourceBytes)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create search request")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if req.URL.Query().Get("query") != "" {
		searchRequest.Query = bleve.NewQueryStringQuery(req.URL.Query().Get("query"))
	}

	if req.URL.Query().Get("size") != "" {
		i, err := strconv.Atoi(req.URL.Query().Get("size"))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to convert size")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		searchRequest.Size = i
	}
	if searchRequest.Size == 0 {
		searchRequest.Size = DefaultSize
	}

	if req.URL.Query().Get("from") != "" {
		i, err := strconv.Atoi(req.URL.Query().Get("from"))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to convert from")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		searchRequest.From = i
	}
	if searchRequest.From < 0 {
		searchRequest.From = DefaultFrom
	}

	if req.URL.Query().Get("explain") != "" {
		if req.URL.Query().Get("explain") == "true" {
			searchRequest.Explain = true
		} else {
			searchRequest.Explain = false
		}
	}

	if req.URL.Query().Get("fields") != "" {
		searchRequest.Fields = strings.Split(req.URL.Query().Get("fields"), ",")
	}

	if req.URL.Query().Get("sort") != "" {
		searchRequest.SortBy(strings.Split(req.URL.Query().Get("sort"), ","))
	}

	if req.URL.Query().Get("facets") != "" {
		facetRequest := bleve.FacetsRequest{}
		err := json.Unmarshal([]byte(req.URL.Query().Get("facets")), &facetRequest)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create facets")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		searchRequest.Facets = facetRequest
	}

	if req.URL.Query().Get("highlight") != "" {
		highlightRequest := bleve.NewHighlight()
		err := json.Unmarshal([]byte(req.URL.Query().Get("highlight")), highlightRequest)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create highlight")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		searchRequest.Highlight = highlightRequest
	}

	if req.URL.Query().Get("highlightStyle") != "" || req.URL.Query().Get("highlightField") != "" {
		highlightRequest := bleve.NewHighlightWithStyle(req.URL.Query().Get("highlightStyle"))
		highlightRequest.Fields = strings.Split(req.URL.Query().Get("highlightField"), ",")
		searchRequest.Highlight = highlightRequest
	}

	if req.URL.Query().Get("include-locations") != "" {
		if req.URL.Query().Get("include-locations") == "true" {
			searchRequest.IncludeLocations = true
		} else {
			searchRequest.IncludeLocations = false
		}
	}

	searchRequestAny, err := util.MarshalAny(searchRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create search request")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoSearchRequest := &proto.SearchRequest{
		SearchRequest: &searchRequestAny,
	}

	resp, err := h.client.Search(context.Background(), protoSearchRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to search documents")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	searchResult, err := util.UnmarshalAny(resp.SearchResult)

	r := util.SearchResponse{
		SearchResult: searchResult.(*bleve.SearchResult),
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