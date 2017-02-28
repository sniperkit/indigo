// Code generated by protoc-gen-go.
// source: proto/indigo_service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/indigo_service.proto

It has these top-level messages:
	ListIndexRequest
	ListIndexResponse
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

type ListIndexRequest struct {
}

func (m *ListIndexRequest) Reset()                    { *m = ListIndexRequest{} }
func (m *ListIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexRequest) ProtoMessage()               {}
func (*ListIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListIndexResponse struct {
	IndexNames []string `protobuf:"bytes,1,rep,name=indexNames" json:"indexNames,omitempty"`
}

func (m *ListIndexResponse) Reset()                    { *m = ListIndexResponse{} }
func (m *ListIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexResponse) ProtoMessage()               {}
func (*ListIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListIndexResponse) GetIndexNames() []string {
	if m != nil {
		return m.IndexNames
	}
	return nil
}

type CreateIndexRequest struct {
	IndexName    string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	IndexMapping []byte `protobuf:"bytes,2,opt,name=indexMapping,proto3" json:"indexMapping,omitempty"`
	IndexType    string `protobuf:"bytes,3,opt,name=indexType" json:"indexType,omitempty"`
	KvStore      string `protobuf:"bytes,4,opt,name=kvStore" json:"kvStore,omitempty"`
	KvConfig     []byte `protobuf:"bytes,5,opt,name=kvConfig,proto3" json:"kvConfig,omitempty"`
}

func (m *CreateIndexRequest) Reset()                    { *m = CreateIndexRequest{} }
func (m *CreateIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexRequest) ProtoMessage()               {}
func (*CreateIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *CreateIndexRequest) GetKvStore() string {
	if m != nil {
		return m.KvStore
	}
	return ""
}

func (m *CreateIndexRequest) GetKvConfig() []byte {
	if m != nil {
		return m.KvConfig
	}
	return nil
}

type CreateIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *CreateIndexResponse) Reset()                    { *m = CreateIndexResponse{} }
func (m *CreateIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexResponse) ProtoMessage()               {}
func (*CreateIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type OpenIndexRequest struct {
	IndexName     string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	RuntimeConfig []byte `protobuf:"bytes,2,opt,name=runtimeConfig,proto3" json:"runtimeConfig,omitempty"`
}

func (m *OpenIndexRequest) Reset()                    { *m = OpenIndexRequest{} }
func (m *OpenIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexRequest) ProtoMessage()               {}
func (*OpenIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *OpenIndexResponse) Reset()                    { *m = OpenIndexResponse{} }
func (m *OpenIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexResponse) ProtoMessage()               {}
func (*OpenIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OpenIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type GetIndexRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *GetIndexRequest) Reset()                    { *m = GetIndexRequest{} }
func (m *GetIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetIndexRequest) ProtoMessage()               {}
func (*GetIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type GetIndexResponse struct {
	DocumentCount uint64 `protobuf:"varint,1,opt,name=documentCount" json:"documentCount,omitempty"`
	IndexStats    []byte `protobuf:"bytes,2,opt,name=indexStats,proto3" json:"indexStats,omitempty"`
	IndexMapping  []byte `protobuf:"bytes,3,opt,name=indexMapping,proto3" json:"indexMapping,omitempty"`
}

func (m *GetIndexResponse) Reset()                    { *m = GetIndexResponse{} }
func (m *GetIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetIndexResponse) ProtoMessage()               {}
func (*GetIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *CloseIndexRequest) Reset()                    { *m = CloseIndexRequest{} }
func (m *CloseIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexRequest) ProtoMessage()               {}
func (*CloseIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CloseIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type CloseIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *CloseIndexResponse) Reset()                    { *m = CloseIndexResponse{} }
func (m *CloseIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexResponse) ProtoMessage()               {}
func (*CloseIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CloseIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type DeleteIndexRequest struct {
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *DeleteIndexRequest) Reset()                    { *m = DeleteIndexRequest{} }
func (m *DeleteIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexRequest) ProtoMessage()               {}
func (*DeleteIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteIndexRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type DeleteIndexResponse struct {
	IndexName string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
}

func (m *DeleteIndexResponse) Reset()                    { *m = DeleteIndexResponse{} }
func (m *DeleteIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexResponse) ProtoMessage()               {}
func (*DeleteIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteIndexResponse) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type PutDocumentRequest struct {
	IndexName  string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	DocumentID string `protobuf:"bytes,2,opt,name=documentID" json:"documentID,omitempty"`
	Document   []byte `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
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

func (m *PutDocumentRequest) GetDocumentID() string {
	if m != nil {
		return m.DocumentID
	}
	return ""
}

func (m *PutDocumentRequest) GetDocument() []byte {
	if m != nil {
		return m.Document
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
	IndexName  string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	DocumentID string `protobuf:"bytes,2,opt,name=documentID" json:"documentID,omitempty"`
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

func (m *GetDocumentRequest) GetDocumentID() string {
	if m != nil {
		return m.DocumentID
	}
	return ""
}

type GetDocumentResponse struct {
	Document []byte `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (m *GetDocumentResponse) Reset()                    { *m = GetDocumentResponse{} }
func (m *GetDocumentResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetDocumentResponse) ProtoMessage()               {}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetDocumentResponse) GetDocument() []byte {
	if m != nil {
		return m.Document
	}
	return nil
}

type DeleteDocumentRequest struct {
	IndexName  string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	DocumentID string `protobuf:"bytes,2,opt,name=documentID" json:"documentID,omitempty"`
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

func (m *DeleteDocumentRequest) GetDocumentID() string {
	if m != nil {
		return m.DocumentID
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
	IndexName   string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	BulkRequest []byte `protobuf:"bytes,2,opt,name=bulkRequest,proto3" json:"bulkRequest,omitempty"`
	BatchSize   int32  `protobuf:"varint,3,opt,name=batchSize" json:"batchSize,omitempty"`
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
	PutCount      int32 `protobuf:"varint,1,opt,name=putCount" json:"putCount,omitempty"`
	PutErrorCount int32 `protobuf:"varint,2,opt,name=putErrorCount" json:"putErrorCount,omitempty"`
	DeleteCount   int32 `protobuf:"varint,3,opt,name=deleteCount" json:"deleteCount,omitempty"`
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
	IndexName     string `protobuf:"bytes,1,opt,name=indexName" json:"indexName,omitempty"`
	SearchRequest []byte `protobuf:"bytes,2,opt,name=searchRequest,proto3" json:"searchRequest,omitempty"`
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
	SearchResult []byte `protobuf:"bytes,1,opt,name=searchResult,proto3" json:"searchResult,omitempty"`
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
	proto1.RegisterType((*ListIndexRequest)(nil), "proto.ListIndexRequest")
	proto1.RegisterType((*ListIndexResponse)(nil), "proto.ListIndexResponse")
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
	ListIndex(ctx context.Context, in *ListIndexRequest, opts ...grpc.CallOption) (*ListIndexResponse, error)
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error)
	OpenIndex(ctx context.Context, in *OpenIndexRequest, opts ...grpc.CallOption) (*OpenIndexResponse, error)
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error)
	CloseIndex(ctx context.Context, in *CloseIndexRequest, opts ...grpc.CallOption) (*CloseIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
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

func (c *indigoClient) ListIndex(ctx context.Context, in *ListIndexRequest, opts ...grpc.CallOption) (*ListIndexResponse, error) {
	out := new(ListIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/ListIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	ListIndex(context.Context, *ListIndexRequest) (*ListIndexResponse, error)
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error)
	OpenIndex(context.Context, *OpenIndexRequest) (*OpenIndexResponse, error)
	GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error)
	CloseIndex(context.Context, *CloseIndexRequest) (*CloseIndexResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	PutDocument(context.Context, *PutDocumentRequest) (*PutDocumentResponse, error)
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	Bulk(context.Context, *BulkRequest) (*BulkResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterIndigoServer(s *grpc.Server, srv IndigoServer) {
	s.RegisterService(&_Indigo_serviceDesc, srv)
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
			MethodName: "ListIndex",
			Handler:    _Indigo_ListIndex_Handler,
		},
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
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x51, 0x4f, 0xdb, 0x3c,
	0x14, 0xa5, 0x40, 0x81, 0x5e, 0x0a, 0x1f, 0xdc, 0x7e, 0x40, 0x88, 0x18, 0xaa, 0x2c, 0x1e, 0x78,
	0x02, 0x51, 0x26, 0xed, 0x69, 0x12, 0x5a, 0xd9, 0x10, 0xd2, 0x36, 0xa6, 0x74, 0xdb, 0xeb, 0x14,
	0x52, 0x0f, 0xa2, 0x96, 0x24, 0x8b, 0x1d, 0xb4, 0x4d, 0xfb, 0x45, 0xfb, 0x6f, 0xfb, 0x0f, 0x53,
	0x6c, 0xc7, 0xb1, 0x93, 0x6e, 0x0b, 0x12, 0x4f, 0x91, 0x8f, 0xef, 0x3d, 0xf7, 0xf8, 0xe6, 0xfa,
	0x18, 0xdc, 0x24, 0x8d, 0x79, 0x7c, 0x1c, 0x46, 0xe3, 0xf0, 0x26, 0xfe, 0xc4, 0x68, 0x7a, 0x1f,
	0x06, 0xf4, 0x48, 0x80, 0xd8, 0x16, 0x1f, 0x82, 0xb0, 0xf1, 0x3a, 0x64, 0xfc, 0x32, 0x1a, 0xd3,
	0xaf, 0x1e, 0xfd, 0x92, 0x51, 0xc6, 0xc9, 0x29, 0x6c, 0x1a, 0x18, 0x4b, 0xe2, 0x88, 0x51, 0xdc,
	0x07, 0x08, 0x73, 0xe0, 0xad, 0x7f, 0x47, 0x99, 0xd3, 0xea, 0x2f, 0x1c, 0x76, 0x3c, 0x03, 0x21,
	0x3f, 0x5b, 0x80, 0xc3, 0x94, 0xfa, 0x9c, 0x9a, 0x5c, 0xb8, 0x07, 0x1d, 0x1d, 0xe4, 0xb4, 0xfa,
	0xad, 0xc3, 0x8e, 0x57, 0x02, 0x48, 0xa0, 0x2b, 0x16, 0x6f, 0xfc, 0x24, 0x09, 0xa3, 0x1b, 0x67,
	0xbe, 0xdf, 0x3a, 0xec, 0x7a, 0x16, 0xa6, 0x19, 0xde, 0x7f, 0x4b, 0xa8, 0xb3, 0x60, 0x30, 0xe4,
	0x00, 0x3a, 0xb0, 0x3c, 0xb9, 0x1f, 0xf1, 0x38, 0xa5, 0xce, 0xa2, 0xd8, 0x2b, 0x96, 0xe8, 0xc2,
	0xca, 0xe4, 0x7e, 0x18, 0x47, 0x9f, 0xc3, 0x1b, 0xa7, 0x2d, 0x78, 0xf5, 0x9a, 0x9c, 0x42, 0xcf,
	0xd2, 0xaa, 0xce, 0xf8, 0x57, 0xb1, 0xe4, 0x23, 0x6c, 0x5c, 0x25, 0x34, 0x7a, 0xc0, 0xf1, 0x0e,
	0x60, 0x2d, 0xcd, 0x22, 0x1e, 0xde, 0x51, 0xa5, 0x43, 0x9e, 0xcf, 0x06, 0xc9, 0x09, 0x6c, 0x1a,
	0xbc, 0x8d, 0xa4, 0x1c, 0xc3, 0x7f, 0x17, 0x94, 0x37, 0x57, 0x42, 0x7e, 0xc0, 0x46, 0x99, 0xa0,
	0x4a, 0x1c, 0xc0, 0xda, 0x38, 0x0e, 0xb2, 0x3b, 0x1a, 0xf1, 0x61, 0x9c, 0x45, 0x5c, 0x64, 0x2d,
	0x7a, 0x36, 0xa8, 0xff, 0xfb, 0x88, 0xfb, 0x9c, 0xa9, 0x03, 0x18, 0x48, 0xed, 0x17, 0x2e, 0xd4,
	0x7f, 0x61, 0x7e, 0xc2, 0xe1, 0x34, 0x66, 0x0f, 0x98, 0x0c, 0x32, 0x00, 0x34, 0x53, 0x1a, 0x75,
	0x65, 0x00, 0x78, 0x4e, 0xa7, 0xf4, 0x21, 0x13, 0x98, 0x4f, 0x82, 0x95, 0xd3, 0xa8, 0x50, 0x04,
	0xf8, 0x2e, 0xe3, 0xe7, 0xaa, 0x4f, 0xcd, 0x66, 0x61, 0x1f, 0xa0, 0x68, 0xec, 0xe5, 0xb9, 0xe8,
	0x63, 0xc7, 0x33, 0x90, 0x7c, 0x5c, 0x8b, 0x95, 0xea, 0xa1, 0x5e, 0x93, 0x63, 0xe8, 0x59, 0xf5,
	0x94, 0x48, 0x07, 0x96, 0x59, 0x16, 0x04, 0x94, 0x31, 0x51, 0x6e, 0xc5, 0x2b, 0x96, 0xc4, 0x03,
	0xbc, 0xa0, 0x8f, 0x2b, 0x90, 0x9c, 0x40, 0xcf, 0xe2, 0x54, 0x22, 0x4c, 0xdd, 0xad, 0x8a, 0xee,
	0x0f, 0xb0, 0x25, 0x9b, 0xfb, 0xb8, 0x4a, 0x06, 0xb0, 0x5d, 0xa5, 0xfd, 0x67, 0x47, 0x26, 0xb0,
	0xfa, 0x22, 0x9b, 0x4e, 0x9a, 0x09, 0xe8, 0xc3, 0xea, 0x75, 0x19, 0xac, 0x86, 0xde, 0x84, 0xf2,
	0xfc, 0x6b, 0x9f, 0x07, 0xb7, 0xa3, 0xf0, 0xbb, 0x34, 0xa5, 0xb6, 0x57, 0x02, 0x24, 0x85, 0xae,
	0x2c, 0x56, 0xf6, 0x28, 0xc9, 0x8c, 0x4b, 0xd6, 0xf6, 0xf4, 0x3a, 0xbf, 0x85, 0x49, 0xc6, 0x5f,
	0xa6, 0x69, 0x9c, 0xca, 0x80, 0x79, 0x11, 0x60, 0x83, 0xb9, 0xa2, 0xb1, 0x38, 0xb2, 0x8c, 0x91,
	0x15, 0x4d, 0x88, 0x8c, 0x60, 0x6d, 0x44, 0xfd, 0x34, 0xb8, 0x6d, 0x6c, 0x4d, 0xcc, 0x0c, 0x2f,
	0xac, 0xc9, 0x02, 0xc9, 0x53, 0x58, 0x2f, 0x48, 0xd5, 0x51, 0x08, 0x74, 0x8b, 0x10, 0x96, 0x4d,
	0x8b, 0x5f, 0x6e, 0x61, 0x83, 0x5f, 0x6d, 0x58, 0xba, 0x14, 0x6f, 0x0e, 0x9e, 0x41, 0x47, 0x3f,
	0x25, 0xb8, 0x23, 0x9f, 0x9e, 0xa3, 0xea, 0x83, 0xe3, 0x3a, 0xf5, 0x0d, 0x59, 0x8e, 0xcc, 0xe1,
	0x2b, 0x58, 0x35, 0xac, 0x1a, 0x77, 0x55, 0x68, 0xfd, 0xa9, 0x71, 0xdd, 0x59, 0x5b, 0x9a, 0xe7,
	0x0c, 0x3a, 0xda, 0x65, 0xb5, 0x92, 0xaa, 0x9f, 0x6b, 0x25, 0x35, 0x43, 0x26, 0x73, 0xf8, 0x1c,
	0x56, 0x0a, 0x0f, 0xc5, 0x6d, 0x15, 0x57, 0x71, 0x61, 0x77, 0xa7, 0x86, 0xeb, 0xf4, 0x21, 0x40,
	0xe9, 0x68, 0x58, 0x14, 0xaa, 0xf9, 0xa2, 0xbb, 0x3b, 0x63, 0xc7, 0xec, 0x86, 0x61, 0x57, 0xba,
	0x1b, 0x75, 0xdb, 0xd3, 0xdd, 0x98, 0xe1, 0x6e, 0x92, 0xc7, 0x70, 0x14, 0xcd, 0x53, 0x77, 0x35,
	0xcd, 0x33, 0xc3, 0x80, 0x24, 0x8f, 0x61, 0x0a, 0x9a, 0xa7, 0x6e, 0x3e, 0x9a, 0x67, 0x86, 0x87,
	0x90, 0x39, 0xbc, 0x82, 0x75, 0xfb, 0x4a, 0xe3, 0x9e, 0xa5, 0xbf, 0xca, 0xf6, 0xe4, 0x0f, 0xbb,
	0x9a, 0xf0, 0x04, 0x16, 0xf3, 0x2b, 0x88, 0xa8, 0x02, 0x8d, 0xcb, 0xef, 0xf6, 0x2c, 0x4c, 0xa7,
	0x3c, 0x83, 0x25, 0x39, 0xec, 0xf8, 0xbf, 0x0a, 0xb0, 0x2e, 0x94, 0xbb, 0x55, 0x41, 0x8b, 0xc4,
	0xeb, 0x25, 0x81, 0x9f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xa2, 0x18, 0xd5, 0x6f, 0x09,
	0x00, 0x00,
}
