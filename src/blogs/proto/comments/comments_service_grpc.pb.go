// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: comments_service.proto

package comments

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
	CommentsService_GetCommentById_FullMethodName     = "/CommentsService/GetCommentById"
	CommentsService_GetAllComments_FullMethodName     = "/CommentsService/GetAllComments"
	CommentsService_GetCommentByBlogId_FullMethodName = "/CommentsService/GetCommentByBlogId"
	CommentsService_CreateComment_FullMethodName      = "/CommentsService/CreateComment"
	CommentsService_DeleteComment_FullMethodName      = "/CommentsService/DeleteComment"
	CommentsService_UpdateComment_FullMethodName      = "/CommentsService/UpdateComment"
)

// CommentsServiceClient is the client API for CommentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsServiceClient interface {
	GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error)
	GetAllComments(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error)
	GetCommentByBlogId(ctx context.Context, in *GetCommentByBlogIdRequest, opts ...grpc.CallOption) (*GetCommentByBlogIdResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
}

type commentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsServiceClient(cc grpc.ClientConnInterface) CommentsServiceClient {
	return &commentsServiceClient{cc}
}

func (c *commentsServiceClient) GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error) {
	out := new(GetCommentByIdResponse)
	err := c.cc.Invoke(ctx, CommentsService_GetCommentById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServiceClient) GetAllComments(ctx context.Context, in *GetAllCommentsRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error) {
	out := new(GetAllCommentsResponse)
	err := c.cc.Invoke(ctx, CommentsService_GetAllComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServiceClient) GetCommentByBlogId(ctx context.Context, in *GetCommentByBlogIdRequest, opts ...grpc.CallOption) (*GetCommentByBlogIdResponse, error) {
	out := new(GetCommentByBlogIdResponse)
	err := c.cc.Invoke(ctx, CommentsService_GetCommentByBlogId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, CommentsService_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, CommentsService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, CommentsService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsServiceServer is the server API for CommentsService service.
// All implementations must embed UnimplementedCommentsServiceServer
// for forward compatibility
type CommentsServiceServer interface {
	GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error)
	GetAllComments(context.Context, *GetAllCommentsRequest) (*GetAllCommentsResponse, error)
	GetCommentByBlogId(context.Context, *GetCommentByBlogIdRequest) (*GetCommentByBlogIdResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	mustEmbedUnimplementedCommentsServiceServer()
}

// UnimplementedCommentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentsServiceServer struct {
}

func (UnimplementedCommentsServiceServer) GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentsServiceServer) GetAllComments(context.Context, *GetAllCommentsRequest) (*GetAllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllComments not implemented")
}
func (UnimplementedCommentsServiceServer) GetCommentByBlogId(context.Context, *GetCommentByBlogIdRequest) (*GetCommentByBlogIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentByBlogId not implemented")
}
func (UnimplementedCommentsServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentsServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentsServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentsServiceServer) mustEmbedUnimplementedCommentsServiceServer() {}

// UnsafeCommentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsServiceServer will
// result in compilation errors.
type UnsafeCommentsServiceServer interface {
	mustEmbedUnimplementedCommentsServiceServer()
}

func RegisterCommentsServiceServer(s grpc.ServiceRegistrar, srv CommentsServiceServer) {
	s.RegisterService(&CommentsService_ServiceDesc, srv)
}

func _CommentsService_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).GetCommentById(ctx, req.(*GetCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsService_GetAllComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).GetAllComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_GetAllComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).GetAllComments(ctx, req.(*GetAllCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsService_GetCommentByBlogId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByBlogIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).GetCommentByBlogId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_GetCommentByBlogId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).GetCommentByBlogId(ctx, req.(*GetCommentByBlogIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentsService_ServiceDesc is the grpc.ServiceDesc for CommentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommentsService",
	HandlerType: (*CommentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommentById",
			Handler:    _CommentsService_GetCommentById_Handler,
		},
		{
			MethodName: "GetAllComments",
			Handler:    _CommentsService_GetAllComments_Handler,
		},
		{
			MethodName: "GetCommentByBlogId",
			Handler:    _CommentsService_GetCommentByBlogId_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CommentsService_CreateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentsService_DeleteComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentsService_UpdateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments_service.proto",
}
