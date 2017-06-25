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

package util

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"io"
	"io/ioutil"
)

func NewIndexMapping(reader io.Reader) (*mapping.IndexMappingImpl, error) {
	indexMapping := mapping.NewIndexMapping()

	resourceBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resourceBytes, indexMapping)
	if err != nil {
		return nil, err
	}

	return indexMapping, nil
}

func NewKvconfig(reader io.Reader) (map[string]interface{}, error) {
	kvconfig := make(map[string]interface{})

	resourceBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resourceBytes, &kvconfig)
	if err != nil {
		return nil, err
	}

	return kvconfig, nil
}

type GetIndexRequest struct {
	IncludeIndexMapping bool `json:"include_index_mapping,omitempty"`
	IncludeIndexType    bool `json:"include_index_type,omitempty"`
	IncludeKvstore      bool `json:"include_kvstore,omitempty"`
	IncludeKvconfig     bool `json:"include_kvconfig,omitempty"`
}

func NewGetIndexRequest(includeIndexMapping bool, includeIndexType bool, includeKvstore bool, includeKvconfig bool) (*GetIndexRequest, error) {
	req := GetIndexRequest{
		IncludeIndexMapping: includeIndexMapping,
		IncludeIndexType:    includeIndexType,
		IncludeKvstore:      includeKvstore,
		IncludeKvconfig:     includeKvconfig,
	}

	return &req, nil
}

func (r *GetIndexRequest) MarshalProto() (*proto.GetIndexRequest, error) {
	req := &proto.GetIndexRequest{}

	req.IncludeIndexMapping = r.IncludeIndexMapping
	req.IncludeIndexType = r.IncludeIndexType
	req.IncludeKvstore = r.IncludeKvstore
	req.IncludeKvconfig = r.IncludeKvconfig

	return req, nil
}

type GetIndexResponse struct {
	Path         string                    `json:"path,omitempty"`
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     interface{}               `json:"kvconfig,omitempty"`
}

func NewGetIndexRespone(protoResp *proto.GetIndexResponse) (*GetIndexResponse, error) {
	resp := GetIndexResponse{
		Path:      protoResp.Path,
		IndexType: protoResp.IndexType,
		Kvstore:   protoResp.Kvstore,
	}

	if protoResp.IndexMapping != nil {
		indexMapping, err := UnmarshalAny(protoResp.IndexMapping)
		if err != nil {
			return nil, err
		}
		resp.IndexMapping = indexMapping.(*mapping.IndexMappingImpl)
	}

	if resp.Kvconfig != nil {
		kvconfig, err := UnmarshalAny(protoResp.Kvconfig)
		if err != nil {
			return nil, err
		}
		resp.Kvconfig = kvconfig
	}

	return &resp, nil
}

