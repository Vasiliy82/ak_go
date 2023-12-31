// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package report

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReportClient is the client API for Report service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportClient interface {
	User(ctx context.Context, in *Dates, opts ...grpc.CallOption) (*Users, error)
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type reportClient struct {
	cc grpc.ClientConnInterface
}

func NewReportClient(cc grpc.ClientConnInterface) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) User(ctx context.Context, in *Dates, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/report.Report/user", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/report.Report/healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServer is the server API for Report service.
// All implementations must embed UnimplementedReportServer
// for forward compatibility
type ReportServer interface {
	User(context.Context, *Dates) (*Users, error)
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
	mustEmbedUnimplementedReportServer()
}

// UnimplementedReportServer must be embedded to have forward compatible implementations.
type UnimplementedReportServer struct {
}

func (UnimplementedReportServer) User(context.Context, *Dates) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedReportServer) Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedReportServer) mustEmbedUnimplementedReportServer() {}

// UnsafeReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServer will
// result in compilation errors.
type UnsafeReportServer interface {
	mustEmbedUnimplementedReportServer()
}

func RegisterReportServer(s grpc.ServiceRegistrar, srv ReportServer) {
	s.RegisterService(&Report_ServiceDesc, srv)
}

func _Report_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Report/user",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).User(ctx, req.(*Dates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Report/healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Report_ServiceDesc is the grpc.ServiceDesc for Report service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Report_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "report.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "user",
			Handler:    _Report_User_Handler,
		},
		{
			MethodName: "healthz",
			Handler:    _Report_Healthz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/report.proto",
}
