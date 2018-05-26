// Code generated by protoc-gen-go. DO NOT EDIT.
// source: poll.proto

package poll

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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
	Title                string               `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Id                   int64                `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Poll) Reset()         { *m = Poll{} }
func (m *Poll) String() string { return proto.CompactTextString(m) }
func (*Poll) ProtoMessage()    {}
func (*Poll) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_6549b5b410e4958b, []int{0}
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

type PollRequest struct {
	Id                   int64    `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollRequest) Reset()         { *m = PollRequest{} }
func (m *PollRequest) String() string { return proto.CompactTextString(m) }
func (*PollRequest) ProtoMessage()    {}
func (*PollRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_6549b5b410e4958b, []int{1}
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
	Tick                 string   `protobuf:"bytes,6,opt,name=tick" json:"tick,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickerResponse) Reset()         { *m = TickerResponse{} }
func (m *TickerResponse) String() string { return proto.CompactTextString(m) }
func (*TickerResponse) ProtoMessage()    {}
func (*TickerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poll_6549b5b410e4958b, []int{2}
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

func init() { proto.RegisterFile("poll.proto", fileDescriptor_poll_6549b5b410e4958b) }

var fileDescriptor_poll_6549b5b410e4958b = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0x8f, 0x37, 0xf0, 0x4e, 0xb5, 0xe0, 0x50, 0x24, 0x44, 0xc4, 0x10, 0x44, 0x72,
	0x4a, 0xa5, 0xde, 0xbc, 0x78, 0x12, 0x4f, 0x82, 0xac, 0xbd, 0x4b, 0xdb, 0x4c, 0xcb, 0xd2, 0x5d,
	0x77, 0xcd, 0x4e, 0x05, 0xff, 0x7b, 0xc9, 0x6e, 0x83, 0x5f, 0xb7, 0xc9, 0xe4, 0xb7, 0xbf, 0xe7,
	0x61, 0x00, 0xac, 0x51, 0xaa, 0xb1, 0x9d, 0x61, 0x83, 0x69, 0x3f, 0x17, 0x17, 0x5b, 0x63, 0xb6,
	0x8a, 0x66, 0x7e, 0xb7, 0xda, 0x6f, 0x66, 0x2c, 0x35, 0x39, 0x5e, 0x6a, 0x1b, 0xb0, 0xe2, 0xec,
	0x37, 0x70, 0xaf, 0x2d, 0x7f, 0x84, 0x9f, 0x95, 0x86, 0xf4, 0xc9, 0x28, 0x85, 0x53, 0xf8, 0xc7,
	0x92, 0x15, 0xe5, 0x51, 0x19, 0xd5, 0xff, 0x45, 0xf8, 0xc0, 0x09, 0xc4, 0xb2, 0xcd, 0xe3, 0x32,
	0xaa, 0x13, 0x11, 0xcb, 0x16, 0xef, 0xe0, 0x58, 0x2d, 0x1d, 0xbf, 0x68, 0xd3, 0xca, 0x8d, 0xa4,
	0x36, 0x4f, 0xca, 0xa8, 0x1e, 0xcf, 0x8b, 0x26, 0x44, 0x34, 0x43, 0x44, 0xb3, 0x18, 0x3a, 0x88,
	0xa3, 0xfe, 0xc1, 0xe3, 0x81, 0xaf, 0xce, 0x61, 0xdc, 0xc7, 0x09, 0x7a, 0xdb, 0x93, 0xe3, 0x83,
	0x3f, 0x1d, 0xfc, 0xd5, 0x25, 0x4c, 0x16, 0x72, 0xbd, 0xa3, 0x4e, 0x90, 0xb3, 0xe6, 0xd5, 0x11,
	0x22, 0xa4, 0x2c, 0xd7, 0xbb, 0x3c, 0xf3, 0xb5, 0xfc, 0x3c, 0xb7, 0x00, 0xbd, 0xe4, 0x99, 0xba,
	0x77, 0xea, 0xf0, 0x0a, 0x92, 0x07, 0x62, 0x3c, 0x69, 0xfc, 0x65, 0xbe, 0xd9, 0x0b, 0xf8, 0x5a,
	0x55, 0x23, 0xbc, 0x85, 0x2c, 0xb8, 0xf1, 0xf4, 0x4f, 0x5d, 0x7f, 0x91, 0x62, 0x1a, 0xf8, 0x9f,
	0x0d, 0xaa, 0xd1, 0x75, 0xb4, 0xca, 0x3c, 0x79, 0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0x57, 0x5c,
	0x67, 0xa4, 0x7e, 0x01, 0x00, 0x00,
}