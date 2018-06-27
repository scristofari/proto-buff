// Code generated by protoc-gen-go. DO NOT EDIT.
// source: poll.proto

package poll

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Poll struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Id                   int64                `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	Published            *wrappers.BoolValue  `protobuf:"bytes,4,opt,name=published,proto3" json:"published,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Poll) Reset()         { *m = Poll{} }
func (m *Poll) String() string { return proto.CompactTextString(m) }
func (*Poll) ProtoMessage()    {}
func (*Poll) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_a5e53eb1030d5dae, []int{0}
}
func (m *Poll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Poll.Unmarshal(m, b)
}
func (m *Poll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Poll.Marshal(b, m, deterministic)
}
func (dst *Poll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Poll.Merge(dst, src)
}
func (m *Poll) XXX_Size() int {
	return xxx_messageInfo_Poll.Size(m)
}
func (m *Poll) XXX_DiscardUnknown() {
	xxx_messageInfo_Poll.DiscardUnknown(m)
}

var xxx_messageInfo_Poll proto.InternalMessageInfo

func (m *Poll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Poll) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Poll) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *Poll) GetPublished() *wrappers.BoolValue {
	if m != nil {
		return m.Published
	}
	return nil
}

type PollRequest struct {
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollRequest) Reset()         { *m = PollRequest{} }
func (m *PollRequest) String() string { return proto.CompactTextString(m) }
func (*PollRequest) ProtoMessage()    {}
func (*PollRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_a5e53eb1030d5dae, []int{1}
}
func (m *PollRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollRequest.Unmarshal(m, b)
}
func (m *PollRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollRequest.Marshal(b, m, deterministic)
}
func (dst *PollRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollRequest.Merge(dst, src)
}
func (m *PollRequest) XXX_Size() int {
	return xxx_messageInfo_PollRequest.Size(m)
}
func (m *PollRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PollRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PollRequest proto.InternalMessageInfo

func (m *PollRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TickerResponse struct {
	Tick                 string   `protobuf:"bytes,6,opt,name=tick,proto3" json:"tick,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickerResponse) Reset()         { *m = TickerResponse{} }
func (m *TickerResponse) String() string { return proto.CompactTextString(m) }
func (*TickerResponse) ProtoMessage()    {}
func (*TickerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_a5e53eb1030d5dae, []int{2}
}
func (m *TickerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TickerResponse.Unmarshal(m, b)
}
func (m *TickerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TickerResponse.Marshal(b, m, deterministic)
}
func (dst *TickerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickerResponse.Merge(dst, src)
}
func (m *TickerResponse) XXX_Size() int {
	return xxx_messageInfo_TickerResponse.Size(m)
}
func (m *TickerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TickerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TickerResponse proto.InternalMessageInfo

func (m *TickerResponse) GetTick() string {
	if m != nil {
		return m.Tick
	}
	return ""
}

func init() {
	proto.RegisterType((*Poll)(nil), "poll.Poll")
	proto.RegisterType((*PollRequest)(nil), "poll.PollRequest")
	proto.RegisterType((*TickerResponse)(nil), "poll.TickerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PollServerClient is the client API for PollServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PollServerClient interface {
	Get(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*Poll, error)
	Ticker(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (PollServer_TickerClient, error)
}

type pollServerClient struct {
	cc *grpc.ClientConn
}

func NewPollServerClient(cc *grpc.ClientConn) PollServerClient {
	return &pollServerClient{cc}
}

func (c *pollServerClient) Get(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*Poll, error) {
	out := new(Poll)
	err := c.cc.Invoke(ctx, "/poll.PollServer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollServerClient) Ticker(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (PollServer_TickerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PollServer_serviceDesc.Streams[0], "/poll.PollServer/Ticker", opts...)
	if err != nil {
		return nil, err
	}
	x := &pollServerTickerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PollServer_TickerClient interface {
	Recv() (*TickerResponse, error)
	grpc.ClientStream
}

type pollServerTickerClient struct {
	grpc.ClientStream
}

func (x *pollServerTickerClient) Recv() (*TickerResponse, error) {
	m := new(TickerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PollServerServer is the server API for PollServer service.
type PollServerServer interface {
	Get(context.Context, *PollRequest) (*Poll, error)
	Ticker(*empty.Empty, PollServer_TickerServer) error
}

func RegisterPollServerServer(s *grpc.Server, srv PollServerServer) {
	s.RegisterService(&_PollServer_serviceDesc, srv)
}

func _PollServer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poll.PollServer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServerServer).Get(ctx, req.(*PollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollServer_Ticker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PollServerServer).Ticker(m, &pollServerTickerServer{stream})
}

type PollServer_TickerServer interface {
	Send(*TickerResponse) error
	grpc.ServerStream
}

type pollServerTickerServer struct {
	grpc.ServerStream
}

func (x *pollServerTickerServer) Send(m *TickerResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PollServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poll.PollServer",
	HandlerType: (*PollServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PollServer_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ticker",
			Handler:       _PollServer_Ticker_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "poll.proto",
}

func init() { proto.RegisterFile("poll.proto", fileDescriptor_poll_a5e53eb1030d5dae) }

var fileDescriptor_poll_a5e53eb1030d5dae = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0x69, 0x6e, 0xa1, 0xa7, 0x3f, 0x97, 0x3b, 0xf4, 0x5e, 0x42, 0xae, 0x3f, 0x25, 0xb8,
	0xe8, 0x2a, 0xd1, 0xba, 0x11, 0x5c, 0x08, 0x82, 0xb8, 0x51, 0x90, 0x58, 0xdc, 0x4a, 0xda, 0x9c,
	0xd6, 0xa1, 0x93, 0xcc, 0x98, 0x9c, 0x54, 0xdc, 0xfa, 0x0a, 0x2e, 0x7c, 0x04, 0x1f, 0xc8, 0x57,
	0xf0, 0x41, 0x64, 0x26, 0x29, 0xd5, 0x8a, 0xbb, 0x99, 0xf3, 0xfd, 0x9c, 0xef, 0xe3, 0x00, 0x28,
	0x29, 0x44, 0xa0, 0x72, 0x49, 0x92, 0x39, 0xfa, 0xed, 0xed, 0xce, 0xa5, 0x9c, 0x0b, 0x0c, 0xcd,
	0x6c, 0x52, 0xce, 0x42, 0xe2, 0x29, 0x16, 0x14, 0xa7, 0xaa, 0xa2, 0x79, 0x5b, 0x35, 0x21, 0x56,
	0x3c, 0x8c, 0xb3, 0x4c, 0x52, 0x4c, 0x5c, 0x66, 0x45, 0x8d, 0xfe, 0xdf, 0x94, 0x9f, 0xa5, 0x8a,
	0x1e, 0x6b, 0x70, 0x67, 0x13, 0x7c, 0xc8, 0x63, 0xa5, 0x30, 0xaf, 0xc5, 0xfe, 0xab, 0x05, 0xce,
	0x95, 0x14, 0x82, 0xf5, 0xe1, 0x17, 0x71, 0x12, 0xe8, 0x5a, 0x03, 0x6b, 0xd8, 0x8a, 0xaa, 0x0f,
	0xeb, 0x81, 0xcd, 0x13, 0xd7, 0x1e, 0x58, 0xc3, 0x46, 0x64, 0xf3, 0x84, 0x9d, 0x40, 0x57, 0xc4,
	0x05, 0xdd, 0xa6, 0x32, 0xe1, 0x33, 0x8e, 0x89, 0xdb, 0x18, 0x58, 0xc3, 0xf6, 0xc8, 0x0b, 0xaa,
	0x35, 0xc1, 0x6a, 0x4d, 0x30, 0x5e, 0x55, 0x88, 0x3a, 0x5a, 0x70, 0x59, 0xf3, 0xd9, 0x11, 0xb4,
	0x54, 0x39, 0x11, 0xbc, 0xb8, 0xc3, 0xc4, 0x75, 0x7e, 0x10, 0x9f, 0x4a, 0x29, 0x6e, 0x62, 0x51,
	0x62, 0xb4, 0x26, 0xfb, 0xdb, 0xd0, 0xd6, 0x41, 0x23, 0xbc, 0x2f, 0xb1, 0xa0, 0x3a, 0x99, 0xb3,
	0x4a, 0xe6, 0xef, 0x41, 0x6f, 0xcc, 0xa7, 0x0b, 0xcc, 0x23, 0x2c, 0x94, 0xcc, 0x0a, 0x64, 0x0c,
	0x1c, 0xe2, 0xd3, 0x85, 0xdb, 0x34, 0x85, 0xcc, 0x7b, 0xf4, 0x62, 0x01, 0x68, 0x97, 0x6b, 0xcc,
	0x97, 0x98, 0xb3, 0x63, 0x68, 0x9c, 0x23, 0xb1, 0x3f, 0x81, 0xb9, 0xc9, 0x27, 0x7b, 0x0f, 0xd6,
	0x23, 0xff, 0xef, 0xd3, 0xdb, 0xfb, 0xb3, 0xfd, 0x9b, 0x75, 0xc3, 0xe5, 0x41, 0xa8, 0xc7, 0x21,
	0xe9, 0x04, 0x17, 0xd0, 0xac, 0x36, 0xb2, 0x7f, 0xdf, 0x1a, 0x98, 0x13, 0x78, 0xfd, 0xca, 0xe4,
	0x6b, 0x2e, 0x9f, 0x19, 0xbb, 0x0e, 0x03, 0x6d, 0x47, 0x06, 0xdb, 0xb7, 0x26, 0x4d, 0xa3, 0x3d,
	0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x3d, 0xf6, 0x56, 0x1f, 0x02, 0x00, 0x00,
}
