// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: equipment_service.proto

package equipment

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
	EquipmentService_GetById_FullMethodName = "/EquipmentService/GetById"
	EquipmentService_GetAll_FullMethodName  = "/EquipmentService/GetAll"
	EquipmentService_Create_FullMethodName  = "/EquipmentService/Create"
)

// EquipmentServiceClient is the client API for EquipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EquipmentServiceClient interface {
	GetById(ctx context.Context, in *EquipmentGetByIdRequest, opts ...grpc.CallOption) (*EquipmentGetByIdResponse, error)
	GetAll(ctx context.Context, in *EquipmentGetAllRequest, opts ...grpc.CallOption) (*EquipmentGetAllResponse, error)
	Create(ctx context.Context, in *EquipmentCreateRequest, opts ...grpc.CallOption) (*EquipmentCreateResponse, error)
}

type equipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquipmentServiceClient(cc grpc.ClientConnInterface) EquipmentServiceClient {
	return &equipmentServiceClient{cc}
}

func (c *equipmentServiceClient) GetById(ctx context.Context, in *EquipmentGetByIdRequest, opts ...grpc.CallOption) (*EquipmentGetByIdResponse, error) {
	out := new(EquipmentGetByIdResponse)
	err := c.cc.Invoke(ctx, EquipmentService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) GetAll(ctx context.Context, in *EquipmentGetAllRequest, opts ...grpc.CallOption) (*EquipmentGetAllResponse, error) {
	out := new(EquipmentGetAllResponse)
	err := c.cc.Invoke(ctx, EquipmentService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) Create(ctx context.Context, in *EquipmentCreateRequest, opts ...grpc.CallOption) (*EquipmentCreateResponse, error) {
	out := new(EquipmentCreateResponse)
	err := c.cc.Invoke(ctx, EquipmentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquipmentServiceServer is the server API for EquipmentService service.
// All implementations must embed UnimplementedEquipmentServiceServer
// for forward compatibility
type EquipmentServiceServer interface {
	GetById(context.Context, *EquipmentGetByIdRequest) (*EquipmentGetByIdResponse, error)
	GetAll(context.Context, *EquipmentGetAllRequest) (*EquipmentGetAllResponse, error)
	Create(context.Context, *EquipmentCreateRequest) (*EquipmentCreateResponse, error)
	mustEmbedUnimplementedEquipmentServiceServer()
}

// UnimplementedEquipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEquipmentServiceServer struct {
}

func (UnimplementedEquipmentServiceServer) GetById(context.Context, *EquipmentGetByIdRequest) (*EquipmentGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEquipmentServiceServer) GetAll(context.Context, *EquipmentGetAllRequest) (*EquipmentGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEquipmentServiceServer) Create(context.Context, *EquipmentCreateRequest) (*EquipmentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEquipmentServiceServer) mustEmbedUnimplementedEquipmentServiceServer() {}

// UnsafeEquipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EquipmentServiceServer will
// result in compilation errors.
type UnsafeEquipmentServiceServer interface {
	mustEmbedUnimplementedEquipmentServiceServer()
}

func RegisterEquipmentServiceServer(s grpc.ServiceRegistrar, srv EquipmentServiceServer) {
	s.RegisterService(&EquipmentService_ServiceDesc, srv)
}

func _EquipmentService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EquipmentService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).GetById(ctx, req.(*EquipmentGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EquipmentService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).GetAll(ctx, req.(*EquipmentGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EquipmentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).Create(ctx, req.(*EquipmentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EquipmentService_ServiceDesc is the grpc.ServiceDesc for EquipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EquipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EquipmentService",
	HandlerType: (*EquipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _EquipmentService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EquipmentService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EquipmentService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "equipment_service.proto",
}
