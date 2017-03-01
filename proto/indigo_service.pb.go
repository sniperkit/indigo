// Code generated by protoc-gen-go.
// source: proto/indigo_service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/indigo_service.proto

It has these top-level messages:
	CreateIndexRequest
	CreateIndexResponse
	OpenIndexRequest
	OpenIndexResponse
	GetIndexRequest
	GetIndexResponse
	CloseIndexRequest
	CloseIndexResponse
	DeleteIndexRequest
	DeleteIndexResponse
	ListIndexRequest
	ListIndexResponse
	PutDocumentRequest
	PutDocumentResponse
	GetDocumentRequest
	GetDocumentResponse
	DeleteDocumentRequest
	DeleteDocumentResponse
	BulkRequest
	BulkResponse
	SearchRequest
	SearchResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateIndexRequest struct {
	IndexName    string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	IndexMapping []byte `protobuf:"bytes,2,opt,name=index_mapping,json=indexMapping,proto3" json:"index_mapping,omitempty"`
	IndexType    string `protobuf:"bytes,3,opt,name=index_type,json=indexType" json:"index_type,omitempty"`
	Kvstore      string `protobuf:"bytes,4,opt,name=kvstore" json:"kvstore,omitempty"`
	Kvconfig     []byte `protobuf:"bytes,5,opt,name=kvconfig,proto3" json:"kvconfig,omitempty"`
}

func (m *CreateIndexRequest) Reset()                    { *m = CreateIndexRequest{} }
func (m *CreateIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexRequest) ProtoMessage()               {}
func (*CreateIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *CreateIndexRequest) GetIndexMapping() []byte {
	if m != nil {
		return m.IndexMapping
	}
	return nil
}

func (m *CreateIndexRequest) GetIndexType() string {
	if m != nil {
		return m.IndexType
	}
	return ""
}

func (m *CreateIndexRequest) GetKvstore() string {
	if m != nil {
		return m.Kvstore
	}
	return ""
}

func (m *CreateIndexRequest) GetKvconfig() []byte {
	if m != nil {
		return m.Kvconfig
	}
	return nil
}

type CreateIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	IndexDir  string `protobuf:"bytes,2,opt,name=index_dir,json=indexDir" json:"index_dir,omitempty"`
}

func (m *CreateIndexResponse) Reset()                    { *m = CreateIndexResponse{} }
func (m *CreateIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexResponse) ProtoMessage()               {}
func (*CreateIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *CreateIndexResponse) GetIndexDir() string {
	if m != nil {
		return m.IndexDir
	}
	return ""
}

type OpenIndexRequest struct {
	IndexName     string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	RuntimeConfig []byte `protobuf:"bytes,2,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty"`
}

func (m *OpenIndexRequest) Reset()                    { *m = OpenIndexRequest{} }
func (m *OpenIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexRequest) ProtoMessage()               {}
func (*OpenIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OpenIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *OpenIndexRequest) GetRuntimeConfig() []byte {
	if m != nil {
		return m.RuntimeConfig
	}
	return nil
}

type OpenIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	IndexDir  string `protobuf:"bytes,2,opt,name=index_dir,json=indexDir" json:"index_dir,omitempty"`
}

func (m *OpenIndexResponse) Reset()                    { *m = OpenIndexResponse{} }
func (m *OpenIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexResponse) ProtoMessage()               {}
func (*OpenIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OpenIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *OpenIndexResponse) GetIndexDir() string {
	if m != nil {
		return m.IndexDir
	}
	return ""
}

type GetIndexRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *GetIndexRequest) Reset()                    { *m = GetIndexRequest{} }
func (m *GetIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetIndexRequest) ProtoMessage()               {}
func (*GetIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type GetIndexResponse struct {
	DocumentCount uint64 `protobuf:"varint,1,opt,name=document_count,json=documentCount" json:"document_count,omitempty"`
	IndexStats    []byte `protobuf:"bytes,2,opt,name=index_stats,json=indexStats,proto3" json:"index_stats,omitempty"`
	IndexMapping  []byte `protobuf:"bytes,3,opt,name=index_mapping,json=indexMapping,proto3" json:"index_mapping,omitempty"`
}

func (m *GetIndexResponse) Reset()                    { *m = GetIndexResponse{} }
func (m *GetIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetIndexResponse) ProtoMessage()               {}
func (*GetIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetIndexResponse) GetDocumentCount() uint64 {
	if m != nil {
		return m.DocumentCount
	}
	return 0
}

func (m *GetIndexResponse) GetIndexStats() []byte {
	if m != nil {
		return m.IndexStats
	}
	return nil
}

func (m *GetIndexResponse) GetIndexMapping() []byte {
	if m != nil {
		return m.IndexMapping
	}
	return nil
}

type CloseIndexRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *CloseIndexRequest) Reset()                    { *m = CloseIndexRequest{} }
func (m *CloseIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexRequest) ProtoMessage()               {}
func (*CloseIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CloseIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type CloseIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *CloseIndexResponse) Reset()                    { *m = CloseIndexResponse{} }
func (m *CloseIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexResponse) ProtoMessage()               {}
func (*CloseIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CloseIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type DeleteIndexRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *DeleteIndexRequest) Reset()                    { *m = DeleteIndexRequest{} }
func (m *DeleteIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexRequest) ProtoMessage()               {}
func (*DeleteIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type DeleteIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *DeleteIndexResponse) Reset()                    { *m = DeleteIndexResponse{} }
func (m *DeleteIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexResponse) ProtoMessage()               {}
func (*DeleteIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type ListIndexRequest struct {
}

func (m *ListIndexRequest) Reset()                    { *m = ListIndexRequest{} }
func (m *ListIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexRequest) ProtoMessage()               {}
func (*ListIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ListIndexResponse struct {
	Indices []string `protobuf:"bytes,1,rep,name=indices" json:"indices,omitempty"`
}

func (m *ListIndexResponse) Reset()                    { *m = ListIndexResponse{} }
func (m *ListIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexResponse) ProtoMessage()               {}
func (*ListIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListIndexResponse) GetIndices() []string {
	if m != nil {
		return m.Indices
	}
	return nil
}

type PutDocumentRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Fields    []byte `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (m *PutDocumentRequest) Reset()                    { *m = PutDocumentRequest{} }
func (m *PutDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*PutDocumentRequest) ProtoMessage()               {}
func (*PutDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PutDocumentRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *PutDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutDocumentRequest) GetFields() []byte {
	if m != nil {
		return m.Fields
	}
	return nil
}

type PutDocumentResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *PutDocumentResponse) Reset()                    { *m = PutDocumentResponse{} }
func (m *PutDocumentResponse) String() string            { return proto1.CompactTextString(m) }
func (*PutDocumentResponse) ProtoMessage()               {}
func (*PutDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PutDocumentResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetDocumentRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetDocumentRequest) Reset()                    { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()               {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetDocumentRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *GetDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDocumentResponse struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Fields []byte `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (m *GetDocumentResponse) Reset()                    { *m = GetDocumentResponse{} }
func (m *GetDocumentResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetDocumentResponse) ProtoMessage()               {}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetDocumentResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDocumentResponse) GetFields() []byte {
	if m != nil {
		return m.Fields
	}
	return nil
}

type DeleteDocumentRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteDocumentRequest) Reset()                    { *m = DeleteDocumentRequest{} }
func (m *DeleteDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteDocumentRequest) ProtoMessage()               {}
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *DeleteDocumentRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *DeleteDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteDocumentResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteDocumentResponse) Reset()                    { *m = DeleteDocumentResponse{} }
func (m *DeleteDocumentResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteDocumentResponse) ProtoMessage()               {}
func (*DeleteDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *DeleteDocumentResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type BulkRequest struct {
	IndexName   string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	BulkRequest []byte `protobuf:"bytes,2,opt,name=bulk_request,json=bulkRequest,proto3" json:"bulk_request,omitempty"`
	BatchSize   int32  `protobuf:"varint,3,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
}

func (m *BulkRequest) Reset()                    { *m = BulkRequest{} }
func (m *BulkRequest) String() string            { return proto1.CompactTextString(m) }
func (*BulkRequest) ProtoMessage()               {}
func (*BulkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *BulkRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *BulkRequest) GetBulkRequest() []byte {
	if m != nil {
		return m.BulkRequest
	}
	return nil
}

func (m *BulkRequest) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

type BulkResponse struct {
	PutCount      int32 `protobuf:"varint,1,opt,name=put_count,json=putCount" json:"put_count,omitempty"`
	PutErrorCount int32 `protobuf:"varint,2,opt,name=put_error_count,json=putErrorCount" json:"put_error_count,omitempty"`
	DeleteCount   int32 `protobuf:"varint,3,opt,name=delete_count,json=deleteCount" json:"delete_count,omitempty"`
}

func (m *BulkResponse) Reset()                    { *m = BulkResponse{} }
func (m *BulkResponse) String() string            { return proto1.CompactTextString(m) }
func (*BulkResponse) ProtoMessage()               {}
func (*BulkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *BulkResponse) GetPutCount() int32 {
	if m != nil {
		return m.PutCount
	}
	return 0
}

func (m *BulkResponse) GetPutErrorCount() int32 {
	if m != nil {
		return m.PutErrorCount
	}
	return 0
}

func (m *BulkResponse) GetDeleteCount() int32 {
	if m != nil {
		return m.DeleteCount
	}
	return 0
}

type SearchRequest struct {
	IndexName     string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	SearchRequest []byte `protobuf:"bytes,2,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto1.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *SearchRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *SearchRequest) GetSearchRequest() []byte {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

type SearchResponse struct {
	SearchResult []byte `protobuf:"bytes,1,opt,name=search_result,json=searchResult,proto3" json:"search_result,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto1.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *SearchResponse) GetSearchResult() []byte {
	if m != nil {
		return m.SearchResult
	}
	return nil
}

func init() {
	proto1.RegisterType((*CreateIndexRequest)(nil), "proto.CreateIndexRequest")
	proto1.RegisterType((*CreateIndexResponse)(nil), "proto.CreateIndexResponse")
	proto1.RegisterType((*OpenIndexRequest)(nil), "proto.OpenIndexRequest")
	proto1.RegisterType((*OpenIndexResponse)(nil), "proto.OpenIndexResponse")
	proto1.RegisterType((*GetIndexRequest)(nil), "proto.GetIndexRequest")
	proto1.RegisterType((*GetIndexResponse)(nil), "proto.GetIndexResponse")
	proto1.RegisterType((*CloseIndexRequest)(nil), "proto.CloseIndexRequest")
	proto1.RegisterType((*CloseIndexResponse)(nil), "proto.CloseIndexResponse")
	proto1.RegisterType((*DeleteIndexRequest)(nil), "proto.DeleteIndexRequest")
	proto1.RegisterType((*DeleteIndexResponse)(nil), "proto.DeleteIndexResponse")
	proto1.RegisterType((*ListIndexRequest)(nil), "proto.ListIndexRequest")
	proto1.RegisterType((*ListIndexResponse)(nil), "proto.ListIndexResponse")
	proto1.RegisterType((*PutDocumentRequest)(nil), "proto.PutDocumentRequest")
	proto1.RegisterType((*PutDocumentResponse)(nil), "proto.PutDocumentResponse")
	proto1.RegisterType((*GetDocumentRequest)(nil), "proto.GetDocumentRequest")
	proto1.RegisterType((*GetDocumentResponse)(nil), "proto.GetDocumentResponse")
	proto1.RegisterType((*DeleteDocumentRequest)(nil), "proto.DeleteDocumentRequest")
	proto1.RegisterType((*DeleteDocumentResponse)(nil), "proto.DeleteDocumentResponse")
	proto1.RegisterType((*BulkRequest)(nil), "proto.BulkRequest")
	proto1.RegisterType((*BulkResponse)(nil), "proto.BulkResponse")
	proto1.RegisterType((*SearchRequest)(nil), "proto.SearchRequest")
	proto1.RegisterType((*SearchResponse)(nil), "proto.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Indigo service

type IndigoClient interface {
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error)
	OpenIndex(ctx context.Context, in *OpenIndexRequest, opts ...grpc.CallOption) (*OpenIndexResponse, error)
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error)
	CloseIndex(ctx context.Context, in *CloseIndexRequest, opts ...grpc.CallOption) (*CloseIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	ListIndex(ctx context.Context, in *ListIndexRequest, opts ...grpc.CallOption) (*ListIndexResponse, error)
	PutDocument(ctx context.Context, in *PutDocumentRequest, opts ...grpc.CallOption) (*PutDocumentResponse, error)
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error)
	Bulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type indigoClient struct {
	cc *grpc.ClientConn
}

func NewIndigoClient(cc *grpc.ClientConn) IndigoClient {
	return &indigoClient{cc}
}

func (c *indigoClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error) {
	out := new(CreateIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/CreateIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) OpenIndex(ctx context.Context, in *OpenIndexRequest, opts ...grpc.CallOption) (*OpenIndexResponse, error) {
	out := new(OpenIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/OpenIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error) {
	out := new(GetIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/GetIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) CloseIndex(ctx context.Context, in *CloseIndexRequest, opts ...grpc.CallOption) (*CloseIndexResponse, error) {
	out := new(CloseIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/CloseIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/DeleteIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) ListIndex(ctx context.Context, in *ListIndexRequest, opts ...grpc.CallOption) (*ListIndexResponse, error) {
	out := new(ListIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/ListIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) PutDocument(ctx context.Context, in *PutDocumentRequest, opts ...grpc.CallOption) (*PutDocumentResponse, error) {
	out := new(PutDocumentResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/PutDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/GetDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error) {
	out := new(DeleteDocumentResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/DeleteDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) Bulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error) {
	out := new(BulkResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/Bulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indigoClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Indigo service

type IndigoServer interface {
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error)
	OpenIndex(context.Context, *OpenIndexRequest) (*OpenIndexResponse, error)
	GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error)
	CloseIndex(context.Context, *CloseIndexRequest) (*CloseIndexResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	ListIndex(context.Context, *ListIndexRequest) (*ListIndexResponse, error)
	PutDocument(context.Context, *PutDocumentRequest) (*PutDocumentResponse, error)
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	Bulk(context.Context, *BulkRequest) (*BulkResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterIndigoServer(s *grpc.Server, srv IndigoServer) {
	s.RegisterService(&_Indigo_serviceDesc, srv)
}

func _Indigo_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_OpenIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).OpenIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/OpenIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).OpenIndex(ctx, req.(*OpenIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/GetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).GetIndex(ctx, req.(*GetIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_CloseIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).CloseIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/CloseIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).CloseIndex(ctx, req.(*CloseIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_ListIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).ListIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/ListIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).ListIndex(ctx, req.(*ListIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_PutDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).PutDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/PutDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).PutDocument(ctx, req.(*PutDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_Bulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).Bulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/Bulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).Bulk(ctx, req.(*BulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indigo_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndigoServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Indigo/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndigoServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Indigo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Indigo",
	HandlerType: (*IndigoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIndex",
			Handler:    _Indigo_CreateIndex_Handler,
		},
		{
			MethodName: "OpenIndex",
			Handler:    _Indigo_OpenIndex_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _Indigo_GetIndex_Handler,
		},
		{
			MethodName: "CloseIndex",
			Handler:    _Indigo_CloseIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _Indigo_DeleteIndex_Handler,
		},
		{
			MethodName: "ListIndex",
			Handler:    _Indigo_ListIndex_Handler,
		},
		{
			MethodName: "PutDocument",
			Handler:    _Indigo_PutDocument_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _Indigo_GetDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _Indigo_DeleteDocument_Handler,
		},
		{
			MethodName: "Bulk",
			Handler:    _Indigo_Bulk_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Indigo_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indigo_service.proto",
}

func init() { proto1.RegisterFile("proto/indigo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0x25, 0x81, 0x84, 0x64, 0xf2, 0x01, 0x6c, 0x2e, 0x60, 0xcc, 0x45, 0x17, 0x16, 0x71, 0xc5,
	0x4b, 0xa1, 0x85, 0x56, 0x7d, 0x42, 0xaa, 0x1a, 0x0a, 0x42, 0x6a, 0x4b, 0x6b, 0x5a, 0xa9, 0x52,
	0x1f, 0xa2, 0xc4, 0x5e, 0x60, 0x95, 0xc4, 0x76, 0xbd, 0xeb, 0xa8, 0xf0, 0xd2, 0x5f, 0xd4, 0x9f,
	0xd6, 0xff, 0x50, 0x79, 0x77, 0xbd, 0xf1, 0x47, 0x2a, 0x12, 0xd1, 0xa7, 0x68, 0xcf, 0xce, 0x9c,
	0x39, 0x33, 0x99, 0x3d, 0x06, 0xd3, 0x0f, 0x3c, 0xee, 0x1d, 0x52, 0xd7, 0xa1, 0x37, 0x5e, 0x87,
	0x91, 0x60, 0x44, 0x6d, 0x72, 0x20, 0x40, 0x54, 0x12, 0x3f, 0xf8, 0x67, 0x01, 0x50, 0x3b, 0x20,
	0x5d, 0x4e, 0x2e, 0x5c, 0x87, 0x7c, 0xb7, 0xc8, 0xb7, 0x90, 0x30, 0x8e, 0xb6, 0x00, 0x68, 0x74,
	0xee, 0xb8, 0xdd, 0x21, 0x31, 0x0a, 0xdb, 0x85, 0xfd, 0xaa, 0x55, 0x15, 0xc8, 0xfb, 0xee, 0x90,
	0xa0, 0x5d, 0x68, 0xc8, 0xeb, 0x61, 0xd7, 0xf7, 0xa9, 0x7b, 0x63, 0x14, 0xb7, 0x0b, 0xfb, 0x75,
	0xab, 0x2e, 0xc0, 0x77, 0x12, 0x1b, 0x73, 0xf0, 0x3b, 0x9f, 0x18, 0xf3, 0x09, 0x8e, 0x4f, 0x77,
	0x3e, 0x41, 0x06, 0x2c, 0xf6, 0x47, 0x8c, 0x7b, 0x01, 0x31, 0x16, 0xc4, 0x5d, 0x7c, 0x44, 0x26,
	0x54, 0xfa, 0x23, 0xdb, 0x73, 0xaf, 0xe9, 0x8d, 0x51, 0x12, 0xc4, 0xfa, 0x8c, 0x3f, 0x42, 0x2b,
	0x25, 0x97, 0xf9, 0x9e, 0xcb, 0xc8, 0x43, 0x7a, 0x37, 0x41, 0x1e, 0x3a, 0x0e, 0x0d, 0x84, 0xd6,
	0xaa, 0x55, 0x11, 0xc0, 0x29, 0x0d, 0xf0, 0x17, 0x58, 0xbe, 0xf4, 0x89, 0x3b, 0x4b, 0xff, 0x7b,
	0xd0, 0x0c, 0x42, 0x97, 0xd3, 0x21, 0xe9, 0x28, 0x9d, 0x72, 0x00, 0x0d, 0x85, 0xb6, 0xa5, 0xd8,
	0x4b, 0x58, 0x49, 0x30, 0xff, 0x05, 0xa9, 0x4f, 0x61, 0xe9, 0x9c, 0xf0, 0x19, 0x94, 0xe2, 0x1f,
	0xb0, 0x3c, 0xce, 0x50, 0x0a, 0xf6, 0xa0, 0xe9, 0x78, 0x76, 0x38, 0x24, 0x2e, 0xef, 0xd8, 0x5e,
	0xe8, 0x72, 0x91, 0xb6, 0x60, 0x35, 0x62, 0xb4, 0x1d, 0x81, 0xe8, 0x3f, 0xa8, 0x49, 0x66, 0xc6,
	0xbb, 0x9c, 0xa9, 0x0e, 0x65, 0xb1, 0xab, 0x08, 0xc9, 0x6f, 0xc1, 0x7c, 0x7e, 0x0b, 0xf0, 0x11,
	0xac, 0xb4, 0x07, 0x1e, 0x9b, 0x65, 0xbd, 0xf0, 0x31, 0xa0, 0x64, 0xce, 0x54, 0x83, 0x8b, 0x92,
	0x4e, 0xc9, 0x80, 0xcc, 0xb4, 0xc8, 0xf8, 0x39, 0xb4, 0x52, 0x49, 0xd3, 0x95, 0x42, 0xb0, 0xfc,
	0x96, 0xb2, 0xd4, 0xff, 0x80, 0x9f, 0xc0, 0x4a, 0x02, 0x53, 0x3c, 0x06, 0x2c, 0x46, 0x8f, 0xcf,
	0x26, 0xcc, 0x28, 0x6c, 0xcf, 0x47, 0x3b, 0xae, 0x8e, 0xf8, 0x2b, 0xa0, 0x0f, 0x21, 0x3f, 0x55,
	0x03, 0x9f, 0x72, 0xed, 0x9a, 0x50, 0xa4, 0x8e, 0x5a, 0x8a, 0x22, 0x75, 0xd0, 0x1a, 0x94, 0xaf,
	0x29, 0x19, 0x38, 0x4c, 0x4d, 0x5e, 0x9d, 0xf0, 0x21, 0xb4, 0x52, 0xe4, 0x63, 0x35, 0x2c, 0xb4,
	0x6d, 0xc2, 0x98, 0xa0, 0xae, 0x58, 0xf1, 0x11, 0xb7, 0x01, 0x9d, 0x93, 0x47, 0xaa, 0xc1, 0x27,
	0xd0, 0x4a, 0x91, 0xa8, 0xaa, 0x32, 0xac, 0x30, 0x41, 0x74, 0x31, 0x25, 0xfa, 0x0c, 0x56, 0xe5,
	0x5f, 0xf1, 0x48, 0x19, 0x47, 0xb0, 0x96, 0xe5, 0x79, 0xb0, 0x7f, 0x17, 0x6a, 0xaf, 0xc3, 0x41,
	0x7f, 0xca, 0x8a, 0x3b, 0x50, 0xef, 0x85, 0x83, 0x7e, 0x27, 0x90, 0xe1, 0xaa, 0x8f, 0x5a, 0x2f,
	0xcd, 0xd0, 0xeb, 0x72, 0xfb, 0xb6, 0xc3, 0xe8, 0xbd, 0xf4, 0xbe, 0x92, 0x55, 0x15, 0xc8, 0x15,
	0xbd, 0x27, 0x78, 0x04, 0x75, 0x59, 0x4f, 0x29, 0xdb, 0x84, 0xaa, 0x1f, 0x26, 0x1f, 0x63, 0xc9,
	0xaa, 0xf8, 0xa1, 0x7a, 0x87, 0xff, 0xc3, 0x52, 0x74, 0x49, 0x82, 0xc0, 0x0b, 0x54, 0x48, 0x51,
	0x84, 0x34, 0xfc, 0x90, 0xbf, 0x89, 0x50, 0x19, 0xb7, 0x03, 0x75, 0x47, 0x34, 0xae, 0x82, 0x64,
	0xd5, 0x9a, 0xc4, 0x44, 0x08, 0xfe, 0x0c, 0x8d, 0x2b, 0xd2, 0x0d, 0xec, 0xdb, 0xe9, 0x7d, 0x8e,
	0x89, 0xf8, 0x4c, 0xaf, 0x0d, 0x96, 0x64, 0xc1, 0x2f, 0xa0, 0x19, 0xd3, 0xaa, 0x86, 0x76, 0xa1,
	0xa1, 0x13, 0x59, 0x38, 0x90, 0x4d, 0xd5, 0xad, 0x7a, 0x9c, 0x17, 0x61, 0x47, 0xbf, 0x4a, 0x50,
	0xbe, 0x10, 0xdf, 0x26, 0x74, 0x06, 0xb5, 0x84, 0xad, 0xa3, 0x0d, 0xf9, 0x91, 0x3a, 0xc8, 0x7f,
	0x99, 0x4c, 0x73, 0xd2, 0x95, 0xac, 0x8a, 0xe7, 0xd0, 0x2b, 0xa8, 0x6a, 0xc7, 0x45, 0xeb, 0x2a,
	0x34, 0xeb, 0xee, 0xa6, 0x91, 0xbf, 0xd0, 0x0c, 0x27, 0x50, 0x89, 0x0d, 0x13, 0xad, 0xa9, 0xb8,
	0x8c, 0xe7, 0x9a, 0xeb, 0x39, 0x5c, 0xa7, 0xb7, 0x01, 0xc6, 0xd6, 0x85, 0xe2, 0x42, 0x39, 0x07,
	0x34, 0x37, 0x26, 0xdc, 0x68, 0x92, 0x33, 0xa8, 0x25, 0x5c, 0x49, 0x4f, 0x23, 0x6f, 0x6f, 0x7a,
	0x1a, 0x13, 0x4c, 0x4c, 0x4e, 0x43, 0x7b, 0x92, 0x9e, 0x46, 0xd6, 0xb9, 0xf4, 0x34, 0x72, 0xf6,
	0x25, 0x95, 0x24, 0x9c, 0x44, 0x2b, 0xc9, 0x5b, 0x97, 0x56, 0x32, 0xc1, 0x78, 0x24, 0x4f, 0xc2,
	0x1b, 0x34, 0x4f, 0xde, 0x74, 0x34, 0xcf, 0x04, 0x2b, 0xc1, 0x73, 0xe8, 0x12, 0x9a, 0xe9, 0xc7,
	0x8d, 0xfe, 0x4d, 0x4d, 0x20, 0xcb, 0xb6, 0xf5, 0x87, 0x5b, 0x4d, 0xf8, 0x0c, 0x16, 0xa2, 0x97,
	0x88, 0x90, 0x0a, 0x4c, 0xd8, 0x80, 0xd9, 0x4a, 0x61, 0x3a, 0xe5, 0x25, 0x94, 0xe5, 0xb6, 0xa3,
	0x7f, 0x54, 0x40, 0xea, 0x4d, 0x99, 0xab, 0x19, 0x34, 0x4e, 0xec, 0x95, 0x05, 0x7e, 0xfc, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x5f, 0x0b, 0xe7, 0x97, 0x09, 0x00, 0x00,
}
