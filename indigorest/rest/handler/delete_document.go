//  Copyright (c) 2015 Minoru Osuka
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
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

type DeleteDocumentHandler struct {
	client proto.IndigoClient
}

func NewDeleteDocumentHandler(client proto.IndigoClient) *DeleteDocumentHandler {
	return &DeleteDocumentHandler{
		client: client,
	}
}

func (h *DeleteDocumentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{
		"req": req,
	}).Info("")

	vars := mux.Vars(req)
	id := vars["id"]

	protoDeleteDocumentRequest := &proto.DeleteDocumentRequest{
		Id: id,
	}

	resp, err := h.client.DeleteDocument(context.Background(), protoDeleteDocumentRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to delete document")

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
