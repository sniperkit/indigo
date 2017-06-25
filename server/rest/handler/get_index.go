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
	"encoding/json"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
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

	// create request
	getIndexRequest, err := util.NewGetIndexRequest(
		false,
		false,
		false,
		false,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create get index request")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// overwrite request
	if req.URL.Query().Get("includeIndexMapping") == "true" {
		getIndexRequest.IncludeIndexMapping = true
	}
	if req.URL.Query().Get("includeIndexType") == "true" {
		getIndexRequest.IncludeIndexType = true
	}
	if req.URL.Query().Get("includeKvstore") == "true" {
		getIndexRequest.IncludeKvstore = true
	}
	if req.URL.Query().Get("includeKvconfig") == "true" {
		getIndexRequest.IncludeKvconfig = true
	}

	// create proto message
	protoReq, err := getIndexRequest.MarshalProto()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create proto message")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// request
	resp, err := h.client.GetIndex(context.Background(), protoReq)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get index")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// create response
	getIndexResponse, err := util.NewGetIndexRespone(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create get index response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// output response
	output, err := json.MarshalIndent(getIndexResponse, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return
}