type Document struct {
	Id     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type UpdateRequest struct {
	Method   string   `json:"method,omitempty"`
	Document Document `json:"document,omitempty"`
}

type PutDocumentRequest struct {
	Document Document `json:"document,omitempty"`
}

func NewPutDocumentRequest(reader io.Reader) (*PutDocumentRequest, error) {
	req := PutDocumentRequest{}

	resourceBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resourceBytes, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *PutDocumentRequest) MarshalProto() (*proto.PutDocumentRequest, error) {
	req := &proto.PutDocumentRequest{}

	fieldsAny, err := MarshalAny(r.Document.Fields)
	if err != nil {
		return nil, err
	}

	document := &proto.Document{
		Id:     r.Document.Id,
		Fields: &fieldsAny,
	}

	req.Document = document

	return req, nil
}

type PutDocumentResponse struct {
	PutCount int32 `json:"put_count,omitempty"`
}

func NewPutDocumentResponse(protoResp *proto.PutDocumentResponse) (*PutDocumentResponse, error) {
	resp := &PutDocumentResponse{
		PutCount: protoResp.PutCount,
	}

	return resp, nil
}

type GetDocumentRequest struct {
	Id string `json:"id,omitempty"`
}

func NewGetDocumentRequest(id string) (*GetDocumentRequest, error) {
	req := GetDocumentRequest{
		Id: id,
	}

	return &req, nil
}

func (r *GetDocumentRequest) MarshalProto() (*proto.GetDocumentRequest, error) {
	req := &proto.GetDocumentRequest{}

	req.Id = r.Id

	return req, nil
}

type GetDocumentResponse struct {
	Document Document `json:"document,omitempty"`
}

func NewGetDocumentRespone(protoResp *proto.GetDocumentResponse) (*GetDocumentResponse, error) {
	fields, err := UnmarshalAny(protoResp.Document.Fields)
	if err != nil {
		return nil, err
	}

	document := Document{
		Id:     protoResp.Document.Id,
		Fields: *fields.(*map[string]interface{}),
	}

	resp := GetDocumentResponse{
		Document: document,
	}

	return &resp, nil
}

type DeleteDocumentRequest struct {
	Id string `json:"id,omitempty"`
}

func NewDeleteDocumentRequest(id string) (*DeleteDocumentRequest, error) {
	req := DeleteDocumentRequest{
		Id: id,
	}

	return &req, nil
}

func (r *DeleteDocumentRequest) MarshalProto() (*proto.DeleteDocumentRequest, error) {
	req := &proto.DeleteDocumentRequest{}

	req.Id = r.Id

	return req, nil
}

type DeleteDocumentResponse struct {
	DeleteCount int32 `json:"delete_count,omitempty"`
}

func NewDeleteDocumentResponse(protoResp *proto.DeleteDocumentResponse) (*DeleteDocumentResponse, error) {
	resp := &DeleteDocumentResponse{
		DeleteCount: protoResp.DeleteCount,
	}

	return resp, nil
}

type BulkRequest struct {
	BatchSize int32           `json:"batch_size,omitempty"`
	Requests  []UpdateRequest `json:"requests,omitempty"`
}

func NewBulkRequest(reader io.Reader) (*BulkRequest, error) {
	req := BulkRequest{}

	resourceBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resourceBytes, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *BulkRequest) MarshalProto() (*proto.BulkRequest, error) {
	var requests []*proto.UpdateRequest
	for _, request := range r.Requests {
		fieldsAny, err := MarshalAny(request.Document.Fields)
		if err != nil {
			return nil, err
		}
		document := proto.Document{
			Id:     request.Document.Id,
			Fields: &fieldsAny,
		}
		request := proto.UpdateRequest{
			Method:   request.Method,
			Document: &document,
		}
		requests = append(requests, &request)
	}

	req := &proto.BulkRequest{
		BatchSize: r.BatchSize,
		Requests:  requests,
	}

	return req, nil
}

type BulkResponse struct {
	PutCount      int32 `json:"put_count,omitempty"`
	PutErrorCount int32 `json:"put_error_count,omitempty"`
	DeleteCount   int32 `json:"delete_count,omitempty"`
}

func NewBulkResponse(protoResp *proto.BulkResponse) (*BulkResponse, error) {
	resp := &BulkResponse{
		PutCount:      protoResp.PutCount,
		PutErrorCount: protoResp.PutErrorCount,
		DeleteCount:   protoResp.DeleteCount,
	}

	return resp, nil
}

type SearchRequest struct {
	SearchRequest *bleve.SearchRequest `json:"search_request,omitempty"`
}

func NewSearchRequest(reader io.Reader) (*SearchRequest, error) {
	req := SearchRequest{}

	resourceBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resourceBytes, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *SearchRequest) MarshalProto() (*proto.SearchRequest, error) {
	req := &proto.SearchRequest{}

	searchRequestAny, err := MarshalAny(r.SearchRequest)
	if err != nil {
		return nil, err
	}

	req.SearchRequest = &searchRequestAny

	return req, nil
}

type SearchResponse struct {
	SearchResult *bleve.SearchResult `json:"search_result,omitempty"`
}

func NewSearchResonse(protoRes *proto.SearchResponse) (*SearchResponse, error) {
	resp := &SearchResponse{}

	searchResultAny, err := UnmarshalAny(protoRes.SearchResult)
	if err != nil {
		return nil, err
	}

	resp.SearchResult = searchResultAny.(*bleve.SearchResult)

	return resp, nil
}
