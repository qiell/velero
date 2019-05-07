// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ObjectStore.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutObjectRequest) Reset()         { *m = PutObjectRequest{} }
func (m *PutObjectRequest) String() string { return proto.CompactTextString(m) }
func (*PutObjectRequest) ProtoMessage()    {}
func (*PutObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{0}
}
func (m *PutObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectRequest.Unmarshal(m, b)
}
func (m *PutObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectRequest.Marshal(b, m, deterministic)
}
func (dst *PutObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectRequest.Merge(dst, src)
}
func (m *PutObjectRequest) XXX_Size() int {
	return xxx_messageInfo_PutObjectRequest.Size(m)
}
func (m *PutObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectRequest proto.InternalMessageInfo

func (m *PutObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *PutObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *PutObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutObjectRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{1}
}
func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (dst *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(dst, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *GetObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *GetObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Bytes struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{2}
}
func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (dst *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(dst, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListCommonPrefixesRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Delimiter            string   `protobuf:"bytes,3,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Prefix               string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCommonPrefixesRequest) Reset()         { *m = ListCommonPrefixesRequest{} }
func (m *ListCommonPrefixesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCommonPrefixesRequest) ProtoMessage()    {}
func (*ListCommonPrefixesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{3}
}
func (m *ListCommonPrefixesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommonPrefixesRequest.Unmarshal(m, b)
}
func (m *ListCommonPrefixesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommonPrefixesRequest.Marshal(b, m, deterministic)
}
func (dst *ListCommonPrefixesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommonPrefixesRequest.Merge(dst, src)
}
func (m *ListCommonPrefixesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCommonPrefixesRequest.Size(m)
}
func (m *ListCommonPrefixesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommonPrefixesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommonPrefixesRequest proto.InternalMessageInfo

func (m *ListCommonPrefixesRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type ListCommonPrefixesResponse struct {
	Prefixes             []string `protobuf:"bytes,1,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCommonPrefixesResponse) Reset()         { *m = ListCommonPrefixesResponse{} }
func (m *ListCommonPrefixesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCommonPrefixesResponse) ProtoMessage()    {}
func (*ListCommonPrefixesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{4}
}
func (m *ListCommonPrefixesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommonPrefixesResponse.Unmarshal(m, b)
}
func (m *ListCommonPrefixesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommonPrefixesResponse.Marshal(b, m, deterministic)
}
func (dst *ListCommonPrefixesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommonPrefixesResponse.Merge(dst, src)
}
func (m *ListCommonPrefixesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCommonPrefixesResponse.Size(m)
}
func (m *ListCommonPrefixesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommonPrefixesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommonPrefixesResponse proto.InternalMessageInfo

func (m *ListCommonPrefixesResponse) GetPrefixes() []string {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type ListObjectsRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListObjectsRequest) Reset()         { *m = ListObjectsRequest{} }
func (m *ListObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListObjectsRequest) ProtoMessage()    {}
func (*ListObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{5}
}
func (m *ListObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsRequest.Unmarshal(m, b)
}
func (m *ListObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsRequest.Marshal(b, m, deterministic)
}
func (dst *ListObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsRequest.Merge(dst, src)
}
func (m *ListObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListObjectsRequest.Size(m)
}
func (m *ListObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsRequest proto.InternalMessageInfo

func (m *ListObjectsRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *ListObjectsRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ListObjectsRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type ListObjectsResponse struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListObjectsResponse) Reset()         { *m = ListObjectsResponse{} }
func (m *ListObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListObjectsResponse) ProtoMessage()    {}
func (*ListObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{6}
}
func (m *ListObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsResponse.Unmarshal(m, b)
}
func (m *ListObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsResponse.Marshal(b, m, deterministic)
}
func (dst *ListObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsResponse.Merge(dst, src)
}
func (m *ListObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListObjectsResponse.Size(m)
}
func (m *ListObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsResponse proto.InternalMessageInfo

func (m *ListObjectsResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectRequest) Reset()         { *m = DeleteObjectRequest{} }
func (m *DeleteObjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectRequest) ProtoMessage()    {}
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{7}
}
func (m *DeleteObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectRequest.Unmarshal(m, b)
}
func (m *DeleteObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectRequest.Merge(dst, src)
}
func (m *DeleteObjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectRequest.Size(m)
}
func (m *DeleteObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectRequest proto.InternalMessageInfo

func (m *DeleteObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *DeleteObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *DeleteObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type CreateSignedURLRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Ttl                  int64    `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSignedURLRequest) Reset()         { *m = CreateSignedURLRequest{} }
func (m *CreateSignedURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSignedURLRequest) ProtoMessage()    {}
func (*CreateSignedURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{8}
}
func (m *CreateSignedURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSignedURLRequest.Unmarshal(m, b)
}
func (m *CreateSignedURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSignedURLRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSignedURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSignedURLRequest.Merge(dst, src)
}
func (m *CreateSignedURLRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSignedURLRequest.Size(m)
}
func (m *CreateSignedURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSignedURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSignedURLRequest proto.InternalMessageInfo

func (m *CreateSignedURLRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *CreateSignedURLRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *CreateSignedURLRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateSignedURLRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type CreateSignedURLResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSignedURLResponse) Reset()         { *m = CreateSignedURLResponse{} }
func (m *CreateSignedURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSignedURLResponse) ProtoMessage()    {}
func (*CreateSignedURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ObjectStore_c3671bb7cc2cbc44, []int{9}
}
func (m *CreateSignedURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSignedURLResponse.Unmarshal(m, b)
}
func (m *CreateSignedURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSignedURLResponse.Marshal(b, m, deterministic)
}
func (dst *CreateSignedURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSignedURLResponse.Merge(dst, src)
}
func (m *CreateSignedURLResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSignedURLResponse.Size(m)
}
func (m *CreateSignedURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSignedURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSignedURLResponse proto.InternalMessageInfo

func (m *CreateSignedURLResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*PutObjectRequest)(nil), "generated.PutObjectRequest")
	proto.RegisterType((*GetObjectRequest)(nil), "generated.GetObjectRequest")
	proto.RegisterType((*Bytes)(nil), "generated.Bytes")
	proto.RegisterType((*ListCommonPrefixesRequest)(nil), "generated.ListCommonPrefixesRequest")
	proto.RegisterType((*ListCommonPrefixesResponse)(nil), "generated.ListCommonPrefixesResponse")
	proto.RegisterType((*ListObjectsRequest)(nil), "generated.ListObjectsRequest")
	proto.RegisterType((*ListObjectsResponse)(nil), "generated.ListObjectsResponse")
	proto.RegisterType((*DeleteObjectRequest)(nil), "generated.DeleteObjectRequest")
	proto.RegisterType((*CreateSignedURLRequest)(nil), "generated.CreateSignedURLRequest")
	proto.RegisterType((*CreateSignedURLResponse)(nil), "generated.CreateSignedURLResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ObjectStoreClient is the client API for ObjectStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObjectStoreClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	PutObject(ctx context.Context, opts ...grpc.CallOption) (ObjectStore_PutObjectClient, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (ObjectStore_GetObjectClient, error)
	ListCommonPrefixes(ctx context.Context, in *ListCommonPrefixesRequest, opts ...grpc.CallOption) (*ListCommonPrefixesResponse, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateSignedURL(ctx context.Context, in *CreateSignedURLRequest, opts ...grpc.CallOption) (*CreateSignedURLResponse, error)
}

type objectStoreClient struct {
	cc *grpc.ClientConn
}

func NewObjectStoreClient(cc *grpc.ClientConn) ObjectStoreClient {
	return &objectStoreClient{cc}
}

func (c *objectStoreClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) PutObject(ctx context.Context, opts ...grpc.CallOption) (ObjectStore_PutObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ObjectStore_serviceDesc.Streams[0], "/generated.ObjectStore/PutObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectStorePutObjectClient{stream}
	return x, nil
}

type ObjectStore_PutObjectClient interface {
	Send(*PutObjectRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type objectStorePutObjectClient struct {
	grpc.ClientStream
}

func (x *objectStorePutObjectClient) Send(m *PutObjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *objectStorePutObjectClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectStoreClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (ObjectStore_GetObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ObjectStore_serviceDesc.Streams[1], "/generated.ObjectStore/GetObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectStoreGetObjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectStore_GetObjectClient interface {
	Recv() (*Bytes, error)
	grpc.ClientStream
}

type objectStoreGetObjectClient struct {
	grpc.ClientStream
}

func (x *objectStoreGetObjectClient) Recv() (*Bytes, error) {
	m := new(Bytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectStoreClient) ListCommonPrefixes(ctx context.Context, in *ListCommonPrefixesRequest, opts ...grpc.CallOption) (*ListCommonPrefixesResponse, error) {
	out := new(ListCommonPrefixesResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/ListCommonPrefixes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) CreateSignedURL(ctx context.Context, in *CreateSignedURLRequest, opts ...grpc.CallOption) (*CreateSignedURLResponse, error) {
	out := new(CreateSignedURLResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/CreateSignedURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStoreServer is the server API for ObjectStore service.
type ObjectStoreServer interface {
	Init(context.Context, *InitRequest) (*Empty, error)
	PutObject(ObjectStore_PutObjectServer) error
	GetObject(*GetObjectRequest, ObjectStore_GetObjectServer) error
	ListCommonPrefixes(context.Context, *ListCommonPrefixesRequest) (*ListCommonPrefixesResponse, error)
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*Empty, error)
	CreateSignedURL(context.Context, *CreateSignedURLRequest) (*CreateSignedURLResponse, error)
}

func RegisterObjectStoreServer(s *grpc.Server, srv ObjectStoreServer) {
	s.RegisterService(&_ObjectStore_serviceDesc, srv)
}

func _ObjectStore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_PutObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ObjectStoreServer).PutObject(&objectStorePutObjectServer{stream})
}

type ObjectStore_PutObjectServer interface {
	SendAndClose(*Empty) error
	Recv() (*PutObjectRequest, error)
	grpc.ServerStream
}

type objectStorePutObjectServer struct {
	grpc.ServerStream
}

func (x *objectStorePutObjectServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *objectStorePutObjectServer) Recv() (*PutObjectRequest, error) {
	m := new(PutObjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ObjectStore_GetObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetObjectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectStoreServer).GetObject(m, &objectStoreGetObjectServer{stream})
}

type ObjectStore_GetObjectServer interface {
	Send(*Bytes) error
	grpc.ServerStream
}

type objectStoreGetObjectServer struct {
	grpc.ServerStream
}

func (x *objectStoreGetObjectServer) Send(m *Bytes) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectStore_ListCommonPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommonPrefixesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).ListCommonPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/ListCommonPrefixes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).ListCommonPrefixes(ctx, req.(*ListCommonPrefixesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_CreateSignedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSignedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).CreateSignedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/CreateSignedURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).CreateSignedURL(ctx, req.(*CreateSignedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.ObjectStore",
	HandlerType: (*ObjectStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ObjectStore_Init_Handler,
		},
		{
			MethodName: "ListCommonPrefixes",
			Handler:    _ObjectStore_ListCommonPrefixes_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _ObjectStore_ListObjects_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _ObjectStore_DeleteObject_Handler,
		},
		{
			MethodName: "CreateSignedURL",
			Handler:    _ObjectStore_CreateSignedURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutObject",
			Handler:       _ObjectStore_PutObject_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetObject",
			Handler:       _ObjectStore_GetObject_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ObjectStore.proto",
}

func init() { proto.RegisterFile("ObjectStore.proto", fileDescriptor_ObjectStore_c3671bb7cc2cbc44) }

var fileDescriptor_ObjectStore_c3671bb7cc2cbc44 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x26, 0x26, 0x16, 0xf3, 0xba, 0x60, 0x9c, 0x42, 0x8d, 0xa9, 0x8a, 0x0e, 0x0a, 0x15, 0x61,
	0x29, 0x7a, 0xf1, 0xe0, 0x41, 0xac, 0x22, 0xc2, 0x82, 0x25, 0xab, 0xe8, 0xc1, 0x4b, 0xd2, 0x3c,
	0xb7, 0xb1, 0xf9, 0xe5, 0x64, 0x02, 0xe6, 0xe8, 0xcd, 0x3f, 0xdb, 0x99, 0xc9, 0x98, 0x4e, 0xb2,
	0x59, 0x85, 0x65, 0x6f, 0x6f, 0xbe, 0x79, 0x3f, 0xbe, 0xbc, 0xf9, 0xbe, 0xc0, 0xad, 0x0f, 0xf1,
	0x77, 0x3c, 0xe7, 0x4b, 0x5e, 0x32, 0x9c, 0x57, 0xac, 0xe4, 0x25, 0x71, 0x57, 0x58, 0x20, 0x8b,
	0x38, 0x26, 0xc1, 0x6c, 0x79, 0x11, 0x31, 0x4c, 0xba, 0x0b, 0x7a, 0x01, 0xde, 0x59, 0xc3, 0xbb,
	0x82, 0x10, 0x7f, 0x34, 0x58, 0x73, 0x72, 0x08, 0x7b, 0x55, 0xd6, 0xac, 0xd2, 0xc2, 0xb7, 0x1e,
	0x58, 0xc7, 0x6e, 0xa8, 0x4f, 0x12, 0x8f, 0x9b, 0xf3, 0x4b, 0xe4, 0xfe, 0xb5, 0x0e, 0xef, 0x4e,
	0xc4, 0x03, 0xfb, 0x12, 0x5b, 0xdf, 0x56, 0xa0, 0x0c, 0x09, 0x01, 0x27, 0x2e, 0x93, 0xd6, 0x77,
	0x04, 0x34, 0x0b, 0x55, 0x4c, 0x3f, 0x82, 0xf7, 0x0e, 0x77, 0x3d, 0x89, 0x1e, 0xc1, 0xf5, 0xd7,
	0x2d, 0xc7, 0x5a, 0x8e, 0x4c, 0x22, 0x1e, 0xa9, 0x46, 0x62, 0xa4, 0x8c, 0xe9, 0x2f, 0x0b, 0xee,
	0x2c, 0xd2, 0x9a, 0x9f, 0x96, 0x79, 0x5e, 0x16, 0x67, 0x0c, 0xbf, 0xa5, 0x3f, 0xb1, 0xde, 0x76,
	0xf8, 0x5d, 0x70, 0x13, 0xcc, 0xd2, 0x3c, 0xe5, 0xc8, 0x34, 0x85, 0x2b, 0x40, 0x75, 0x53, 0x03,
	0xd4, 0x47, 0xcb, 0x6e, 0xea, 0x44, 0x5f, 0x40, 0x30, 0x45, 0xa1, 0xae, 0xca, 0xa2, 0x46, 0x12,
	0xc0, 0x8d, 0x4a, 0x63, 0x82, 0x85, 0x2d, 0xea, 0xfa, 0x33, 0xfd, 0x0a, 0x44, 0x56, 0x76, 0x1b,
	0xdb, 0x9a, 0xf5, 0x15, 0x2f, 0x7b, 0xc0, 0xeb, 0x09, 0x1c, 0x0c, 0xba, 0x6b, 0x42, 0x62, 0x8d,
	0x62, 0xad, 0x7f, 0xc9, 0xa8, 0x98, 0x7e, 0x86, 0x83, 0x37, 0x98, 0x21, 0xc7, 0x5d, 0x3f, 0x5e,
	0x06, 0x87, 0xa7, 0x0c, 0x85, 0x2a, 0x97, 0xe9, 0xaa, 0xc0, 0xe4, 0x53, 0xb8, 0xd8, 0x9d, 0x04,
	0x05, 0xc2, 0x79, 0xa6, 0x1e, 0xc3, 0x0e, 0x65, 0x48, 0x9f, 0xc2, 0xed, 0xb5, 0x69, 0xfa, 0xab,
	0x45, 0x72, 0xc3, 0x32, 0x3d, 0x4b, 0x86, 0xcf, 0x7e, 0x3b, 0xb0, 0x6f, 0xd8, 0x88, 0x9c, 0x80,
	0xf3, 0xbe, 0x48, 0x05, 0xb1, 0x79, 0xef, 0xa4, 0xb9, 0x04, 0x34, 0xe1, 0xc0, 0x33, 0xf0, 0xb7,
	0x79, 0xc5, 0x5b, 0xf2, 0x12, 0xdc, 0xde, 0x59, 0xe4, 0xc8, 0xb8, 0x1e, 0xfb, 0x6d, 0xbd, 0xf6,
	0xd8, 0x92, 0xd5, 0xbd, 0x5b, 0x06, 0xd5, 0x63, 0x0f, 0x0d, 0xaa, 0x95, 0x15, 0x4e, 0x2c, 0x12,
	0x75, 0xd2, 0x19, 0x8a, 0x8e, 0x3c, 0x32, 0x32, 0x37, 0xda, 0x22, 0x78, 0xfc, 0x9f, 0x2c, 0xbd,
	0xb2, 0x05, 0xec, 0x1b, 0xfa, 0x21, 0xf7, 0x46, 0x55, 0x43, 0xd5, 0x06, 0xf7, 0x37, 0x5d, 0xeb,
	0x6e, 0xaf, 0x60, 0x66, 0x4a, 0x8c, 0x98, 0xf9, 0x13, 0xda, 0x9b, 0x58, 0xf7, 0x17, 0xb8, 0x39,
	0x7a, 0x5d, 0xf2, 0xd0, 0x48, 0x9a, 0xd6, 0x59, 0x40, 0xff, 0x95, 0xd2, 0x71, 0x8b, 0xf7, 0xd4,
	0x9f, 0xf2, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xac, 0xfe, 0xa7, 0x57, 0x05, 0x00,
	0x00,
}
