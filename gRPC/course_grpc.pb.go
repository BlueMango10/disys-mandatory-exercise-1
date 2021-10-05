// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// ITUCourseManagerClient is the client API for ITUCourseManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ITUCourseManagerClient interface {
	GetAllCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourseList, error)
	PostGourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error)
	GetCourse(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Course, error)
	PutCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error)
	DeleteCourse(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type iTUCourseManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewITUCourseManagerClient(cc grpc.ClientConnInterface) ITUCourseManagerClient {
	return &iTUCourseManagerClient{cc}
}

func (c *iTUCourseManagerClient) GetAllCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourseList, error) {
	out := new(CourseList)
	err := c.cc.Invoke(ctx, "/ITUCourseManager/GetAllCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iTUCourseManagerClient) PostGourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ITUCourseManager/PostGourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iTUCourseManagerClient) GetCourse(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/ITUCourseManager/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iTUCourseManagerClient) PutCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ITUCourseManager/PutCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iTUCourseManagerClient) DeleteCourse(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ITUCourseManager/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ITUCourseManagerServer is the server API for ITUCourseManager service.
// All implementations must embed UnimplementedITUCourseManagerServer
// for forward compatibility
type ITUCourseManagerServer interface {
	GetAllCourses(context.Context, *Empty) (*CourseList, error)
	PostGourse(context.Context, *Course) (*Empty, error)
	GetCourse(context.Context, *Id) (*Course, error)
	PutCourse(context.Context, *Course) (*Empty, error)
	DeleteCourse(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedITUCourseManagerServer()
}

// UnimplementedITUCourseManagerServer must be embedded to have forward compatible implementations.
type UnimplementedITUCourseManagerServer struct {
}

func (UnimplementedITUCourseManagerServer) GetAllCourses(context.Context, *Empty) (*CourseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCourses not implemented")
}
func (UnimplementedITUCourseManagerServer) PostGourse(context.Context, *Course) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGourse not implemented")
}
func (UnimplementedITUCourseManagerServer) GetCourse(context.Context, *Id) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedITUCourseManagerServer) PutCourse(context.Context, *Course) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCourse not implemented")
}
func (UnimplementedITUCourseManagerServer) DeleteCourse(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedITUCourseManagerServer) mustEmbedUnimplementedITUCourseManagerServer() {}

// UnsafeITUCourseManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ITUCourseManagerServer will
// result in compilation errors.
type UnsafeITUCourseManagerServer interface {
	mustEmbedUnimplementedITUCourseManagerServer()
}

func RegisterITUCourseManagerServer(s grpc.ServiceRegistrar, srv ITUCourseManagerServer) {
	s.RegisterService(&ITUCourseManager_ServiceDesc, srv)
}

func _ITUCourseManager_GetAllCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITUCourseManagerServer).GetAllCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ITUCourseManager/GetAllCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITUCourseManagerServer).GetAllCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ITUCourseManager_PostGourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITUCourseManagerServer).PostGourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ITUCourseManager/PostGourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITUCourseManagerServer).PostGourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _ITUCourseManager_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITUCourseManagerServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ITUCourseManager/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITUCourseManagerServer).GetCourse(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ITUCourseManager_PutCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITUCourseManagerServer).PutCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ITUCourseManager/PutCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITUCourseManagerServer).PutCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _ITUCourseManager_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITUCourseManagerServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ITUCourseManager/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITUCourseManagerServer).DeleteCourse(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ITUCourseManager_ServiceDesc is the grpc.ServiceDesc for ITUCourseManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ITUCourseManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ITUCourseManager",
	HandlerType: (*ITUCourseManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCourses",
			Handler:    _ITUCourseManager_GetAllCourses_Handler,
		},
		{
			MethodName: "PostGourse",
			Handler:    _ITUCourseManager_PostGourse_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _ITUCourseManager_GetCourse_Handler,
		},
		{
			MethodName: "PutCourse",
			Handler:    _ITUCourseManager_PutCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _ITUCourseManager_DeleteCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course.proto",
}