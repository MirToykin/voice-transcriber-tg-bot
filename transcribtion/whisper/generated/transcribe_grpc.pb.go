// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: transcribe.proto

package generated

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

// TranscriptionServiceClient is the client API for TranscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscriptionServiceClient interface {
	// Endpoint to transcribe audio by file path
	TranscribeByPath(ctx context.Context, in *TranscribeByPathRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error)
	// Endpoint to transcribe audio by binary data
	TranscribeByBinary(ctx context.Context, in *TranscribeByBinaryRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error)
	// Endpoint to get available languages list
	GetAvailableLanguages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AvailableLanguagesResponse, error)
}

type transcriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscriptionServiceClient(cc grpc.ClientConnInterface) TranscriptionServiceClient {
	return &transcriptionServiceClient{cc}
}

func (c *transcriptionServiceClient) TranscribeByPath(ctx context.Context, in *TranscribeByPathRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error) {
	out := new(TranscriptionResponse)
	err := c.cc.Invoke(ctx, "/TranscriptionService/TranscribeByPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcriptionServiceClient) TranscribeByBinary(ctx context.Context, in *TranscribeByBinaryRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error) {
	out := new(TranscriptionResponse)
	err := c.cc.Invoke(ctx, "/TranscriptionService/TranscribeByBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcriptionServiceClient) GetAvailableLanguages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AvailableLanguagesResponse, error) {
	out := new(AvailableLanguagesResponse)
	err := c.cc.Invoke(ctx, "/TranscriptionService/GetAvailableLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranscriptionServiceServer is the server API for TranscriptionService service.
// All implementations must embed UnimplementedTranscriptionServiceServer
// for forward compatibility
type TranscriptionServiceServer interface {
	// Endpoint to transcribe audio by file path
	TranscribeByPath(context.Context, *TranscribeByPathRequest) (*TranscriptionResponse, error)
	// Endpoint to transcribe audio by binary data
	TranscribeByBinary(context.Context, *TranscribeByBinaryRequest) (*TranscriptionResponse, error)
	// Endpoint to get available languages list
	GetAvailableLanguages(context.Context, *emptypb.Empty) (*AvailableLanguagesResponse, error)
	mustEmbedUnimplementedTranscriptionServiceServer()
}

// UnimplementedTranscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTranscriptionServiceServer struct {
}

func (UnimplementedTranscriptionServiceServer) TranscribeByPath(context.Context, *TranscribeByPathRequest) (*TranscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscribeByPath not implemented")
}
func (UnimplementedTranscriptionServiceServer) TranscribeByBinary(context.Context, *TranscribeByBinaryRequest) (*TranscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscribeByBinary not implemented")
}
func (UnimplementedTranscriptionServiceServer) GetAvailableLanguages(context.Context, *emptypb.Empty) (*AvailableLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableLanguages not implemented")
}
func (UnimplementedTranscriptionServiceServer) mustEmbedUnimplementedTranscriptionServiceServer() {}

// UnsafeTranscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscriptionServiceServer will
// result in compilation errors.
type UnsafeTranscriptionServiceServer interface {
	mustEmbedUnimplementedTranscriptionServiceServer()
}

func RegisterTranscriptionServiceServer(s grpc.ServiceRegistrar, srv TranscriptionServiceServer) {
	s.RegisterService(&TranscriptionService_ServiceDesc, srv)
}

func _TranscriptionService_TranscribeByPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscribeByPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscriptionServiceServer).TranscribeByPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TranscriptionService/TranscribeByPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscriptionServiceServer).TranscribeByPath(ctx, req.(*TranscribeByPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranscriptionService_TranscribeByBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscribeByBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscriptionServiceServer).TranscribeByBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TranscriptionService/TranscribeByBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscriptionServiceServer).TranscribeByBinary(ctx, req.(*TranscribeByBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranscriptionService_GetAvailableLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscriptionServiceServer).GetAvailableLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TranscriptionService/GetAvailableLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscriptionServiceServer).GetAvailableLanguages(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TranscriptionService_ServiceDesc is the grpc.ServiceDesc for TranscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TranscriptionService",
	HandlerType: (*TranscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranscribeByPath",
			Handler:    _TranscriptionService_TranscribeByPath_Handler,
		},
		{
			MethodName: "TranscribeByBinary",
			Handler:    _TranscriptionService_TranscribeByBinary_Handler,
		},
		{
			MethodName: "GetAvailableLanguages",
			Handler:    _TranscriptionService_GetAvailableLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transcribe.proto",
}
