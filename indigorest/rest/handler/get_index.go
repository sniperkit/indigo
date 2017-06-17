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
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/resource"
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

	protoGetIndexRequest := &proto.GetIndexRequest{}

	if req.URL.Query().Get("includeIndexMapping") == "true" {
		protoGetIndexRequest.IncludeIndexMapping = true
	}

	if req.URL.Query().Get("includeIndexType") == "true" {
		protoGetIndexRequest.IncludeIndexType = true
	}

	if req.URL.Query().Get("includeKvstore") == "true" {
		protoGetIndexRequest.IncludeKvstore = true
	}

	if req.URL.Query().Get("includeKvconfig") == "true" {
		protoGetIndexRequest.IncludeKvconfig = true
	}

	resp, err := h.client.GetIndex(context.Background(), protoGetIndexRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	r := resource.GetIndexResponse{
		Path:      resp.Path,
		IndexType: resp.IndexType,
		Kvstore:   resp.Kvstore,
	}

	if req.URL.Query().Get("includeIndexMapping") == "true" {
		indexMapping, err := proto.UnmarshalAny(resp.IndexMapping)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create index mapping")

			Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		r.IndexMapping = indexMapping.(*mapping.IndexMappingImpl)
	}

	if req.URL.Query().Get("includeKvconfig") == "true" {
		kvconfig, err := proto.UnmarshalAny(resp.Kvconfig)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create kvconfig")

			Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		r.Kvconfig = kvconfig
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
