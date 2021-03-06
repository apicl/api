// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/service.proto

package istio_mixer_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mixer service

type MixerClient interface {
	// Checks preconditions and allocate quota before performing an operation.
	// The preconditions enforced depend on the set of supplied attributes and
	// the active configuration.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Reports telemetry, such as logs and metrics.
	// The reported information depends on the set of supplied attributes and the
	// active configuration.
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type mixerClient struct {
	cc *grpc.ClientConn
}

func NewMixerClient(cc *grpc.ClientConn) MixerClient {
	return &mixerClient{cc}
}

func (c *mixerClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/istio.mixer.v1.Mixer/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixerClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := grpc.Invoke(ctx, "/istio.mixer.v1.Mixer/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mixer service

type MixerServer interface {
	// Checks preconditions and allocate quota before performing an operation.
	// The preconditions enforced depend on the set of supplied attributes and
	// the active configuration.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Reports telemetry, such as logs and metrics.
	// The reported information depends on the set of supplied attributes and the
	// active configuration.
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterMixerServer(s *grpc.Server, srv MixerServer) {
	s.RegisterService(&_Mixer_serviceDesc, srv)
}

func _Mixer_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.mixer.v1.Mixer/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mixer_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixerServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.mixer.v1.Mixer/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixerServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mixer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.mixer.v1.Mixer",
	HandlerType: (*MixerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Mixer_Check_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _Mixer_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/v1/service.proto",
}

func init() { proto.RegisterFile("mixer/v1/service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x03, 0xcb, 0xea, 0x95, 0x19, 0x4a,
	0x89, 0xc0, 0xd5, 0x25, 0x67, 0xa4, 0x26, 0x67, 0x43, 0x54, 0x49, 0x89, 0xc2, 0x45, 0x8b, 0x52,
	0x0b, 0xf2, 0x8b, 0x4a, 0x20, 0xc2, 0x46, 0xb3, 0x18, 0xb9, 0x58, 0x7d, 0x41, 0x32, 0x42, 0x6e,
	0x5c, 0xac, 0xce, 0x20, 0xf5, 0x42, 0x32, 0x7a, 0xa8, 0x06, 0xea, 0x81, 0x85, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x64, 0x71, 0xc8, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x31,
	0x08, 0x79, 0x72, 0xb1, 0x05, 0x81, 0x6d, 0x10, 0xc2, 0x50, 0x0a, 0x11, 0x87, 0x99, 0x24, 0x87,
	0x4b, 0x1a, 0x66, 0x94, 0x93, 0xfe, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78,
	0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0x68, 0x60, 0x64, 0x4c, 0x62, 0x03, 0x7b, 0xca, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x1e, 0x4d, 0x13, 0x2b, 0x01, 0x00, 0x00,
}
