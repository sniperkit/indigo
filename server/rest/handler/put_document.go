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
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

type PutDocumentHandler struct {
	client proto.IndigoClient
}

func NewPutDocumentHandler(client proto.IndigoClient) *PutDocumentHandler {
	return &PutDocumentHandler{
		client: client,
	}
}

func (h *PutDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)

	// create request
	putDocumentRequest, err := util.NewPutDocumentRequest(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create put document request")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// overwrite request
	putDocumentRequest.Document.Id = vars["id"]

	// create proto message
	protoReq, err := putDocumentRequest.MarshalProto()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create proto message")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// request
	resp, err := h.client.PutDocument(context.Background(), protoReq)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to put document")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// create response
	putDocumentResponse, err := util.NewPutDocumentResponse(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create put document response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// output response
	output, err := json.MarshalIndent(putDocumentResponse, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to create response")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return
}
