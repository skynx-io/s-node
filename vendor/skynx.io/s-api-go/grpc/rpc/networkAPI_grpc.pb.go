// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: skynx/protobuf/rpc/v1/networkAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "skynx.io/s-api-go/grpc/common/empty"
	nac "skynx.io/s-api-go/grpc/network/nac"
	sxsp "skynx.io/s-api-go/grpc/network/sxsp"
	controller "skynx.io/s-api-go/grpc/resources/controller"
	topology "skynx.io/s-api-go/grpc/resources/topology"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NetworkAPI_NetworkAdmissionControl_FullMethodName = "/network.NetworkAPI/NetworkAdmissionControl"
	NetworkAPI_NATProbe_FullMethodName                = "/network.NetworkAPI/NATProbe"
	NetworkAPI_RegisterEndpoint_FullMethodName        = "/network.NetworkAPI/RegisterEndpoint"
	NetworkAPI_RemoveEndpoint_FullMethodName          = "/network.NetworkAPI/RemoveEndpoint"
	NetworkAPI_RegisterNode_FullMethodName            = "/network.NetworkAPI/RegisterNode"
	NetworkAPI_Control_FullMethodName                 = "/network.NetworkAPI/Control"
	NetworkAPI_Metrics_FullMethodName                 = "/network.NetworkAPI/Metrics"
	NetworkAPI_FederationEndpoints_FullMethodName     = "/network.NetworkAPI/FederationEndpoints"
)

// NetworkAPIClient is the client API for NetworkAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [n-api] NetworkAPI Definition
type NetworkAPIClient interface {
	NetworkAdmissionControl(ctx context.Context, in *nac.NetworkAdmissionRequest, opts ...grpc.CallOption) (*nac.NetworkAdmissionResponse, error)
	NATProbe(ctx context.Context, in *nac.NATProbe, opts ...grpc.CallOption) (*nac.NATProbe, error)
	RegisterEndpoint(ctx context.Context, in *nac.EndpointRegRequest, opts ...grpc.CallOption) (*nac.EndpointRegResponse, error)
	RemoveEndpoint(ctx context.Context, in *topology.EndpointRequest, opts ...grpc.CallOption) (*topology.Node, error)
	RegisterNode(ctx context.Context, in *nac.NodeRegRequest, opts ...grpc.CallOption) (*nac.NodeRegResponse, error)
	// rpc Routing(stream routing.LSA) returns (stream routing.Status) {}
	// rpc RT(routing.RTRequest) returns (routing.RTResponse) {}
	Control(ctx context.Context, opts ...grpc.CallOption) (NetworkAPI_ControlClient, error)
	Metrics(ctx context.Context, in *topology.Node, opts ...grpc.CallOption) (*empty.Response, error)
	FederationEndpoints(ctx context.Context, in *topology.NodeReq, opts ...grpc.CallOption) (*controller.FederationEndpoints, error)
}

type networkAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkAPIClient(cc grpc.ClientConnInterface) NetworkAPIClient {
	return &networkAPIClient{cc}
}

