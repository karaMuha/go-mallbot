// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: notificationspb/api.proto

package notificationspb

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

// NotificationsServiceClient is the client API for NotificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsServiceClient interface {
	NotifyOrderCreated(ctx context.Context, in *NotifyOrderCreatedRequest, opts ...grpc.CallOption) (*NotifyOrderCreatedResponse, error)
	NotifyOrderCanceled(ctx context.Context, in *NotifyOrderCanceledRequest, opts ...grpc.CallOption) (*NotifyOrderCanceledResponse, error)
	NotifyOrderReady(ctx context.Context, in *NotifyOrderReadyRequest, opts ...grpc.CallOption) (*NotifyOrderReadyResponse, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) NotifyOrderCreated(ctx context.Context, in *NotifyOrderCreatedRequest, opts ...grpc.CallOption) (*NotifyOrderCreatedResponse, error) {
	out := new(NotifyOrderCreatedResponse)
	err := c.cc.Invoke(ctx, "/notificationspb.NotificationsService/NotifyOrderCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) NotifyOrderCanceled(ctx context.Context, in *NotifyOrderCanceledRequest, opts ...grpc.CallOption) (*NotifyOrderCanceledResponse, error) {
	out := new(NotifyOrderCanceledResponse)
	err := c.cc.Invoke(ctx, "/notificationspb.NotificationsService/NotifyOrderCanceled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) NotifyOrderReady(ctx context.Context, in *NotifyOrderReadyRequest, opts ...grpc.CallOption) (*NotifyOrderReadyResponse, error) {
	out := new(NotifyOrderReadyResponse)
	err := c.cc.Invoke(ctx, "/notificationspb.NotificationsService/NotifyOrderReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations must embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	NotifyOrderCreated(context.Context, *NotifyOrderCreatedRequest) (*NotifyOrderCreatedResponse, error)
	NotifyOrderCanceled(context.Context, *NotifyOrderCanceledRequest) (*NotifyOrderCanceledResponse, error)
	NotifyOrderReady(context.Context, *NotifyOrderReadyRequest) (*NotifyOrderReadyResponse, error)
	mustEmbedUnimplementedNotificationsServiceServer()
}

// UnimplementedNotificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct {
}

func (UnimplementedNotificationsServiceServer) NotifyOrderCreated(context.Context, *NotifyOrderCreatedRequest) (*NotifyOrderCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyOrderCreated not implemented")
}
func (UnimplementedNotificationsServiceServer) NotifyOrderCanceled(context.Context, *NotifyOrderCanceledRequest) (*NotifyOrderCanceledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyOrderCanceled not implemented")
}
func (UnimplementedNotificationsServiceServer) NotifyOrderReady(context.Context, *NotifyOrderReadyRequest) (*NotifyOrderReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyOrderReady not implemented")
}
func (UnimplementedNotificationsServiceServer) mustEmbedUnimplementedNotificationsServiceServer() {}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_NotifyOrderCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyOrderCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifyOrderCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationspb.NotificationsService/NotifyOrderCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifyOrderCreated(ctx, req.(*NotifyOrderCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_NotifyOrderCanceled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyOrderCanceledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifyOrderCanceled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationspb.NotificationsService/NotifyOrderCanceled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifyOrderCanceled(ctx, req.(*NotifyOrderCanceledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_NotifyOrderReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyOrderReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifyOrderReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationspb.NotificationsService/NotifyOrderReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifyOrderReady(ctx, req.(*NotifyOrderReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notificationspb.NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyOrderCreated",
			Handler:    _NotificationsService_NotifyOrderCreated_Handler,
		},
		{
			MethodName: "NotifyOrderCanceled",
			Handler:    _NotificationsService_NotifyOrderCanceled_Handler,
		},
		{
			MethodName: "NotifyOrderReady",
			Handler:    _NotificationsService_NotifyOrderReady_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notificationspb/api.proto",
}
