// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go.
// source: operator.proto
// DO NOT EDIT!

/*
Package m3em is a generated protocol buffer package.

It is generated from these files:
	operator.proto

It has these top-level messages:
	SetupRequest
	SetupResponse
	StartRequest
	StartResponse
	StopRequest
	StopResponse
	TeardownRequest
	TeardownResponse
	PullFileRequest
	PullFileResponse
	PushFileRequest
	PushFileResponse
	DataChunk
*/
package m3em

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

type PullFileType int32

const (
	PullFileType_PULL_FILE_TYPE_UNKNOWN        PullFileType = 0
	PullFileType_PULL_FILE_TYPE_SERVICE_STDOUT PullFileType = 1
	PullFileType_PULL_FILE_TYPE_SERVICE_STDERR PullFileType = 2
)

var PullFileType_name = map[int32]string{
	0: "PULL_FILE_TYPE_UNKNOWN",
	1: "PULL_FILE_TYPE_SERVICE_STDOUT",
	2: "PULL_FILE_TYPE_SERVICE_STDERR",
}
var PullFileType_value = map[string]int32{
	"PULL_FILE_TYPE_UNKNOWN":        0,
	"PULL_FILE_TYPE_SERVICE_STDOUT": 1,
	"PULL_FILE_TYPE_SERVICE_STDERR": 2,
}

