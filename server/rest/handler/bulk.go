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
	"strconv"
)

type BulkHandler struct {
	client proto.IndigoClient
}

func NewBulkHandler(client proto.IndigoClient) *BulkHandler {
	return &BulkHandler{
		client: client,
	}
}

func (h *BulkHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	// create request
	bulkRequest, err := util.NewBulkRequest(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk request")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// overwrite request
	if req.URL.Query().Get("batchSize") != "" {
		i, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to set batch size")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bulkRequest.BatchSize = int32(i)
	}
	if bulkRequest.BatchSize <= 0 {
		bulkRequest.BatchSize = DefaultBatchSize
	}

	// create proto message
	protoReq, err := bulkRequest.MarshalProto()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create proto message")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// request
	resp, err := h.client.Bulk(context.Background(), protoReq)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to index documents in bulk")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// create response
	bulkResponse, err := util.NewBulkResponse(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// output response
	output, err := json.MarshalIndent(bulkResponse, "", "  ")
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
