// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/order.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderService_PlaceOrder_FullMethodName      = "/order.OrderService/PlaceOrder"
	OrderService_CancelOrder_FullMethodName     = "/order.OrderService/CancelOrder"
	OrderService_GetOrderHistory_FullMethodName = "/order.OrderService/GetOrderHistory"
	OrderService_GetCurrentPrice_FullMethodName = "/order.OrderService/GetCurrentPrice"
	OrderService_CompleteOrder_FullMethodName   = "/order.OrderService/CompleteOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	GetOrderHistory(ctx context.Context, in *OrderHistoryRequest, opts ...grpc.CallOption) (*OrderHistoryResponse, error)
	GetCurrentPrice(ctx context.Context, in *GetCurrentPriceRequest, opts ...grpc.CallOption) (*GetCurrentPriceResponse, error)
	CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderHistory(ctx context.Context, in *OrderHistoryRequest, opts ...grpc.CallOption) (*OrderHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderHistoryResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetCurrentPrice(ctx context.Context, in *GetCurrentPriceRequest, opts ...grpc.CallOption) (*GetCurrentPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentPriceResponse)
	err := c.cc.Invoke(ctx, OrderService_GetCurrentPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CompleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	PlaceOrder(context.Context, *OrderRequest) (*OrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	GetOrderHistory(context.Context, *OrderHistoryRequest) (*OrderHistoryResponse, error)
	GetCurrentPrice(context.Context, *GetCurrentPriceRequest) (*GetCurrentPriceResponse, error)
	CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) PlaceOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderHistory(context.Context, *OrderHistoryRequest) (*OrderHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderHistory not implemented")
}
func (UnimplementedOrderServiceServer) GetCurrentPrice(context.Context, *GetCurrentPriceRequest) (*GetCurrentPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentPrice not implemented")
}
func (UnimplementedOrderServiceServer) CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PlaceOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderHistory(ctx, req.(*OrderHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetCurrentPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetCurrentPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetCurrentPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetCurrentPrice(ctx, req.(*GetCurrentPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CompleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CompleteOrder(ctx, req.(*CompleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderService_PlaceOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrderHistory",
			Handler:    _OrderService_GetOrderHistory_Handler,
		},
		{
			MethodName: "GetCurrentPrice",
			Handler:    _OrderService_GetCurrentPrice_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrderService_CompleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
