// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	TerminalStateRequest
	TerminalStateResponse
	TerminalBuyRequest
	TerminalBuyResponse
	TerminalScanRequest
	TerminalScanResponse
*/
package rpc

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

type TerminalStateRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *TerminalStateRequest) Reset()                    { *m = TerminalStateRequest{} }
func (m *TerminalStateRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateRequest) ProtoMessage()               {}
func (*TerminalStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TerminalStateRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type TerminalStateResponse struct {
}

func (m *TerminalStateResponse) Reset()                    { *m = TerminalStateResponse{} }
func (m *TerminalStateResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateResponse) ProtoMessage()               {}
func (*TerminalStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TerminalBuyRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *TerminalBuyRequest) Reset()                    { *m = TerminalBuyRequest{} }
func (m *TerminalBuyRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyRequest) ProtoMessage()               {}
func (*TerminalBuyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TerminalBuyRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type TerminalBuyResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalBuyResponse) Reset()                    { *m = TerminalBuyResponse{} }
func (m *TerminalBuyResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyResponse) ProtoMessage()               {}
func (*TerminalBuyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TerminalBuyResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TerminalScanRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	ProductID  string `protobuf:"bytes,2,opt,name=ProductID" json:"ProductID,omitempty"`
}

func (m *TerminalScanRequest) Reset()                    { *m = TerminalScanRequest{} }
func (m *TerminalScanRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanRequest) ProtoMessage()               {}
func (*TerminalScanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TerminalScanRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalScanRequest) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

type TerminalScanResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalScanResponse) Reset()                    { *m = TerminalScanResponse{} }
func (m *TerminalScanResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanResponse) ProtoMessage()               {}
func (*TerminalScanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TerminalScanResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*TerminalStateRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateRequest")
	proto.RegisterType((*TerminalStateResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse")
	proto.RegisterType((*TerminalBuyRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyRequest")
	proto.RegisterType((*TerminalBuyResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyResponse")
	proto.RegisterType((*TerminalScanRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanRequest")
	proto.RegisterType((*TerminalScanResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TerminalBackend service

type TerminalBackendClient interface {
	GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error)
	Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error)
	Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error)
}

type terminalBackendClient struct {
	cc *grpc.ClientConn
}

func NewTerminalBackendClient(cc *grpc.ClientConn) TerminalBackendClient {
	return &terminalBackendClient{cc}
}

func (c *terminalBackendClient) GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error) {
	out := new(TerminalStateResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error) {
	out := new(TerminalBuyResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error) {
	out := new(TerminalScanResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TerminalBackend service

type TerminalBackendServer interface {
	GetState(context.Context, *TerminalStateRequest) (*TerminalStateResponse, error)
	Buy(context.Context, *TerminalBuyRequest) (*TerminalBuyResponse, error)
	Scan(context.Context, *TerminalScanRequest) (*TerminalScanResponse, error)
}

func RegisterTerminalBackendServer(s *grpc.Server, srv TerminalBackendServer) {
	s.RegisterService(&_TerminalBackend_serviceDesc, srv)
}

func _TerminalBackend_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).GetState(ctx, req.(*TerminalStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalBuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Buy(ctx, req.(*TerminalBuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Scan(ctx, req.(*TerminalScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TerminalBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "i6getraenkeabrechnungssystem3000.rpc.TerminalBackend",
	HandlerType: (*TerminalBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _TerminalBackend_GetState_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _TerminalBackend_Buy_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _TerminalBackend_Scan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

func init() { proto.RegisterFile("main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0xcc, 0xcc,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc9, 0x34, 0x4b, 0x4f, 0x2d, 0x29, 0x4a, 0x4c,
	0xcd, 0xcb, 0x4e, 0x4d, 0x4c, 0x2a, 0x4a, 0x4d, 0xce, 0xc8, 0x2b, 0xcd, 0x4b, 0x2f, 0x2e, 0xae,
	0x2c, 0x2e, 0x49, 0xcd, 0x35, 0x36, 0x30, 0x30, 0xd0, 0x2b, 0x2a, 0x48, 0x56, 0x32, 0xe3, 0x12,
	0x09, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x4b, 0xcc, 0x09, 0x2e, 0x49, 0x2c, 0x49, 0x0d, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0x82, 0x89, 0x7b, 0xba, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x21, 0x89, 0x28, 0x89, 0x73, 0x89, 0xa2, 0xe9, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x55, 0x32, 0xe1, 0x12, 0x82, 0x49, 0x38, 0x95, 0x56, 0x12, 0x6b, 0x9c, 0x36, 0x97, 0x30,
	0x8a, 0x2e, 0x88, 0x61, 0x42, 0x22, 0x5c, 0xac, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x50, 0x1d, 0x10,
	0x8e, 0x52, 0x30, 0x42, 0x71, 0x70, 0x72, 0x62, 0x1e, 0x91, 0x76, 0x08, 0xc9, 0x70, 0x71, 0x06,
	0x14, 0xe5, 0xa7, 0x94, 0x26, 0x97, 0x78, 0xba, 0x48, 0x30, 0x81, 0xa5, 0x11, 0x02, 0x4a, 0x3a,
	0x48, 0x01, 0x01, 0x36, 0x14, 0x9f, 0x13, 0x8c, 0x66, 0x30, 0x73, 0xf1, 0xc3, 0x1d, 0x9c, 0x98,
	0x9c, 0x9d, 0x9a, 0x97, 0x22, 0xd4, 0xca, 0xc8, 0xc5, 0xe1, 0x9e, 0x5a, 0x02, 0x0e, 0x0e, 0x21,
	0x2b, 0x3d, 0x62, 0x82, 0x5f, 0x0f, 0x5b, 0xd8, 0x4b, 0x59, 0x93, 0xa5, 0x17, 0x1a, 0xfe, 0x0c,
	0x42, 0x35, 0x5c, 0xcc, 0x4e, 0xa5, 0x95, 0x42, 0x16, 0xa4, 0x99, 0x82, 0x88, 0x2c, 0x29, 0x4b,
	0x32, 0x74, 0xc2, 0x6d, 0xaf, 0xe7, 0x62, 0x01, 0x85, 0x9f, 0x10, 0x89, 0x86, 0x20, 0x45, 0xa4,
	0x94, 0x15, 0x39, 0x5a, 0x61, 0x0e, 0x70, 0x62, 0x8d, 0x62, 0x2e, 0x2a, 0x48, 0x4e, 0x62, 0x03,
	0xe7, 0x02, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xf6, 0xf2, 0x2e, 0x13, 0x03, 0x00,
	0x00,
}
