// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: admin/v1/feature_flag.proto

package v1

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

const (
	FeatureFlag_CreateFeatureFlag_FullMethodName = "/api.admin.v1.FeatureFlag/CreateFeatureFlag"
	FeatureFlag_UpdateFeatureFlag_FullMethodName = "/api.admin.v1.FeatureFlag/UpdateFeatureFlag"
	FeatureFlag_DeleteFeatureFlag_FullMethodName = "/api.admin.v1.FeatureFlag/DeleteFeatureFlag"
	FeatureFlag_GetFeatureFlag_FullMethodName    = "/api.admin.v1.FeatureFlag/GetFeatureFlag"
	FeatureFlag_ListFeatureFlag_FullMethodName   = "/api.admin.v1.FeatureFlag/ListFeatureFlag"
)

// FeatureFlagClient is the client API for FeatureFlag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureFlagClient interface {
	CreateFeatureFlag(ctx context.Context, in *CreateFeatureFlagRequest, opts ...grpc.CallOption) (*CreateFeatureFlagReply, error)
	UpdateFeatureFlag(ctx context.Context, in *UpdateFeatureFlagRequest, opts ...grpc.CallOption) (*UpdateFeatureFlagReply, error)
	DeleteFeatureFlag(ctx context.Context, in *DeleteFeatureFlagRequest, opts ...grpc.CallOption) (*DeleteFeatureFlagReply, error)
	GetFeatureFlag(ctx context.Context, in *GetFeatureFlagRequest, opts ...grpc.CallOption) (*GetFeatureFlagReply, error)
	ListFeatureFlag(ctx context.Context, in *ListFeatureFlagRequest, opts ...grpc.CallOption) (*ListFeatureFlagReply, error)
}

type featureFlagClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureFlagClient(cc grpc.ClientConnInterface) FeatureFlagClient {
	return &featureFlagClient{cc}
}

func (c *featureFlagClient) CreateFeatureFlag(ctx context.Context, in *CreateFeatureFlagRequest, opts ...grpc.CallOption) (*CreateFeatureFlagReply, error) {
	out := new(CreateFeatureFlagReply)
	err := c.cc.Invoke(ctx, FeatureFlag_CreateFeatureFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureFlagClient) UpdateFeatureFlag(ctx context.Context, in *UpdateFeatureFlagRequest, opts ...grpc.CallOption) (*UpdateFeatureFlagReply, error) {
	out := new(UpdateFeatureFlagReply)
	err := c.cc.Invoke(ctx, FeatureFlag_UpdateFeatureFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureFlagClient) DeleteFeatureFlag(ctx context.Context, in *DeleteFeatureFlagRequest, opts ...grpc.CallOption) (*DeleteFeatureFlagReply, error) {
	out := new(DeleteFeatureFlagReply)
	err := c.cc.Invoke(ctx, FeatureFlag_DeleteFeatureFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureFlagClient) GetFeatureFlag(ctx context.Context, in *GetFeatureFlagRequest, opts ...grpc.CallOption) (*GetFeatureFlagReply, error) {
	out := new(GetFeatureFlagReply)
	err := c.cc.Invoke(ctx, FeatureFlag_GetFeatureFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureFlagClient) ListFeatureFlag(ctx context.Context, in *ListFeatureFlagRequest, opts ...grpc.CallOption) (*ListFeatureFlagReply, error) {
	out := new(ListFeatureFlagReply)
	err := c.cc.Invoke(ctx, FeatureFlag_ListFeatureFlag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureFlagServer is the server API for FeatureFlag service.
// All implementations must embed UnimplementedFeatureFlagServer
// for forward compatibility
type FeatureFlagServer interface {
	CreateFeatureFlag(context.Context, *CreateFeatureFlagRequest) (*CreateFeatureFlagReply, error)
	UpdateFeatureFlag(context.Context, *UpdateFeatureFlagRequest) (*UpdateFeatureFlagReply, error)
	DeleteFeatureFlag(context.Context, *DeleteFeatureFlagRequest) (*DeleteFeatureFlagReply, error)
	GetFeatureFlag(context.Context, *GetFeatureFlagRequest) (*GetFeatureFlagReply, error)
	ListFeatureFlag(context.Context, *ListFeatureFlagRequest) (*ListFeatureFlagReply, error)
	mustEmbedUnimplementedFeatureFlagServer()
}

// UnimplementedFeatureFlagServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureFlagServer struct {
}

func (UnimplementedFeatureFlagServer) CreateFeatureFlag(context.Context, *CreateFeatureFlagRequest) (*CreateFeatureFlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeatureFlag not implemented")
}
func (UnimplementedFeatureFlagServer) UpdateFeatureFlag(context.Context, *UpdateFeatureFlagRequest) (*UpdateFeatureFlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeatureFlag not implemented")
}
func (UnimplementedFeatureFlagServer) DeleteFeatureFlag(context.Context, *DeleteFeatureFlagRequest) (*DeleteFeatureFlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeatureFlag not implemented")
}
func (UnimplementedFeatureFlagServer) GetFeatureFlag(context.Context, *GetFeatureFlagRequest) (*GetFeatureFlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureFlag not implemented")
}
func (UnimplementedFeatureFlagServer) ListFeatureFlag(context.Context, *ListFeatureFlagRequest) (*ListFeatureFlagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatureFlag not implemented")
}
func (UnimplementedFeatureFlagServer) mustEmbedUnimplementedFeatureFlagServer() {}

// UnsafeFeatureFlagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureFlagServer will
// result in compilation errors.
type UnsafeFeatureFlagServer interface {
	mustEmbedUnimplementedFeatureFlagServer()
}

func RegisterFeatureFlagServer(s grpc.ServiceRegistrar, srv FeatureFlagServer) {
	s.RegisterService(&FeatureFlag_ServiceDesc, srv)
}

func _FeatureFlag_CreateFeatureFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeatureFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).CreateFeatureFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlag_CreateFeatureFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).CreateFeatureFlag(ctx, req.(*CreateFeatureFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureFlag_UpdateFeatureFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeatureFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).UpdateFeatureFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlag_UpdateFeatureFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).UpdateFeatureFlag(ctx, req.(*UpdateFeatureFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureFlag_DeleteFeatureFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeatureFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).DeleteFeatureFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlag_DeleteFeatureFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).DeleteFeatureFlag(ctx, req.(*DeleteFeatureFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureFlag_GetFeatureFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).GetFeatureFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlag_GetFeatureFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).GetFeatureFlag(ctx, req.(*GetFeatureFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureFlag_ListFeatureFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeatureFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).ListFeatureFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureFlag_ListFeatureFlag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).ListFeatureFlag(ctx, req.(*ListFeatureFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureFlag_ServiceDesc is the grpc.ServiceDesc for FeatureFlag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureFlag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.FeatureFlag",
	HandlerType: (*FeatureFlagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeatureFlag",
			Handler:    _FeatureFlag_CreateFeatureFlag_Handler,
		},
		{
			MethodName: "UpdateFeatureFlag",
			Handler:    _FeatureFlag_UpdateFeatureFlag_Handler,
		},
		{
			MethodName: "DeleteFeatureFlag",
			Handler:    _FeatureFlag_DeleteFeatureFlag_Handler,
		},
		{
			MethodName: "GetFeatureFlag",
			Handler:    _FeatureFlag_GetFeatureFlag_Handler,
		},
		{
			MethodName: "ListFeatureFlag",
			Handler:    _FeatureFlag_ListFeatureFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/feature_flag.proto",
}