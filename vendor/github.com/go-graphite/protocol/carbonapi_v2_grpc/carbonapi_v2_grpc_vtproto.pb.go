// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.3.0
// source: carbonapi_v2_grpc/carbonapi_v2_grpc.proto

package carbonapi_v2_grpc

import (
	context "context"
	carbonapi_v2_pb "github.com/go-graphite/protocol/carbonapi_v2_pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CarbonV2Client is the client API for CarbonV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarbonV2Client interface {
	Render(ctx context.Context, in *carbonapi_v2_pb.MultiFetchRequest, opts ...grpc.CallOption) (CarbonV2_RenderClient, error)
	Find(ctx context.Context, in *carbonapi_v2_pb.GlobRequest, opts ...grpc.CallOption) (*carbonapi_v2_pb.GlobResponse, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.ListMetricsResponse, error)
	Info(ctx context.Context, in *carbonapi_v2_pb.InfoRequest, opts ...grpc.CallOption) (*carbonapi_v2_pb.InfoResponse, error)
	Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.MetricDetailsResponse, error)
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.ProtocolVersionResponse, error)
}

type carbonV2Client struct {
	cc grpc.ClientConnInterface
}

func NewCarbonV2Client(cc grpc.ClientConnInterface) CarbonV2Client {
	return &carbonV2Client{cc}
}

func (c *carbonV2Client) Render(ctx context.Context, in *carbonapi_v2_pb.MultiFetchRequest, opts ...grpc.CallOption) (CarbonV2_RenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarbonV2_ServiceDesc.Streams[0], "/carbonapi_v2_grpc.CarbonV2/Render", opts...)
	if err != nil {
		return nil, err
	}
	x := &carbonV2RenderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarbonV2_RenderClient interface {
	Recv() (*carbonapi_v2_pb.FetchResponse, error)
	grpc.ClientStream
}

type carbonV2RenderClient struct {
	grpc.ClientStream
}

func (x *carbonV2RenderClient) Recv() (*carbonapi_v2_pb.FetchResponse, error) {
	m := new(carbonapi_v2_pb.FetchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *carbonV2Client) Find(ctx context.Context, in *carbonapi_v2_pb.GlobRequest, opts ...grpc.CallOption) (*carbonapi_v2_pb.GlobResponse, error) {
	out := new(carbonapi_v2_pb.GlobResponse)
	err := c.cc.Invoke(ctx, "/carbonapi_v2_grpc.CarbonV2/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV2Client) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.ListMetricsResponse, error) {
	out := new(carbonapi_v2_pb.ListMetricsResponse)
	err := c.cc.Invoke(ctx, "/carbonapi_v2_grpc.CarbonV2/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV2Client) Info(ctx context.Context, in *carbonapi_v2_pb.InfoRequest, opts ...grpc.CallOption) (*carbonapi_v2_pb.InfoResponse, error) {
	out := new(carbonapi_v2_pb.InfoResponse)
	err := c.cc.Invoke(ctx, "/carbonapi_v2_grpc.CarbonV2/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV2Client) Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.MetricDetailsResponse, error) {
	out := new(carbonapi_v2_pb.MetricDetailsResponse)
	err := c.cc.Invoke(ctx, "/carbonapi_v2_grpc.CarbonV2/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carbonV2Client) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*carbonapi_v2_pb.ProtocolVersionResponse, error) {
	out := new(carbonapi_v2_pb.ProtocolVersionResponse)
	err := c.cc.Invoke(ctx, "/carbonapi_v2_grpc.CarbonV2/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarbonV2Server is the server API for CarbonV2 service.
// All implementations must embed UnimplementedCarbonV2Server
// for forward compatibility
type CarbonV2Server interface {
	Render(*carbonapi_v2_pb.MultiFetchRequest, CarbonV2_RenderServer) error
	Find(context.Context, *carbonapi_v2_pb.GlobRequest) (*carbonapi_v2_pb.GlobResponse, error)
	List(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.ListMetricsResponse, error)
	Info(context.Context, *carbonapi_v2_pb.InfoRequest) (*carbonapi_v2_pb.InfoResponse, error)
	Stats(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.MetricDetailsResponse, error)
	Version(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.ProtocolVersionResponse, error)
	mustEmbedUnimplementedCarbonV2Server()
}

// UnimplementedCarbonV2Server must be embedded to have forward compatible implementations.
type UnimplementedCarbonV2Server struct {
}

func (UnimplementedCarbonV2Server) Render(*carbonapi_v2_pb.MultiFetchRequest, CarbonV2_RenderServer) error {
	return status.Errorf(codes.Unimplemented, "method Render not implemented")
}
func (UnimplementedCarbonV2Server) Find(context.Context, *carbonapi_v2_pb.GlobRequest) (*carbonapi_v2_pb.GlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedCarbonV2Server) List(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.ListMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCarbonV2Server) Info(context.Context, *carbonapi_v2_pb.InfoRequest) (*carbonapi_v2_pb.InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedCarbonV2Server) Stats(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.MetricDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedCarbonV2Server) Version(context.Context, *emptypb.Empty) (*carbonapi_v2_pb.ProtocolVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCarbonV2Server) mustEmbedUnimplementedCarbonV2Server() {}

// UnsafeCarbonV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarbonV2Server will
// result in compilation errors.
type UnsafeCarbonV2Server interface {
	mustEmbedUnimplementedCarbonV2Server()
}

func RegisterCarbonV2Server(s grpc.ServiceRegistrar, srv CarbonV2Server) {
	s.RegisterService(&CarbonV2_ServiceDesc, srv)
}

func _CarbonV2_Render_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(carbonapi_v2_pb.MultiFetchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarbonV2Server).Render(m, &carbonV2RenderServer{stream})
}

type CarbonV2_RenderServer interface {
	Send(*carbonapi_v2_pb.FetchResponse) error
	grpc.ServerStream
}

type carbonV2RenderServer struct {
	grpc.ServerStream
}

func (x *carbonV2RenderServer) Send(m *carbonapi_v2_pb.FetchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CarbonV2_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(carbonapi_v2_pb.GlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV2Server).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carbonapi_v2_grpc.CarbonV2/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV2Server).Find(ctx, req.(*carbonapi_v2_pb.GlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV2_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV2Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carbonapi_v2_grpc.CarbonV2/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV2Server).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV2_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(carbonapi_v2_pb.InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV2Server).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carbonapi_v2_grpc.CarbonV2/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV2Server).Info(ctx, req.(*carbonapi_v2_pb.InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV2_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV2Server).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carbonapi_v2_grpc.CarbonV2/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV2Server).Stats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarbonV2_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarbonV2Server).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carbonapi_v2_grpc.CarbonV2/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarbonV2Server).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CarbonV2_ServiceDesc is the grpc.ServiceDesc for CarbonV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarbonV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "carbonapi_v2_grpc.CarbonV2",
	HandlerType: (*CarbonV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _CarbonV2_Find_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CarbonV2_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _CarbonV2_Info_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _CarbonV2_Stats_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _CarbonV2_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Render",
			Handler:       _CarbonV2_Render_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "carbonapi_v2_grpc/carbonapi_v2_grpc.proto",
}