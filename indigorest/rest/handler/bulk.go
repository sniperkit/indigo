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
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/resource"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
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

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bulkResource := resource.BulkResource{}
	err = json.Unmarshal(resourceBytes, &bulkResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create bulk resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var b []*proto.BulkRequest_Request
	for _, request := range bulkResource.Requests {
		f, err := proto.MarshalAny(request.Document.Fields)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to create fields")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		d := proto.BulkRequest_Document{
			Id:     request.Document.Id,
			Fields: &f,
		}
		r := proto.BulkRequest_Request{
			Method:   request.Method,
			Document: &d,
		}
		b = append(b, &r)
	}

	protoBulkRequest := &proto.BulkRequest{
		BatchSize: bulkResource.BatchSize,
		Requests:  b,
	}

	if req.URL.Query().Get("batchSize") != "" {
		i, err := strconv.Atoi(req.URL.Query().Get("batchSize"))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to convert batch size")

			Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		protoBulkRequest.BatchSize = int32(i)
	}

	if protoBulkRequest.BatchSize == 0 {
		protoBulkRequest.BatchSize = DefaultBatchSize
	}

	resp, err := h.client.Bulk(context.Background(), protoBulkRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to index documents in bulk")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	output, err := json.MarshalIndent(resp, "", "  ")
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
