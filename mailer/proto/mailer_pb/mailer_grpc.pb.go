// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: mailer.proto

package proto

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

// MailerServiceClient is the client API for MailerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerServiceClient interface {
	SubscribeUser(ctx context.Context, in *SubscribeForm, opts ...grpc.CallOption) (*Response, error)
}

type mailerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerServiceClient(cc grpc.ClientConnInterface) MailerServiceClient {
	return &mailerServiceClient{cc}
}

func (c *mailerServiceClient) SubscribeUser(ctx context.Context, in *SubscribeForm, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MailerService/SubscribeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServiceServer is the server API for MailerService service.
// All implementations must embed UnimplementedMailerServiceServer
// for forward compatibility
type MailerServiceServer interface {
	SubscribeUser(context.Context, *SubscribeForm) (*Response, error)
	mustEmbedUnimplementedMailerServiceServer()
}

// UnimplementedMailerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailerServiceServer struct {
}

func (UnimplementedMailerServiceServer) SubscribeUser(context.Context, *SubscribeForm) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeUser not implemented")
}
func (UnimplementedMailerServiceServer) mustEmbedUnimplementedMailerServiceServer() {}

// UnsafeMailerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServiceServer will
// result in compilation errors.
type UnsafeMailerServiceServer interface {
	mustEmbedUnimplementedMailerServiceServer()
}

func RegisterMailerServiceServer(s grpc.ServiceRegistrar, srv MailerServiceServer) {
	s.RegisterService(&MailerService_ServiceDesc, srv)
}

func _MailerService_SubscribeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServiceServer).SubscribeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MailerService/SubscribeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServiceServer).SubscribeUser(ctx, req.(*SubscribeForm))
	}
	return interceptor(ctx, in, info, handler)
}

// MailerService_ServiceDesc is the grpc.ServiceDesc for MailerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MailerService",
	HandlerType: (*MailerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeUser",
			Handler:    _MailerService_SubscribeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer.proto",
}
