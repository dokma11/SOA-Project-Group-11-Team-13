// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: keypoints_service.proto

package keypoints

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
	KeyPointsService_GetById_FullMethodName     = "/KeyPointsService/GetById"
	KeyPointsService_GetAll_FullMethodName      = "/KeyPointsService/GetAll"
	KeyPointsService_GetByTourId_FullMethodName = "/KeyPointsService/GetByTourId"
	KeyPointsService_Create_FullMethodName      = "/KeyPointsService/Create"
	KeyPointsService_Update_FullMethodName      = "/KeyPointsService/Update"
	KeyPointsService_Delete_FullMethodName      = "/KeyPointsService/Delete"
)

// KeyPointsServiceClient is the client API for KeyPointsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyPointsServiceClient interface {
	GetById(ctx context.Context, in *KeyPointGetByIdRequest, opts ...grpc.CallOption) (*KeyPointGetByIdResponse, error)
	GetAll(ctx context.Context, in *KeyPointGetAllRequest, opts ...grpc.CallOption) (*KeyPointGetAllResponse, error)
	GetByTourId(ctx context.Context, in *KeyPointGetByTourIdRequest, opts ...grpc.CallOption) (*KeyPointGetByTourIdResponse, error)
	Create(ctx context.Context, in *KeyPointCreateRequest, opts ...grpc.CallOption) (*KeyPointCreateResponse, error)
	Update(ctx context.Context, in *KeyPointUpdateRequest, opts ...grpc.CallOption) (*KeyPointUpdateResponse, error)
	Delete(ctx context.Context, in *KeyPointDeleteRequest, opts ...grpc.CallOption) (*KeyPointDeleteResponse, error)
}

type keyPointsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyPointsServiceClient(cc grpc.ClientConnInterface) KeyPointsServiceClient {
	return &keyPointsServiceClient{cc}
}

func (c *keyPointsServiceClient) GetById(ctx context.Context, in *KeyPointGetByIdRequest, opts ...grpc.CallOption) (*KeyPointGetByIdResponse, error) {
	out := new(KeyPointGetByIdResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyPointsServiceClient) GetAll(ctx context.Context, in *KeyPointGetAllRequest, opts ...grpc.CallOption) (*KeyPointGetAllResponse, error) {
	out := new(KeyPointGetAllResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyPointsServiceClient) GetByTourId(ctx context.Context, in *KeyPointGetByTourIdRequest, opts ...grpc.CallOption) (*KeyPointGetByTourIdResponse, error) {
	out := new(KeyPointGetByTourIdResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_GetByTourId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyPointsServiceClient) Create(ctx context.Context, in *KeyPointCreateRequest, opts ...grpc.CallOption) (*KeyPointCreateResponse, error) {
	out := new(KeyPointCreateResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyPointsServiceClient) Update(ctx context.Context, in *KeyPointUpdateRequest, opts ...grpc.CallOption) (*KeyPointUpdateResponse, error) {
	out := new(KeyPointUpdateResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyPointsServiceClient) Delete(ctx context.Context, in *KeyPointDeleteRequest, opts ...grpc.CallOption) (*KeyPointDeleteResponse, error) {
	out := new(KeyPointDeleteResponse)
	err := c.cc.Invoke(ctx, KeyPointsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyPointsServiceServer is the server API for KeyPointsService service.
// All implementations must embed UnimplementedKeyPointsServiceServer
// for forward compatibility
type KeyPointsServiceServer interface {
	GetById(context.Context, *KeyPointGetByIdRequest) (*KeyPointGetByIdResponse, error)
	GetAll(context.Context, *KeyPointGetAllRequest) (*KeyPointGetAllResponse, error)
	GetByTourId(context.Context, *KeyPointGetByTourIdRequest) (*KeyPointGetByTourIdResponse, error)
	Create(context.Context, *KeyPointCreateRequest) (*KeyPointCreateResponse, error)
	Update(context.Context, *KeyPointUpdateRequest) (*KeyPointUpdateResponse, error)
	Delete(context.Context, *KeyPointDeleteRequest) (*KeyPointDeleteResponse, error)
	mustEmbedUnimplementedKeyPointsServiceServer()
}

// UnimplementedKeyPointsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyPointsServiceServer struct {
}

func (UnimplementedKeyPointsServiceServer) GetById(context.Context, *KeyPointGetByIdRequest) (*KeyPointGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedKeyPointsServiceServer) GetAll(context.Context, *KeyPointGetAllRequest) (*KeyPointGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedKeyPointsServiceServer) GetByTourId(context.Context, *KeyPointGetByTourIdRequest) (*KeyPointGetByTourIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTourId not implemented")
}
func (UnimplementedKeyPointsServiceServer) Create(context.Context, *KeyPointCreateRequest) (*KeyPointCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKeyPointsServiceServer) Update(context.Context, *KeyPointUpdateRequest) (*KeyPointUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKeyPointsServiceServer) Delete(context.Context, *KeyPointDeleteRequest) (*KeyPointDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKeyPointsServiceServer) mustEmbedUnimplementedKeyPointsServiceServer() {}

// UnsafeKeyPointsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyPointsServiceServer will
// result in compilation errors.
type UnsafeKeyPointsServiceServer interface {
	mustEmbedUnimplementedKeyPointsServiceServer()
}

func RegisterKeyPointsServiceServer(s grpc.ServiceRegistrar, srv KeyPointsServiceServer) {
	s.RegisterService(&KeyPointsService_ServiceDesc, srv)
}

func _KeyPointsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).GetById(ctx, req.(*KeyPointGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyPointsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).GetAll(ctx, req.(*KeyPointGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyPointsService_GetByTourId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointGetByTourIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).GetByTourId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_GetByTourId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).GetByTourId(ctx, req.(*KeyPointGetByTourIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyPointsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).Create(ctx, req.(*KeyPointCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyPointsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).Update(ctx, req.(*KeyPointUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyPointsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyPointsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyPointsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyPointsServiceServer).Delete(ctx, req.(*KeyPointDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyPointsService_ServiceDesc is the grpc.ServiceDesc for KeyPointsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyPointsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KeyPointsService",
	HandlerType: (*KeyPointsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _KeyPointsService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _KeyPointsService_GetAll_Handler,
		},
		{
			MethodName: "GetByTourId",
			Handler:    _KeyPointsService_GetByTourId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _KeyPointsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _KeyPointsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KeyPointsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keypoints_service.proto",
}
