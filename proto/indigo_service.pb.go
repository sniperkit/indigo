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
	DeleteIndexRequest
	DeleteIndexResponse
	OpenIndexRequest
	OpenIndexResponse
	CloseIndexRequest
	CloseIndexResponse
	ListIndexRequest
	ListIndexResponse
	GetIndexRequest
	GetIndexResponse
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
	Index        string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	IndexMapping []byte `protobuf:"bytes,2,opt,name=index_mapping,json=indexMapping,proto3" json:"index_mapping,omitempty"`
	IndexType    string `protobuf:"bytes,3,opt,name=index_type,json=indexType" json:"index_type,omitempty"`
	Kvstore      string `protobuf:"bytes,4,opt,name=kvstore" json:"kvstore,omitempty"`
	Kvconfig     []byte `protobuf:"bytes,5,opt,name=kvconfig,proto3" json:"kvconfig,omitempty"`
}

func (m *CreateIndexRequest) Reset()                    { *m = CreateIndexRequest{} }
func (m *CreateIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexRequest) ProtoMessage()               {}
func (*CreateIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateIndexRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index    string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	IndexDir string `protobuf:"bytes,2,opt,name=index_dir,json=indexDir" json:"index_dir,omitempty"`
}

func (m *CreateIndexResponse) Reset()                    { *m = CreateIndexResponse{} }
func (m *CreateIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateIndexResponse) ProtoMessage()               {}
func (*CreateIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateIndexResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *CreateIndexResponse) GetIndexDir() string {
	if m != nil {
		return m.IndexDir
	}
	return ""
}

type DeleteIndexRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *DeleteIndexRequest) Reset()                    { *m = DeleteIndexRequest{} }
func (m *DeleteIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexRequest) ProtoMessage()               {}
func (*DeleteIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteIndexRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type DeleteIndexResponse struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *DeleteIndexResponse) Reset()                    { *m = DeleteIndexResponse{} }
func (m *DeleteIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteIndexResponse) ProtoMessage()               {}
func (*DeleteIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteIndexResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type OpenIndexRequest struct {
	Index         string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	RuntimeConfig []byte `protobuf:"bytes,2,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty"`
}

func (m *OpenIndexRequest) Reset()                    { *m = OpenIndexRequest{} }
func (m *OpenIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexRequest) ProtoMessage()               {}
func (*OpenIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OpenIndexRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index    string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	IndexDir string `protobuf:"bytes,2,opt,name=index_dir,json=indexDir" json:"index_dir,omitempty"`
}

func (m *OpenIndexResponse) Reset()                    { *m = OpenIndexResponse{} }
func (m *OpenIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*OpenIndexResponse) ProtoMessage()               {}
func (*OpenIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OpenIndexResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *OpenIndexResponse) GetIndexDir() string {
	if m != nil {
		return m.IndexDir
	}
	return ""
}

type CloseIndexRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *CloseIndexRequest) Reset()                    { *m = CloseIndexRequest{} }
func (m *CloseIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexRequest) ProtoMessage()               {}
func (*CloseIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CloseIndexRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type CloseIndexResponse struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *CloseIndexResponse) Reset()                    { *m = CloseIndexResponse{} }
func (m *CloseIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*CloseIndexResponse) ProtoMessage()               {}
func (*CloseIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CloseIndexResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type ListIndexRequest struct {
}

func (m *ListIndexRequest) Reset()                    { *m = ListIndexRequest{} }
func (m *ListIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexRequest) ProtoMessage()               {}
func (*ListIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListIndexResponse struct {
	Indices []string `protobuf:"bytes,1,rep,name=indices" json:"indices,omitempty"`
}

func (m *ListIndexResponse) Reset()                    { *m = ListIndexResponse{} }
func (m *ListIndexResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListIndexResponse) ProtoMessage()               {}
func (*ListIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListIndexResponse) GetIndices() []string {
	if m != nil {
		return m.Indices
	}
	return nil
}

type GetIndexRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

func (m *GetIndexRequest) Reset()                    { *m = GetIndexRequest{} }
func (m *GetIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetIndexRequest) ProtoMessage()               {}
func (*GetIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetIndexRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
func (*GetIndexResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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

type PutDocumentRequest struct {
	Index  string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Fields []byte `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (m *PutDocumentRequest) Reset()                    { *m = PutDocumentRequest{} }
func (m *PutDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*PutDocumentRequest) ProtoMessage()               {}
func (*PutDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PutDocumentRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetDocumentRequest) Reset()                    { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()               {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetDocumentRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteDocumentRequest) Reset()                    { *m = DeleteDocumentRequest{} }
func (m *DeleteDocumentRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteDocumentRequest) ProtoMessage()               {}
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *DeleteDocumentRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index       string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	BulkRequest []byte `protobuf:"bytes,2,opt,name=bulk_request,json=bulkRequest,proto3" json:"bulk_request,omitempty"`
	BatchSize   int32  `protobuf:"varint,3,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
}

func (m *BulkRequest) Reset()                    { *m = BulkRequest{} }
func (m *BulkRequest) String() string            { return proto1.CompactTextString(m) }
func (*BulkRequest) ProtoMessage()               {}
func (*BulkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *BulkRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	Index         string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	SearchRequest []byte `protobuf:"bytes,2,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto1.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *SearchRequest) GetIndex() string {
	if m != nil {
		return m.Index
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
	proto1.RegisterType((*DeleteIndexRequest)(nil), "proto.DeleteIndexRequest")
	proto1.RegisterType((*DeleteIndexResponse)(nil), "proto.DeleteIndexResponse")
	proto1.RegisterType((*OpenIndexRequest)(nil), "proto.OpenIndexRequest")
	proto1.RegisterType((*OpenIndexResponse)(nil), "proto.OpenIndexResponse")
	proto1.RegisterType((*CloseIndexRequest)(nil), "proto.CloseIndexRequest")
	proto1.RegisterType((*CloseIndexResponse)(nil), "proto.CloseIndexResponse")
	proto1.RegisterType((*ListIndexRequest)(nil), "proto.ListIndexRequest")
	proto1.RegisterType((*ListIndexResponse)(nil), "proto.ListIndexResponse")
	proto1.RegisterType((*GetIndexRequest)(nil), "proto.GetIndexRequest")
	proto1.RegisterType((*GetIndexResponse)(nil), "proto.GetIndexResponse")
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
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	OpenIndex(ctx context.Context, in *OpenIndexRequest, opts ...grpc.CallOption) (*OpenIndexResponse, error)
	CloseIndex(ctx context.Context, in *CloseIndexRequest, opts ...grpc.CallOption) (*CloseIndexResponse, error)
	ListIndex(ctx context.Context, in *ListIndexRequest, opts ...grpc.CallOption) (*ListIndexResponse, error)
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error)
	//    rpc CreateAlias(CreateAliasRequest) returns (CreateAliasResponse) {}
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

func (c *indigoClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/DeleteIndex", in, out, c.cc, opts...)
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

func (c *indigoClient) CloseIndex(ctx context.Context, in *CloseIndexRequest, opts ...grpc.CallOption) (*CloseIndexResponse, error) {
	out := new(CloseIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/CloseIndex", in, out, c.cc, opts...)
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

func (c *indigoClient) GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*GetIndexResponse, error) {
	out := new(GetIndexResponse)
	err := grpc.Invoke(ctx, "/proto.Indigo/GetIndex", in, out, c.cc, opts...)
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
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	OpenIndex(context.Context, *OpenIndexRequest) (*OpenIndexResponse, error)
	CloseIndex(context.Context, *CloseIndexRequest) (*CloseIndexResponse, error)
	ListIndex(context.Context, *ListIndexRequest) (*ListIndexResponse, error)
	GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error)
	//    rpc CreateAlias(CreateAliasRequest) returns (CreateAliasResponse) {}
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
			MethodName: "DeleteIndex",
			Handler:    _Indigo_DeleteIndex_Handler,
		},
		{
			MethodName: "OpenIndex",
			Handler:    _Indigo_OpenIndex_Handler,
		},
		{
			MethodName: "CloseIndex",
			Handler:    _Indigo_CloseIndex_Handler,
		},
		{
			MethodName: "ListIndex",
			Handler:    _Indigo_ListIndex_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _Indigo_GetIndex_Handler,
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
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x6b, 0x4f, 0xdb, 0x3c,
	0x18, 0x6d, 0x0b, 0x2d, 0xed, 0xd3, 0x0b, 0xe0, 0x72, 0x09, 0xe1, 0x45, 0x2f, 0x18, 0xb1, 0x31,
	0xa6, 0x81, 0xc6, 0x34, 0x4d, 0x9a, 0x84, 0x34, 0xad, 0x0c, 0x86, 0xc4, 0xc4, 0x14, 0xf6, 0xbd,
	0x6a, 0x13, 0x03, 0x56, 0x4b, 0x92, 0xc5, 0x4e, 0x35, 0xf8, 0xb2, 0x3f, 0xb3, 0xbf, 0xb7, 0xff,
	0x30, 0xc5, 0x76, 0x4c, 0x2e, 0xa5, 0x54, 0x7c, 0x6a, 0x7d, 0x7c, 0x9e, 0xf3, 0x5c, 0x6c, 0x9f,
	0x80, 0xe9, 0x07, 0x1e, 0xf7, 0x0e, 0xa8, 0xeb, 0xd0, 0x6b, 0xaf, 0xcb, 0x48, 0x30, 0xa2, 0x36,
	0xd9, 0x17, 0x20, 0x2a, 0x8b, 0x1f, 0xfc, 0xa7, 0x08, 0xa8, 0x13, 0x90, 0x1e, 0x27, 0x67, 0xae,
	0x43, 0x7e, 0x59, 0xe4, 0x67, 0x48, 0x18, 0x47, 0x4b, 0x50, 0xa6, 0xd1, 0xda, 0x28, 0x6e, 0x16,
	0x77, 0x6b, 0x96, 0x5c, 0xa0, 0x6d, 0x68, 0x8a, 0x3f, 0xdd, 0xdb, 0x9e, 0xef, 0x53, 0xf7, 0xda,
	0x28, 0x6d, 0x16, 0x77, 0x1b, 0x56, 0x43, 0x80, 0xdf, 0x24, 0x86, 0x36, 0x00, 0x24, 0x89, 0xdf,
	0xf9, 0xc4, 0x98, 0x11, 0xf1, 0x35, 0x81, 0xfc, 0xb8, 0xf3, 0x09, 0x32, 0x60, 0x6e, 0x30, 0x62,
	0xdc, 0x0b, 0x88, 0x31, 0x2b, 0xf6, 0xe2, 0x25, 0x32, 0xa1, 0x3a, 0x18, 0xd9, 0x9e, 0x7b, 0x45,
	0xaf, 0x8d, 0xb2, 0x10, 0xd6, 0x6b, 0xfc, 0x15, 0xda, 0xa9, 0x2a, 0x99, 0xef, 0xb9, 0x8c, 0x3c,
	0x52, 0xe6, 0x3a, 0xc8, 0x7c, 0x5d, 0x87, 0x06, 0xa2, 0xc4, 0x9a, 0x55, 0x15, 0xc0, 0x31, 0x0d,
	0xf0, 0x1e, 0xa0, 0x63, 0x32, 0x24, 0xd3, 0xf4, 0x8b, 0x5f, 0x43, 0x3b, 0xc5, 0x9d, 0x94, 0x15,
	0x5f, 0xc0, 0xc2, 0x85, 0x4f, 0xdc, 0x29, 0xc6, 0xb8, 0x03, 0xad, 0x20, 0x74, 0x39, 0xbd, 0x25,
	0x5d, 0xd5, 0xae, 0x9c, 0x63, 0x53, 0xa1, 0x1d, 0xd9, 0xf3, 0x09, 0x2c, 0x26, 0x04, 0x9f, 0xdf,
	0xf1, 0x2b, 0x58, 0xec, 0x0c, 0x3d, 0x36, 0x4d, 0xc3, 0x7b, 0x80, 0x92, 0xd4, 0x89, 0xfd, 0x22,
	0x58, 0x38, 0xa7, 0x8c, 0x27, 0x55, 0xf1, 0x1b, 0x58, 0x4c, 0x60, 0x2a, 0xdc, 0x80, 0xb9, 0xe8,
	0x06, 0xda, 0x84, 0x19, 0xc5, 0xcd, 0x99, 0xe8, 0xc4, 0xd5, 0x12, 0xbf, 0x84, 0xf9, 0x53, 0xc2,
	0xa7, 0xa8, 0xeb, 0x37, 0x2c, 0x3c, 0x10, 0x95, 0xec, 0x0e, 0xb4, 0x1c, 0xcf, 0x0e, 0x6f, 0x89,
	0xcb, 0xbb, 0xb6, 0x17, 0xba, 0x5c, 0x84, 0xcc, 0x5a, 0xcd, 0x18, 0xed, 0x44, 0x20, 0xfa, 0x1f,
	0xea, 0x72, 0x34, 0x8c, 0xf7, 0x38, 0x53, 0x93, 0x96, 0x37, 0xf4, 0x32, 0x42, 0xf2, 0x97, 0x7a,
	0x26, 0x7f, 0xa9, 0xb1, 0x05, 0xe8, 0x7b, 0xc8, 0x8f, 0x95, 0xf2, 0xe4, 0xe3, 0x6d, 0x41, 0x89,
	0x3a, 0xea, 0x14, 0x4a, 0xd4, 0x41, 0x2b, 0x50, 0xb9, 0xa2, 0x64, 0xe8, 0x30, 0xa5, 0xac, 0x56,
	0xf8, 0x00, 0xda, 0x29, 0xcd, 0x87, 0x71, 0xb1, 0xd0, 0xb6, 0x09, 0x63, 0x42, 0xb6, 0x6a, 0xc5,
	0x4b, 0xfc, 0x11, 0xd0, 0x29, 0x79, 0x5e, 0x11, 0xf8, 0x08, 0xda, 0xa9, 0x58, 0x95, 0x4c, 0xd2,
	0x8a, 0x63, 0x6a, 0x2d, 0xa5, 0x6a, 0x3d, 0x82, 0x65, 0xf9, 0x12, 0x9e, 0x97, 0xfd, 0x10, 0x56,
	0xb2, 0xe1, 0x4f, 0x76, 0x4b, 0xa0, 0xfe, 0x39, 0x1c, 0x0e, 0x26, 0x27, 0xda, 0x82, 0x46, 0x3f,
	0x1c, 0x0e, 0xba, 0x81, 0x64, 0xa9, 0xaa, 0xeb, 0xfd, 0x44, 0xe0, 0x06, 0x40, 0xbf, 0xc7, 0xed,
	0x9b, 0x2e, 0xa3, 0xf7, 0xd2, 0x8f, 0xca, 0x56, 0x4d, 0x20, 0x97, 0xf4, 0x9e, 0xe0, 0x11, 0x34,
	0x64, 0x1a, 0x55, 0xd0, 0x3a, 0xd4, 0xfc, 0x30, 0x79, 0xa3, 0xca, 0x56, 0xd5, 0x0f, 0xd5, 0x65,
	0x7a, 0x01, 0xf3, 0xd1, 0x26, 0x09, 0x02, 0x2f, 0x50, 0x94, 0x92, 0xa0, 0x34, 0xfd, 0x90, 0x7f,
	0x89, 0x50, 0xc9, 0xdb, 0x82, 0x86, 0x23, 0xfa, 0x55, 0x24, 0x99, 0xb5, 0x2e, 0x31, 0x41, 0xc1,
	0xe7, 0xd0, 0xbc, 0x24, 0xbd, 0xc0, 0xbe, 0x79, 0xd2, 0x2b, 0x98, 0xa0, 0x65, 0x5a, 0x6c, 0xb2,
	0x64, 0x30, 0x7e, 0x0f, 0xad, 0x58, 0x4d, 0xf5, 0xb1, 0x0d, 0x4d, 0x1d, 0xc8, 0xc2, 0xa1, 0xec,
	0xa5, 0x61, 0x35, 0xe2, 0xb8, 0x08, 0x3b, 0xfc, 0x5b, 0x86, 0xca, 0x99, 0xf8, 0x3a, 0xa0, 0x13,
	0xa8, 0x27, 0x1c, 0x16, 0xad, 0xc9, 0xcf, 0xc4, 0x7e, 0xfe, 0xdb, 0x60, 0x9a, 0xe3, 0xb6, 0x64,
	0x56, 0x5c, 0x88, 0x74, 0x12, 0x9e, 0xa9, 0x75, 0xf2, 0x9e, 0xab, 0x75, 0xc6, 0x58, 0x2c, 0x2e,
	0xa0, 0x4f, 0x50, 0xd3, 0xee, 0x87, 0x56, 0x15, 0x35, 0x6b, 0xb0, 0xa6, 0x91, 0xdf, 0xd0, 0x0a,
	0x1d, 0x80, 0x07, 0x33, 0x43, 0x31, 0x33, 0x67, 0x85, 0xe6, 0xda, 0x98, 0x9d, 0x64, 0x19, 0xda,
	0xd1, 0x74, 0x19, 0x59, 0xdf, 0xd3, 0x65, 0xe4, 0xcc, 0x0f, 0x17, 0xd0, 0x11, 0x54, 0x63, 0xef,
	0x42, 0x2b, 0x8a, 0x97, 0x71, 0x3d, 0x73, 0x35, 0x87, 0x27, 0xe7, 0x99, 0x70, 0x09, 0x3d, 0xcf,
	0xbc, 0x1b, 0xe9, 0x79, 0x8e, 0x31, 0x15, 0xa9, 0x93, 0x30, 0x00, 0xad, 0x93, 0x37, 0x14, 0xad,
	0x33, 0xc6, 0x2f, 0x70, 0x01, 0x5d, 0x40, 0x2b, 0xfd, 0x94, 0xd1, 0x7f, 0xa9, 0x73, 0xcc, 0xaa,
	0x6d, 0x3c, 0xb2, 0xab, 0x05, 0xdf, 0xc2, 0x6c, 0xf4, 0x00, 0x11, 0x52, 0xc4, 0xc4, 0xa3, 0x37,
	0xdb, 0x29, 0x4c, 0x87, 0x7c, 0x80, 0x8a, 0xbc, 0xed, 0x68, 0x49, 0x11, 0x52, 0x4f, 0xc9, 0x5c,
	0xce, 0xa0, 0x71, 0x60, 0xbf, 0x22, 0xf0, 0x77, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x63,
	0xba, 0x37, 0x19, 0x09, 0x00, 0x00,
}
