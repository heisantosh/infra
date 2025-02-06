// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: template-manager.proto

package template_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	// TemplateCreate is a gRPC service that creates a new template
	TemplateCreate(ctx context.Context, in *TemplateCreateRequest, opts ...grpc.CallOption) (TemplateService_TemplateCreateClient, error)
	// EnvBuildDelete is a gRPC service that deletes files associated with a template build
	TemplateBuildDelete(ctx context.Context, in *TemplateBuildDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) TemplateCreate(ctx context.Context, in *TemplateCreateRequest, opts ...grpc.CallOption) (TemplateService_TemplateCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &TemplateService_ServiceDesc.Streams[0], "/TemplateService/TemplateCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateServiceTemplateCreateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TemplateService_TemplateCreateClient interface {
	Recv() (*TemplateBuildLog, error)
	grpc.ClientStream
}

type templateServiceTemplateCreateClient struct {
	grpc.ClientStream
}

func (x *templateServiceTemplateCreateClient) Recv() (*TemplateBuildLog, error) {
	m := new(TemplateBuildLog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *templateServiceClient) TemplateBuildDelete(ctx context.Context, in *TemplateBuildDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/TemplateService/TemplateBuildDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations must embed UnimplementedTemplateServiceServer
// for forward compatibility
type TemplateServiceServer interface {
	// TemplateCreate is a gRPC service that creates a new template
	TemplateCreate(*TemplateCreateRequest, TemplateService_TemplateCreateServer) error
	// EnvBuildDelete is a gRPC service that deletes files associated with a template build
	TemplateBuildDelete(context.Context, *TemplateBuildDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTemplateServiceServer()
}

// UnimplementedTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct {
}

func (UnimplementedTemplateServiceServer) TemplateCreate(*TemplateCreateRequest, TemplateService_TemplateCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method TemplateCreate not implemented")
}
func (UnimplementedTemplateServiceServer) TemplateBuildDelete(context.Context, *TemplateBuildDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateBuildDelete not implemented")
}
func (UnimplementedTemplateServiceServer) mustEmbedUnimplementedTemplateServiceServer() {}

// UnsafeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServiceServer will
// result in compilation errors.
type UnsafeTemplateServiceServer interface {
	mustEmbedUnimplementedTemplateServiceServer()
}

func RegisterTemplateServiceServer(s grpc.ServiceRegistrar, srv TemplateServiceServer) {
	s.RegisterService(&TemplateService_ServiceDesc, srv)
}

func _TemplateService_TemplateCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TemplateCreateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemplateServiceServer).TemplateCreate(m, &templateServiceTemplateCreateServer{stream})
}

type TemplateService_TemplateCreateServer interface {
	Send(*TemplateBuildLog) error
	grpc.ServerStream
}

type templateServiceTemplateCreateServer struct {
	grpc.ServerStream
}

func (x *templateServiceTemplateCreateServer) Send(m *TemplateBuildLog) error {
	return x.ServerStream.SendMsg(m)
}

func _TemplateService_TemplateBuildDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateBuildDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).TemplateBuildDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TemplateService/TemplateBuildDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).TemplateBuildDelete(ctx, req.(*TemplateBuildDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateService_ServiceDesc is the grpc.ServiceDesc for TemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TemplateBuildDelete",
			Handler:    _TemplateService_TemplateBuildDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TemplateCreate",
			Handler:       _TemplateService_TemplateCreate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "template-manager.proto",
}
