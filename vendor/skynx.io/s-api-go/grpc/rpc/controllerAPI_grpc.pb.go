// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: skynx/protobuf/rpc/v1/controllerAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "skynx.io/s-api-go/grpc/common/empty"
	account "skynx.io/s-api-go/grpc/resources/account"
	controller "skynx.io/s-api-go/grpc/resources/controller"
	topology "skynx.io/s-api-go/grpc/resources/topology"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ControllerAPI_SetupAccount_FullMethodName      = "/api.ControllerAPI/SetupAccount"
	ControllerAPI_GetAccountUsage_FullMethodName   = "/api.ControllerAPI/GetAccountUsage"
	ControllerAPI_GetAccountStats_FullMethodName   = "/api.ControllerAPI/GetAccountStats"
	ControllerAPI_GetNodeController_FullMethodName = "/api.ControllerAPI/GetNodeController"
)

// ControllerAPIClient is the client API for ControllerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [ctrl-api] ControllerAPI Definition: Controller Resources
type ControllerAPIClient interface {
	// account
	SetupAccount(ctx context.Context, in *account.SetupAccountRequest, opts ...grpc.CallOption) (*empty.Response, error)
	GetAccountUsage(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Usage, error)
	GetAccountStats(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Stats, error)
	GetNodeController(ctx context.Context, in *topology.NodeReq, opts ...grpc.CallOption) (*controller.Controller, error)
}

type controllerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerAPIClient(cc grpc.ClientConnInterface) ControllerAPIClient {
	return &controllerAPIClient{cc}
}

func (c *controllerAPIClient) SetupAccount(ctx context.Context, in *account.SetupAccountRequest, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, ControllerAPI_SetupAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAPIClient) GetAccountUsage(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Usage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(account.Usage)
	err := c.cc.Invoke(ctx, ControllerAPI_GetAccountUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAPIClient) GetAccountStats(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Stats, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(account.Stats)
	err := c.cc.Invoke(ctx, ControllerAPI_GetAccountStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerAPIClient) GetNodeController(ctx context.Context, in *topology.NodeReq, opts ...grpc.CallOption) (*controller.Controller, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(controller.Controller)
	err := c.cc.Invoke(ctx, ControllerAPI_GetNodeController_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerAPIServer is the server API for ControllerAPI service.
// All implementations must embed UnimplementedControllerAPIServer
// for forward compatibility
//
// [ctrl-api] ControllerAPI Definition: Controller Resources
type ControllerAPIServer interface {
	// account
	SetupAccount(context.Context, *account.SetupAccountRequest) (*empty.Response, error)
	GetAccountUsage(context.Context, *account.AccountReq) (*account.Usage, error)
	GetAccountStats(context.Context, *account.AccountReq) (*account.Stats, error)
	GetNodeController(context.Context, *topology.NodeReq) (*controller.Controller, error)
	mustEmbedUnimplementedControllerAPIServer()
}

// UnimplementedControllerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedControllerAPIServer struct {
}

func (UnimplementedControllerAPIServer) SetupAccount(context.Context, *account.SetupAccountRequest) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupAccount not implemented")
}
func (UnimplementedControllerAPIServer) GetAccountUsage(context.Context, *account.AccountReq) (*account.Usage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountUsage not implemented")
}
func (UnimplementedControllerAPIServer) GetAccountStats(context.Context, *account.AccountReq) (*account.Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountStats not implemented")
}
func (UnimplementedControllerAPIServer) GetNodeController(context.Context, *topology.NodeReq) (*controller.Controller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeController not implemented")
}
func (UnimplementedControllerAPIServer) mustEmbedUnimplementedControllerAPIServer() {}

// UnsafeControllerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerAPIServer will
// result in compilation errors.
type UnsafeControllerAPIServer interface {
	mustEmbedUnimplementedControllerAPIServer()
}

func RegisterControllerAPIServer(s grpc.ServiceRegistrar, srv ControllerAPIServer) {
	s.RegisterService(&ControllerAPI_ServiceDesc, srv)
}

func _ControllerAPI_SetupAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.SetupAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAPIServer).SetupAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerAPI_SetupAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAPIServer).SetupAccount(ctx, req.(*account.SetupAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAPI_GetAccountUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAPIServer).GetAccountUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerAPI_GetAccountUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAPIServer).GetAccountUsage(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAPI_GetAccountStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAPIServer).GetAccountStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerAPI_GetAccountStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAPIServer).GetAccountStats(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerAPI_GetNodeController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(topology.NodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerAPIServer).GetNodeController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerAPI_GetNodeController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerAPIServer).GetNodeController(ctx, req.(*topology.NodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerAPI_ServiceDesc is the grpc.ServiceDesc for ControllerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ControllerAPI",
	HandlerType: (*ControllerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupAccount",
			Handler:    _ControllerAPI_SetupAccount_Handler,
		},
		{
			MethodName: "GetAccountUsage",
			Handler:    _ControllerAPI_GetAccountUsage_Handler,
		},
		{
			MethodName: "GetAccountStats",
			Handler:    _ControllerAPI_GetAccountStats_Handler,
		},
		{
			MethodName: "GetNodeController",
			Handler:    _ControllerAPI_GetNodeController_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skynx/protobuf/rpc/v1/controllerAPI.proto",
}
