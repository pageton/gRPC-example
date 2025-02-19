// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/service.proto

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
	ImageProcessor_ProcessImage_FullMethodName = "/service.ImageProcessor/ProcessImage"
)

// ImageProcessorClient is the client API for ImageProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageProcessorClient interface {
	ProcessImage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type imageProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewImageProcessorClient(cc grpc.ClientConnInterface) ImageProcessorClient {
	return &imageProcessorClient{cc}
}

func (c *imageProcessorClient) ProcessImage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ImageProcessor_ProcessImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageProcessorServer is the server API for ImageProcessor service.
// All implementations must embed UnimplementedImageProcessorServer
// for forward compatibility.
type ImageProcessorServer interface {
	ProcessImage(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedImageProcessorServer()
}

// UnimplementedImageProcessorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageProcessorServer struct{}

func (UnimplementedImageProcessorServer) ProcessImage(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessImage not implemented")
}
func (UnimplementedImageProcessorServer) mustEmbedUnimplementedImageProcessorServer() {}
func (UnimplementedImageProcessorServer) testEmbeddedByValue()                        {}

// UnsafeImageProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageProcessorServer will
// result in compilation errors.
type UnsafeImageProcessorServer interface {
	mustEmbedUnimplementedImageProcessorServer()
}

func RegisterImageProcessorServer(s grpc.ServiceRegistrar, srv ImageProcessorServer) {
	// If the following call pancis, it indicates UnimplementedImageProcessorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageProcessor_ServiceDesc, srv)
}

func _ImageProcessor_ProcessImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessorServer).ProcessImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageProcessor_ProcessImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessorServer).ProcessImage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageProcessor_ServiceDesc is the grpc.ServiceDesc for ImageProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ImageProcessor",
	HandlerType: (*ImageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessImage",
			Handler:    _ImageProcessor_ProcessImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
