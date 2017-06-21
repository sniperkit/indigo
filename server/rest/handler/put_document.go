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
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
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
	id := vars["id"]

	resourceBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read request body")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	putDocumentResource := util.PutDocumentResource{}
	err = json.Unmarshal(resourceBytes, &putDocumentResource)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create put document resource")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fieldsAny, err := util.MarshalAny(putDocumentResource.Fields)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create fields")

		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protoPutDocumentRequest := &proto.PutDocumentRequest{
		Id:     id,
		Fields: &fieldsAny,
	}

	resp, err := h.client.PutDocument(context.Background(), protoPutDocumentRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).Error("failed to put document")

		Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	output, err := json.MarshalIndent(resp, "", "  ")
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
