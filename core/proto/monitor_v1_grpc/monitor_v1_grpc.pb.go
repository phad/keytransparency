// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor_v1_grpc.proto

/*
Package monitor_v1_grpc is a generated protocol buffer package.

Monitor Service

The Key Transparency monitor server service consists of APIs to fetch
monitor results queried using the mutations API.

It is generated from these files:
	monitor_v1_grpc.proto

It has these top-level messages:
*/
package monitor_v1_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import monitor_v1_proto "github.com/google/keytransparency/core/proto/monitor_v1_proto"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MonitorService service

type MonitorServiceClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(ctx context.Context, in *monitor_v1_proto.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_proto.GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(ctx context.Context, in *monitor_v1_proto.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_proto.GetMonitoringResponse, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) GetSignedMapRoot(ctx context.Context, in *monitor_v1_proto.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_proto.GetMonitoringResponse, error) {
	out := new(monitor_v1_proto.GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.grpc.MonitorService/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetSignedMapRootByRevision(ctx context.Context, in *monitor_v1_proto.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_proto.GetMonitoringResponse, error) {
	out := new(monitor_v1_proto.GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.grpc.MonitorService/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(context.Context, *monitor_v1_proto.GetMonitoringRequest) (*monitor_v1_proto.GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(context.Context, *monitor_v1_proto.GetMonitoringRequest) (*monitor_v1_proto.GetMonitoringResponse, error)
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(monitor_v1_proto.GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.grpc.MonitorService/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, req.(*monitor_v1_proto.GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(monitor_v1_proto.GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.grpc.MonitorService/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, req.(*monitor_v1_proto.GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.v1.grpc.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _MonitorService_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _MonitorService_GetSignedMapRootByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor_v1_grpc.proto",
}

func init() { proto.RegisterFile("monitor_v1_grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xd5, 0x0e, 0x0c, 0x19, 0x00, 0x59, 0x62, 0x89, 0x98, 0x18, 0x28, 0x30, 0xe4, 0x08,
	0x6c, 0x8c, 0x2c, 0x5d, 0xe8, 0x92, 0xc2, 0x1c, 0xb9, 0xe6, 0xe4, 0x5a, 0x4d, 0x7c, 0xc6, 0xbe,
	0x58, 0x8a, 0xaa, 0x2e, 0xbc, 0x02, 0xec, 0x3c, 0x0f, 0x33, 0xaf, 0xc0, 0x83, 0x20, 0x92, 0x20,
	0x55, 0x08, 0xa9, 0x0b, 0xeb, 0x6f, 0x7f, 0xdf, 0x7d, 0xc9, 0x51, 0x4d, 0xd6, 0x30, 0xf9, 0x32,
	0xe6, 0xa5, 0xf6, 0x4e, 0x65, 0xce, 0x13, 0x93, 0x38, 0x18, 0xe6, 0x2c, 0xe6, 0xd9, 0xf7, 0x9c,
	0xde, 0x6b, 0xc3, 0xcb, 0x66, 0x91, 0x29, 0xaa, 0x41, 0x13, 0xe9, 0x0a, 0x61, 0x85, 0x2d, 0x7b,
	0x69, 0x83, 0x93, 0x1e, 0xad, 0x6a, 0x41, 0x91, 0x47, 0xe8, 0x70, 0xd8, 0x92, 0xfe, 0x3d, 0xf4,
	0x67, 0xd2, 0xe3, 0x41, 0x25, 0x9d, 0x01, 0x69, 0x2d, 0xb1, 0x64, 0x43, 0x36, 0xf4, 0xaf, 0x57,
	0xef, 0xe3, 0x64, 0x7f, 0xd6, 0x83, 0x73, 0xf4, 0xd1, 0x28, 0x14, 0xaf, 0xa3, 0xe4, 0x70, 0x8a,
	0x3c, 0x37, 0xda, 0xe2, 0xe3, 0x4c, 0xba, 0x82, 0x88, 0xc5, 0x69, 0xb6, 0x55, 0xdb, 0xeb, 0xa7,
	0xc8, 0x03, 0x69, 0xac, 0x2e, 0xf0, 0xa9, 0xc1, 0xc0, 0xe9, 0x64, 0xe7, 0xbf, 0xe0, 0xc8, 0x06,
	0x3c, 0x81, 0xe7, 0x8f, 0xcf, 0x97, 0xf1, 0xb9, 0x98, 0x40, 0xcc, 0x7f, 0xd2, 0x61, 0xbd, 0xe2,
	0xf2, 0xa1, 0xb8, 0xdb, 0x40, 0x2d, 0x1d, 0x78, 0x0c, 0x4d, 0xc5, 0xe1, 0xa6, 0x92, 0x8c, 0x81,
	0xc5, 0xdb, 0x28, 0x49, 0x7f, 0x67, 0xdd, 0xb6, 0x05, 0x46, 0x13, 0x0c, 0xd9, 0xff, 0x0f, 0xbc,
	0xec, 0x02, 0x2f, 0xc4, 0xd9, 0xae, 0x40, 0x58, 0xa3, 0x23, 0xb5, 0xdc, 0x2c, 0xf6, 0x3a, 0xdd,
	0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x28, 0xda, 0xff, 0xf0, 0x01, 0x00, 0x00,
}