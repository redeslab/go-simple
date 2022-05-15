// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: cmdService.proto

package pbs

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

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmdServiceClient interface {
	SetLogLevel(ctx context.Context, in *LogLevel, opts ...grpc.CallOption) (*CommonResponse, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) SetLogLevel(ctx context.Context, in *LogLevel, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
// All implementations should embed UnimplementedCmdServiceServer
// for forward compatibility
type CmdServiceServer interface {
	SetLogLevel(context.Context, *LogLevel) (*CommonResponse, error)
}

// UnimplementedCmdServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (UnimplementedCmdServiceServer) SetLogLevel(context.Context, *LogLevel) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}

// UnsafeCmdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmdServiceServer will
// result in compilation errors.
type UnsafeCmdServiceServer interface {
	mustEmbedUnimplementedCmdServiceServer()
}

func RegisterCmdServiceServer(s grpc.ServiceRegistrar, srv CmdServiceServer) {
	s.RegisterService(&CmdService_ServiceDesc, srv)
}

func _CmdService_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).SetLogLevel(ctx, req.(*LogLevel))
	}
	return interceptor(ctx, in, info, handler)
}

// CmdService_ServiceDesc is the grpc.ServiceDesc for CmdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLogLevel",
			Handler:    _CmdService_SetLogLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdService.proto",
}