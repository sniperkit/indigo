package service

import (
	"encoding/json"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
)

type CreateIndexRequest struct {
	Index        string                    `json:"index,omitempty"`
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     map[string]interface{}    `json:"kvconfig,omitempty"`
}

func (cir *CreateIndexRequest) ProtoMessage() (*proto.CreateIndexRequest, error) {
	protoCreateIndexRequest := &proto.CreateIndexRequest{}

	// Index
	protoCreateIndexRequest.Index = cir.Index

	// IndexMapping
	indexMappingBytes, err := json.Marshal(cir.IndexMapping)
	if err != nil {
		return protoCreateIndexRequest, err
	}
	protoCreateIndexRequest.IndexMapping = indexMappingBytes

	// IndexType
	protoCreateIndexRequest.IndexType = cir.IndexType

	// Kvstore
	protoCreateIndexRequest.Kvstore = cir.Kvstore

	// Kvconfig
	kvconfigBytes, err := json.Marshal(cir.Kvconfig)
	if err != nil {
		return protoCreateIndexRequest, err
	}
	protoCreateIndexRequest.Kvconfig = kvconfigBytes

	return protoCreateIndexRequest, nil
}