func (x PullFileType) String() string {
	return proto.EnumName(PullFileType_name, int32(x))
}
func (PullFileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PullFileContentType int32

const (
	PullFileContentType_PULL_FILE_CONTENT_TYPE_UNKNOWN PullFileContentType = 0
	PullFileContentType_PULL_FILE_CONTENT_TYPE_LAST    PullFileContentType = 1
	PullFileContentType_PULL_FILE_CONTENT_TYPE_FULL    PullFileContentType = 2
)

var PullFileContentType_name = map[int32]string{
	0: "PULL_FILE_CONTENT_TYPE_UNKNOWN",
	1: "PULL_FILE_CONTENT_TYPE_LAST",
	2: "PULL_FILE_CONTENT_TYPE_FULL",
}
var PullFileContentType_value = map[string]int32{
	"PULL_FILE_CONTENT_TYPE_UNKNOWN": 0,
	"PULL_FILE_CONTENT_TYPE_LAST":    1,
	"PULL_FILE_CONTENT_TYPE_FULL":    2,
}

func (x PullFileContentType) String() string {
	return proto.EnumName(PullFileContentType_name, int32(x))
}
func (PullFileContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PushFileType int32

const (
	PushFileType_PUSH_FILE_TYPE_UNKNOWN        PushFileType = 0
	PushFileType_PUSH_FILE_TYPE_SERVICE_BINARY PushFileType = 1
	PushFileType_PUSH_FILE_TYPE_SERVICE_CONFIG PushFileType = 2
	PushFileType_PUSH_FILE_TYPE_DATA_FILE      PushFileType = 3
)

var PushFileType_name = map[int32]string{
	0: "PUSH_FILE_TYPE_UNKNOWN",
	1: "PUSH_FILE_TYPE_SERVICE_BINARY",
	2: "PUSH_FILE_TYPE_SERVICE_CONFIG",
	3: "PUSH_FILE_TYPE_DATA_FILE",
}
var PushFileType_value = map[string]int32{
	"PUSH_FILE_TYPE_UNKNOWN":        0,
	"PUSH_FILE_TYPE_SERVICE_BINARY": 1,
	"PUSH_FILE_TYPE_SERVICE_CONFIG": 2,
	"PUSH_FILE_TYPE_DATA_FILE":      3,
}

func (x PushFileType) String() string {
	return proto.EnumName(PushFileType_name, int32(x))
}
func (PushFileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SetupRequest struct {
	SessionToken           string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken" json:"session_token,omitempty"`
	OperatorUuid           string `protobuf:"bytes,2,opt,name=operator_uuid,json=operatorUuid" json:"operator_uuid,omitempty"`
	Force                  bool   `protobuf:"varint,3,opt,name=force" json:"force,omitempty"`
	HeartbeatEnabled       bool   `protobuf:"varint,4,opt,name=heartbeat_enabled,json=heartbeatEnabled" json:"heartbeat_enabled,omitempty"`
	HeartbeatEndpoint      string `protobuf:"bytes,5,opt,name=heartbeat_endpoint,json=heartbeatEndpoint" json:"heartbeat_endpoint,omitempty"`
	HeartbeatFrequencySecs uint32 `protobuf:"varint,6,opt,name=heartbeat_frequency_secs,json=heartbeatFrequencySecs" json:"heartbeat_frequency_secs,omitempty"`
}

func (m *SetupRequest) Reset()                    { *m = SetupRequest{} }
func (m *SetupRequest) String() string            { return proto.CompactTextString(m) }
func (*SetupRequest) ProtoMessage()               {}
func (*SetupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetupRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *SetupRequest) GetOperatorUuid() string {
	if m != nil {
		return m.OperatorUuid
	}
	return ""
}

func (m *SetupRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *SetupRequest) GetHeartbeatEnabled() bool {
	if m != nil {
		return m.HeartbeatEnabled
	}
	return false
}

func (m *SetupRequest) GetHeartbeatEndpoint() string {
	if m != nil {
		return m.HeartbeatEndpoint
	}
	return ""
}

func (m *SetupRequest) GetHeartbeatFrequencySecs() uint32 {
	if m != nil {
		return m.HeartbeatFrequencySecs
	}
	return 0
}

type SetupResponse struct {
}

func (m *SetupResponse) Reset()                    { *m = SetupResponse{} }
func (m *SetupResponse) String() string            { return proto.CompactTextString(m) }
func (*SetupResponse) ProtoMessage()               {}
func (*SetupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StartRequest struct {
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StartResponse struct {
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StopRequest struct {
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type StopResponse struct {
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type TeardownRequest struct {
}

func (m *TeardownRequest) Reset()                    { *m = TeardownRequest{} }
func (m *TeardownRequest) String() string            { return proto.CompactTextString(m) }
func (*TeardownRequest) ProtoMessage()               {}
func (*TeardownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type TeardownResponse struct {
}

func (m *TeardownResponse) Reset()                    { *m = TeardownResponse{} }
func (m *TeardownResponse) String() string            { return proto.CompactTextString(m) }
func (*TeardownResponse) ProtoMessage()               {}
func (*TeardownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// PullFileRequest(s) are used to transfer files from remote agents to the master.
type PullFileRequest struct {
	FileType  PullFileType `protobuf:"varint,1,opt,name=file_type,json=fileType,enum=m3em.PullFileType" json:"file_type,omitempty"`
	ChunkSize int64        `protobuf:"varint,2,opt,name=chunk_size,json=chunkSize" json:"chunk_size,omitempty"`
	MaxSize   int64        `protobuf:"varint,3,opt,name=max_size,json=maxSize" json:"max_size,omitempty"`
}

func (m *PullFileRequest) Reset()                    { *m = PullFileRequest{} }
func (m *PullFileRequest) String() string            { return proto.CompactTextString(m) }
func (*PullFileRequest) ProtoMessage()               {}
func (*PullFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PullFileRequest) GetFileType() PullFileType {
	if m != nil {
		return m.FileType
	}
	return PullFileType_PULL_FILE_TYPE_UNKNOWN
}

func (m *PullFileRequest) GetChunkSize() int64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *PullFileRequest) GetMaxSize() int64 {
	if m != nil {
		return m.MaxSize
	}
	return 0
}

type PullFileResponse struct {
	Data      *DataChunk `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Truncated bool       `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
}

func (m *PullFileResponse) Reset()                    { *m = PullFileResponse{} }
func (m *PullFileResponse) String() string            { return proto.CompactTextString(m) }
func (*PullFileResponse) ProtoMessage()               {}
func (*PullFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PullFileResponse) GetData() *DataChunk {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PullFileResponse) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// PushFileRequest(s) are used to transfer files from the master to remote agents.
type PushFileRequest struct {
	Type        PushFileType `protobuf:"varint,1,opt,name=type,enum=m3em.PushFileType" json:"type,omitempty"`
	Overwrite   bool         `protobuf:"varint,2,opt,name=overwrite" json:"overwrite,omitempty"`
	Data        *DataChunk   `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	TargetPaths []string     `protobuf:"bytes,4,rep,name=target_paths,json=targetPaths" json:"target_paths,omitempty"`
}

func (m *PushFileRequest) Reset()                    { *m = PushFileRequest{} }
func (m *PushFileRequest) String() string            { return proto.CompactTextString(m) }
func (*PushFileRequest) ProtoMessage()               {}
func (*PushFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PushFileRequest) GetType() PushFileType {
	if m != nil {
		return m.Type
	}
	return PushFileType_PUSH_FILE_TYPE_UNKNOWN
}

func (m *PushFileRequest) GetOverwrite() bool {
	if m != nil {
		return m.Overwrite
	}
	return false
}

func (m *PushFileRequest) GetData() *DataChunk {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PushFileRequest) GetTargetPaths() []string {
	if m != nil {
		return m.TargetPaths
	}
	return nil
}

type PushFileResponse struct {
	FileChecksum   uint32 `protobuf:"varint,1,opt,name=file_checksum,json=fileChecksum" json:"file_checksum,omitempty"`
	NumChunksRecvd int32  `protobuf:"varint,2,opt,name=num_chunks_recvd,json=numChunksRecvd" json:"num_chunks_recvd,omitempty"`
}

func (m *PushFileResponse) Reset()                    { *m = PushFileResponse{} }
func (m *PushFileResponse) String() string            { return proto.CompactTextString(m) }
func (*PushFileResponse) ProtoMessage()               {}
func (*PushFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PushFileResponse) GetFileChecksum() uint32 {
	if m != nil {
		return m.FileChecksum
	}
	return 0
}

func (m *PushFileResponse) GetNumChunksRecvd() int32 {
	if m != nil {
		return m.NumChunksRecvd
	}
	return 0
}

type DataChunk struct {
	Idx   int32  `protobuf:"varint,1,opt,name=idx" json:"idx,omitempty"`
	Bytes []byte `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (m *DataChunk) Reset()                    { *m = DataChunk{} }
func (m *DataChunk) String() string            { return proto.CompactTextString(m) }
func (*DataChunk) ProtoMessage()               {}
func (*DataChunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DataChunk) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *DataChunk) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*SetupRequest)(nil), "m3em.SetupRequest")
	proto.RegisterType((*SetupResponse)(nil), "m3em.SetupResponse")
	proto.RegisterType((*StartRequest)(nil), "m3em.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "m3em.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "m3em.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "m3em.StopResponse")
	proto.RegisterType((*TeardownRequest)(nil), "m3em.TeardownRequest")
	proto.RegisterType((*TeardownResponse)(nil), "m3em.TeardownResponse")
	proto.RegisterType((*PullFileRequest)(nil), "m3em.PullFileRequest")
	proto.RegisterType((*PullFileResponse)(nil), "m3em.PullFileResponse")
	proto.RegisterType((*PushFileRequest)(nil), "m3em.PushFileRequest")
	proto.RegisterType((*PushFileResponse)(nil), "m3em.PushFileResponse")
	proto.RegisterType((*DataChunk)(nil), "m3em.DataChunk")
	proto.RegisterEnum("m3em.PullFileType", PullFileType_name, PullFileType_value)
	proto.RegisterEnum("m3em.PullFileContentType", PullFileContentType_name, PullFileContentType_value)
	proto.RegisterEnum("m3em.PushFileType", PushFileType_name, PushFileType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Operator service

type OperatorClient interface {
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*SetupResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error)
	PullFile(ctx context.Context, in *PullFileRequest, opts ...grpc.CallOption) (Operator_PullFileClient, error)
	PushFile(ctx context.Context, opts ...grpc.CallOption) (Operator_PushFileClient, error)
}

type operatorClient struct {
	cc *grpc.ClientConn
}

func NewOperatorClient(cc *grpc.ClientConn) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*SetupResponse, error) {
	out := new(SetupResponse)
	err := grpc.Invoke(ctx, "/m3em.Operator/Setup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/m3em.Operator/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/m3em.Operator/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*TeardownResponse, error) {
	out := new(TeardownResponse)
	err := grpc.Invoke(ctx, "/m3em.Operator/Teardown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) PullFile(ctx context.Context, in *PullFileRequest, opts ...grpc.CallOption) (Operator_PullFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Operator_serviceDesc.Streams[0], c.cc, "/m3em.Operator/PullFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &operatorPullFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Operator_PullFileClient interface {
	Recv() (*PullFileResponse, error)
	grpc.ClientStream
}

type operatorPullFileClient struct {
	grpc.ClientStream
}

func (x *operatorPullFileClient) Recv() (*PullFileResponse, error) {
	m := new(PullFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *operatorClient) PushFile(ctx context.Context, opts ...grpc.CallOption) (Operator_PushFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Operator_serviceDesc.Streams[1], c.cc, "/m3em.Operator/PushFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &operatorPushFileClient{stream}
	return x, nil
}

type Operator_PushFileClient interface {
	Send(*PushFileRequest) error
	CloseAndRecv() (*PushFileResponse, error)
	grpc.ClientStream
}

type operatorPushFileClient struct {
	grpc.ClientStream
}

func (x *operatorPushFileClient) Send(m *PushFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *operatorPushFileClient) CloseAndRecv() (*PushFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Operator service

type OperatorServer interface {
	Setup(context.Context, *SetupRequest) (*SetupResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	Teardown(context.Context, *TeardownRequest) (*TeardownResponse, error)
	PullFile(*PullFileRequest, Operator_PullFileServer) error
	PushFile(Operator_PushFileServer) error
}

func RegisterOperatorServer(s *grpc.Server, srv OperatorServer) {
	s.RegisterService(&_Operator_serviceDesc, srv)
}

func _Operator_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m3em.Operator/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Setup(ctx, req.(*SetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m3em.Operator/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m3em.Operator/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m3em.Operator/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Teardown(ctx, req.(*TeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_PullFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperatorServer).PullFile(m, &operatorPullFileServer{stream})
}

type Operator_PullFileServer interface {
	Send(*PullFileResponse) error
	grpc.ServerStream
}

type operatorPullFileServer struct {
	grpc.ServerStream
}

func (x *operatorPullFileServer) Send(m *PullFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Operator_PushFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OperatorServer).PushFile(&operatorPushFileServer{stream})
}

type Operator_PushFileServer interface {
	SendAndClose(*PushFileResponse) error
	Recv() (*PushFileRequest, error)
	grpc.ServerStream
}

type operatorPushFileServer struct {
	grpc.ServerStream
}

func (x *operatorPushFileServer) SendAndClose(m *PushFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *operatorPushFileServer) Recv() (*PushFileRequest, error) {
	m := new(PushFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Operator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "m3em.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Operator_Setup_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Operator_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Operator_Stop_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _Operator_Teardown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullFile",
			Handler:       _Operator_PullFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PushFile",
			Handler:       _Operator_PushFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "operator.proto",
}

func init() { proto.RegisterFile("operator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x55, 0x5d, 0x8f, 0xda, 0x46,
	0x14, 0xad, 0xf9, 0x48, 0xe1, 0x2e, 0xec, 0x7a, 0x27, 0xed, 0xca, 0xa5, 0x49, 0x4b, 0x1c, 0xa9,
	0x42, 0x5b, 0x65, 0x1b, 0xed, 0xbe, 0xb4, 0xea, 0x13, 0x65, 0xa1, 0x45, 0x45, 0x80, 0xc6, 0xa6,
	0x55, 0x9e, 0xac, 0xc1, 0x0c, 0xc5, 0x5a, 0xb0, 0x5d, 0xcf, 0x38, 0x59, 0xd2, 0x3e, 0xf6, 0xb9,
	0x7f, 0xa1, 0xff, 0xa9, 0xbf, 0x28, 0x9a, 0x19, 0x8f, 0x6d, 0x50, 0x76, 0xdf, 0x98, 0x73, 0xce,
	0x9d, 0x7b, 0xcf, 0xbd, 0x9e, 0x0b, 0x9c, 0x46, 0x31, 0x4d, 0x08, 0x8f, 0x92, 0xab, 0x38, 0x89,
	0x78, 0x84, 0x6a, 0xbb, 0x1b, 0xba, 0xb3, 0xff, 0xa9, 0x40, 0xcb, 0xa1, 0x3c, 0x8d, 0x31, 0xfd,
	0x33, 0xa5, 0x8c, 0xa3, 0x97, 0xd0, 0x66, 0x94, 0xb1, 0x20, 0x0a, 0x3d, 0x1e, 0xdd, 0xd1, 0xd0,
	0x32, 0xba, 0x46, 0xaf, 0x89, 0x5b, 0x19, 0xe8, 0x0a, 0x4c, 0x88, 0xf4, 0x6d, 0x5e, 0x9a, 0x06,
	0x2b, 0xab, 0xa2, 0x44, 0x1a, 0x5c, 0xa4, 0xc1, 0x0a, 0x7d, 0x06, 0xf5, 0x75, 0x94, 0xf8, 0xd4,
	0xaa, 0x76, 0x8d, 0x5e, 0x03, 0xab, 0x03, 0xfa, 0x16, 0xce, 0x37, 0x94, 0x24, 0x7c, 0x49, 0x09,
	0xf7, 0x68, 0x48, 0x96, 0x5b, 0xba, 0xb2, 0x6a, 0x52, 0x61, 0xe6, 0xc4, 0x50, 0xe1, 0xe8, 0x15,
	0xa0, 0xb2, 0x78, 0x15, 0x47, 0x41, 0xc8, 0xad, 0xba, 0x4c, 0x76, 0x5e, 0x52, 0x2b, 0x02, 0x7d,
	0x0f, 0x56, 0x21, 0x5f, 0x27, 0xc2, 0x51, 0xe8, 0xef, 0x3d, 0x46, 0x7d, 0x66, 0x3d, 0xe9, 0x1a,
	0xbd, 0x36, 0xbe, 0xc8, 0xf9, 0x91, 0xa6, 0x1d, 0xea, 0x33, 0xfb, 0x0c, 0xda, 0x59, 0x17, 0x58,
	0x1c, 0x85, 0x8c, 0xda, 0xa7, 0xd0, 0x72, 0x38, 0x49, 0x78, 0xd6, 0x16, 0x29, 0x50, 0xe7, 0x4c,
	0xd0, 0x86, 0x13, 0x87, 0x47, 0xba, 0x6d, 0x4a, 0x1f, 0x15, 0xf1, 0xe7, 0x70, 0xe6, 0x52, 0x92,
	0xac, 0xa2, 0x77, 0xa1, 0x96, 0x20, 0x30, 0x0b, 0x28, 0x93, 0xfd, 0x0d, 0x67, 0xf3, 0x74, 0xbb,
	0x1d, 0x05, 0x5b, 0xaa, 0x07, 0xf0, 0x1d, 0x34, 0xd7, 0xc1, 0x96, 0x7a, 0x7c, 0x1f, 0x53, 0xd9,
	0xfc, 0xd3, 0x6b, 0x74, 0x25, 0x66, 0x75, 0xa5, 0x95, 0xee, 0x3e, 0xa6, 0xb8, 0xb1, 0xce, 0x7e,
	0xa1, 0xe7, 0x00, 0xfe, 0x26, 0x0d, 0xef, 0x3c, 0x16, 0xbc, 0xa7, 0x72, 0x12, 0x55, 0xdc, 0x94,
	0x88, 0x13, 0xbc, 0xa7, 0xe8, 0x0b, 0x68, 0xec, 0xc8, 0xbd, 0x22, 0xab, 0x92, 0xfc, 0x74, 0x47,
	0xee, 0x05, 0x65, 0x2f, 0xc0, 0x2c, 0xb2, 0xab, 0x8a, 0xd0, 0x4b, 0xa8, 0xad, 0x08, 0x27, 0x32,
	0xf3, 0xc9, 0xf5, 0x99, 0xca, 0x7c, 0x4b, 0x38, 0x19, 0x88, 0x1b, 0xb1, 0x24, 0xd1, 0x33, 0x68,
	0xf2, 0x24, 0x0d, 0x7d, 0xc2, 0xa9, 0x9a, 0x7d, 0x03, 0x17, 0x80, 0xfd, 0x9f, 0x21, 0x5c, 0xb1,
	0x4d, 0xd9, 0xd5, 0x37, 0x50, 0xfb, 0x98, 0x21, 0x25, 0x92, 0x86, 0x24, 0x2f, 0x6e, 0x8e, 0xde,
	0xd2, 0xe4, 0x5d, 0x12, 0x70, 0xaa, 0x6f, 0xce, 0x81, 0xbc, 0xb8, 0xea, 0x63, 0xc5, 0xbd, 0x80,
	0x16, 0x27, 0xc9, 0x1f, 0x94, 0x7b, 0x31, 0xe1, 0x1b, 0x66, 0xd5, 0xba, 0xd5, 0x5e, 0x13, 0x9f,
	0x28, 0x6c, 0x2e, 0x20, 0x9b, 0x08, 0xe3, 0xba, 0xc0, 0xdc, 0x78, 0x5b, 0xf6, 0xdd, 0xdf, 0x50,
	0xff, 0x8e, 0xa5, 0x3b, 0x59, 0x6a, 0x1b, 0xb7, 0x04, 0x38, 0xc8, 0x30, 0xd4, 0x03, 0x33, 0x4c,
	0x77, 0x9e, 0xec, 0x2e, 0xf3, 0x12, 0xea, 0xbf, 0x55, 0xfe, 0xeb, 0xf8, 0x34, 0x4c, 0x77, 0xb2,
	0x0a, 0x86, 0x05, 0x6a, 0xdf, 0x40, 0x33, 0x2f, 0x0c, 0x99, 0x50, 0x0d, 0x56, 0xf7, 0xf2, 0xc6,
	0x3a, 0x16, 0x3f, 0xc5, 0xe3, 0x58, 0xee, 0x39, 0x65, 0x32, 0xba, 0x85, 0xd5, 0xe1, 0x32, 0x86,
	0x56, 0x79, 0xc8, 0xa8, 0x03, 0x17, 0xf3, 0xc5, 0x64, 0xe2, 0x8d, 0xc6, 0x93, 0xa1, 0xe7, 0xbe,
	0x99, 0x0f, 0xbd, 0xc5, 0xf4, 0xd7, 0xe9, 0xec, 0xf7, 0xa9, 0xf9, 0x09, 0x7a, 0x01, 0xcf, 0x8f,
	0x38, 0x67, 0x88, 0x7f, 0x1b, 0x0f, 0x86, 0x9e, 0xe3, 0xde, 0xce, 0x16, 0xae, 0x69, 0x3c, 0x2e,
	0x19, 0x62, 0x6c, 0x56, 0x2e, 0xff, 0x82, 0xa7, 0x3a, 0xe3, 0x20, 0x0a, 0x39, 0x0d, 0xb9, 0x4c,
	0x6c, 0xc3, 0x57, 0x45, 0xe4, 0x60, 0x36, 0x75, 0x87, 0x53, 0xf7, 0xb8, 0x80, 0xaf, 0xe1, 0xcb,
	0x07, 0x34, 0x93, 0xbe, 0x23, 0xd2, 0x3f, 0x2c, 0x18, 0x2d, 0x26, 0x13, 0xb3, 0x72, 0xf9, 0xaf,
	0x21, 0xfc, 0x16, 0xdf, 0x80, 0xf2, 0xeb, 0xfc, 0xf2, 0xb0, 0xdf, 0x03, 0x4e, 0x9b, 0xf9, 0x69,
	0x3c, 0xed, 0xe3, 0x37, 0xda, 0xef, 0x47, 0x25, 0x83, 0xd9, 0x74, 0x34, 0xfe, 0xd9, 0xac, 0xa0,
	0x67, 0x60, 0x1d, 0x49, 0x6e, 0xfb, 0x6e, 0x5f, 0x1e, 0xcd, 0xea, 0xf5, 0xff, 0x15, 0x68, 0xcc,
	0xb2, 0x1d, 0x86, 0x5e, 0x43, 0x5d, 0xee, 0x04, 0x94, 0x7d, 0xad, 0xe5, 0x35, 0xd9, 0x79, 0x7a,
	0x80, 0x65, 0x9f, 0x90, 0x88, 0x10, 0x4b, 0x22, 0x8f, 0x28, 0x6d, 0x90, 0x3c, 0xa2, 0xbc, 0x45,
	0xd0, 0x2b, 0xa8, 0x89, 0xb5, 0x81, 0xce, 0x35, 0x99, 0x6f, 0x94, 0x0e, 0x2a, 0x43, 0x99, 0xfc,
	0x07, 0x68, 0xe8, 0x15, 0x82, 0x3e, 0x57, 0xfc, 0xd1, 0x96, 0xe9, 0x5c, 0x1c, 0xc3, 0x59, 0xe8,
	0x8f, 0xd0, 0xd0, 0x83, 0xd6, 0xa1, 0x47, 0x9b, 0x47, 0x87, 0x1e, 0xaf, 0x84, 0xd7, 0x86, 0x0a,
	0x56, 0x73, 0x2a, 0x82, 0x0f, 0x1e, 0x78, 0x11, 0x7c, 0xf8, 0xac, 0x7a, 0xc6, 0xf2, 0x89, 0xfc,
	0xbf, 0xb9, 0xf9, 0x10, 0x00, 0x00, 0xff, 0xff, 0xff, 0xef, 0xaf, 0xd1, 0x81, 0x06, 0x00, 0x00,
}