func (c *networkAPIClient) NetworkAdmissionControl(ctx context.Context, in *nac.NetworkAdmissionRequest, opts ...grpc.CallOption) (*nac.NetworkAdmissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(nac.NetworkAdmissionResponse)
	err := c.cc.Invoke(ctx, NetworkAPI_NetworkAdmissionControl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) NATProbe(ctx context.Context, in *nac.NATProbe, opts ...grpc.CallOption) (*nac.NATProbe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(nac.NATProbe)
	err := c.cc.Invoke(ctx, NetworkAPI_NATProbe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) RegisterEndpoint(ctx context.Context, in *nac.EndpointRegRequest, opts ...grpc.CallOption) (*nac.EndpointRegResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(nac.EndpointRegResponse)
	err := c.cc.Invoke(ctx, NetworkAPI_RegisterEndpoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) RemoveEndpoint(ctx context.Context, in *topology.EndpointRequest, opts ...grpc.CallOption) (*topology.Node, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(topology.Node)
	err := c.cc.Invoke(ctx, NetworkAPI_RemoveEndpoint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) RegisterNode(ctx context.Context, in *nac.NodeRegRequest, opts ...grpc.CallOption) (*nac.NodeRegResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(nac.NodeRegResponse)
	err := c.cc.Invoke(ctx, NetworkAPI_RegisterNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) Control(ctx context.Context, opts ...grpc.CallOption) (NetworkAPI_ControlClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NetworkAPI_ServiceDesc.Streams[0], NetworkAPI_Control_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &networkAPIControlClient{ClientStream: stream}
	return x, nil
}

type NetworkAPI_ControlClient interface {
	Send(*sxsp.Payload) error
	Recv() (*sxsp.Payload, error)
	grpc.ClientStream
}

type networkAPIControlClient struct {
	grpc.ClientStream
}

func (x *networkAPIControlClient) Send(m *sxsp.Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkAPIControlClient) Recv() (*sxsp.Payload, error) {
	m := new(sxsp.Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkAPIClient) Metrics(ctx context.Context, in *topology.Node, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, NetworkAPI_Metrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAPIClient) FederationEndpoints(ctx context.Context, in *topology.NodeReq, opts ...grpc.CallOption) (*controller.FederationEndpoints, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(controller.FederationEndpoints)
	err := c.cc.Invoke(ctx, NetworkAPI_FederationEndpoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkAPIServer is the server API for NetworkAPI service.
// All implementations must embed UnimplementedNetworkAPIServer
// for forward compatibility
//
// [n-api] NetworkAPI Definition
type NetworkAPIServer interface {
	NetworkAdmissionControl(context.Context, *nac.NetworkAdmissionRequest) (*nac.NetworkAdmissionResponse, error)
	NATProbe(context.Context, *nac.NATProbe) (*nac.NATProbe, error)
	RegisterEndpoint(context.Context, *nac.EndpointRegRequest) (*nac.EndpointRegResponse, error)
	RemoveEndpoint(context.Context, *topology.EndpointRequest) (*topology.Node, error)
	RegisterNode(context.Context, *nac.NodeRegRequest) (*nac.NodeRegResponse, error)
	// rpc Routing(stream routing.LSA) returns (stream routing.Status) {}
	// rpc RT(routing.RTRequest) returns (routing.RTResponse) {}
	Control(NetworkAPI_ControlServer) error
	Metrics(context.Context, *topology.Node) (*empty.Response, error)
	FederationEndpoints(context.Context, *topology.NodeReq) (*controller.FederationEndpoints, error)
	mustEmbedUnimplementedNetworkAPIServer()
}

// UnimplementedNetworkAPIServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkAPIServer struct {
}

func (UnimplementedNetworkAPIServer) NetworkAdmissionControl(context.Context, *nac.NetworkAdmissionRequest) (*nac.NetworkAdmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkAdmissionControl not implemented")
}
func (UnimplementedNetworkAPIServer) NATProbe(context.Context, *nac.NATProbe) (*nac.NATProbe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NATProbe not implemented")
}
func (UnimplementedNetworkAPIServer) RegisterEndpoint(context.Context, *nac.EndpointRegRequest) (*nac.EndpointRegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEndpoint not implemented")
}
func (UnimplementedNetworkAPIServer) RemoveEndpoint(context.Context, *topology.EndpointRequest) (*topology.Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEndpoint not implemented")
}
func (UnimplementedNetworkAPIServer) RegisterNode(context.Context, *nac.NodeRegRequest) (*nac.NodeRegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (UnimplementedNetworkAPIServer) Control(NetworkAPI_ControlServer) error {
	return status.Errorf(codes.Unimplemented, "method Control not implemented")
}
func (UnimplementedNetworkAPIServer) Metrics(context.Context, *topology.Node) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (UnimplementedNetworkAPIServer) FederationEndpoints(context.Context, *topology.NodeReq) (*controller.FederationEndpoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FederationEndpoints not implemented")
}
func (UnimplementedNetworkAPIServer) mustEmbedUnimplementedNetworkAPIServer() {}

// UnsafeNetworkAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkAPIServer will
// result in compilation errors.
type UnsafeNetworkAPIServer interface {
	mustEmbedUnimplementedNetworkAPIServer()
}

func RegisterNetworkAPIServer(s grpc.ServiceRegistrar, srv NetworkAPIServer) {
	s.RegisterService(&NetworkAPI_ServiceDesc, srv)
}

func _NetworkAPI_NetworkAdmissionControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nac.NetworkAdmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).NetworkAdmissionControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_NetworkAdmissionControl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).NetworkAdmissionControl(ctx, req.(*nac.NetworkAdmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_NATProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nac.NATProbe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).NATProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_NATProbe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).NATProbe(ctx, req.(*nac.NATProbe))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_RegisterEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nac.EndpointRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).RegisterEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_RegisterEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).RegisterEndpoint(ctx, req.(*nac.EndpointRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_RemoveEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(topology.EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).RemoveEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_RemoveEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).RemoveEndpoint(ctx, req.(*topology.EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nac.NodeRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_RegisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).RegisterNode(ctx, req.(*nac.NodeRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_Control_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkAPIServer).Control(&networkAPIControlServer{ServerStream: stream})
}

type NetworkAPI_ControlServer interface {
	Send(*sxsp.Payload) error
	Recv() (*sxsp.Payload, error)
	grpc.ServerStream
}

type networkAPIControlServer struct {
	grpc.ServerStream
}

func (x *networkAPIControlServer) Send(m *sxsp.Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkAPIControlServer) Recv() (*sxsp.Payload, error) {
	m := new(sxsp.Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkAPI_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(topology.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_Metrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).Metrics(ctx, req.(*topology.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAPI_FederationEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(topology.NodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAPIServer).FederationEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkAPI_FederationEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAPIServer).FederationEndpoints(ctx, req.(*topology.NodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkAPI_ServiceDesc is the grpc.ServiceDesc for NetworkAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.NetworkAPI",
	HandlerType: (*NetworkAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NetworkAdmissionControl",
			Handler:    _NetworkAPI_NetworkAdmissionControl_Handler,
		},
		{
			MethodName: "NATProbe",
			Handler:    _NetworkAPI_NATProbe_Handler,
		},
		{
			MethodName: "RegisterEndpoint",
			Handler:    _NetworkAPI_RegisterEndpoint_Handler,
		},
		{
			MethodName: "RemoveEndpoint",
			Handler:    _NetworkAPI_RemoveEndpoint_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _NetworkAPI_RegisterNode_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _NetworkAPI_Metrics_Handler,
		},
		{
			MethodName: "FederationEndpoints",
			Handler:    _NetworkAPI_FederationEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Control",
			Handler:       _NetworkAPI_Control_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "skynx/protobuf/rpc/v1/networkAPI.proto",
}
