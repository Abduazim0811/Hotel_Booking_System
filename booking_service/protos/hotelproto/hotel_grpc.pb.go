// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: hotelproto/hotel.proto

package hotelproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HotelService_CreateHotel_FullMethodName  = "/HotelService/CreateHotel"
	HotelService_GetbyIdHotel_FullMethodName = "/HotelService/GetbyIdHotel"
	HotelService_GetAllHotels_FullMethodName = "/HotelService/GetAllHotels"
	HotelService_UpdateHotel_FullMethodName  = "/HotelService/UpdateHotel"
	HotelService_DeleteHotel_FullMethodName  = "/HotelService/DeleteHotel"
	HotelService_CreateRoom_FullMethodName   = "/HotelService/CreateRoom"
	HotelService_GetbyIdRoom_FullMethodName  = "/HotelService/GetbyIdRoom"
	HotelService_GetAllRooms_FullMethodName  = "/HotelService/GetAllRooms"
	HotelService_UpdateRoom_FullMethodName   = "/HotelService/UpdateRoom"
	HotelService_DeleteRoom_FullMethodName   = "/HotelService/DeleteRoom"
)

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	CreateHotel(ctx context.Context, in *HotelRequest, opts ...grpc.CallOption) (*HotelResponse, error)
	GetbyIdHotel(ctx context.Context, in *HotelResponse, opts ...grpc.CallOption) (*Hotel, error)
	GetAllHotels(ctx context.Context, in *HotelEmpty, opts ...grpc.CallOption) (*ListHotels, error)
	UpdateHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*HotelRes, error)
	DeleteHotel(ctx context.Context, in *HotelResponse, opts ...grpc.CallOption) (*HotelRes, error)
	CreateRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
	GetbyIdRoom(ctx context.Context, in *RoomResponse, opts ...grpc.CallOption) (*Room, error)
	GetAllRooms(ctx context.Context, in *HotelEmpty, opts ...grpc.CallOption) (*ListRooms, error)
	UpdateRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*RoomRes, error)
	DeleteRoom(ctx context.Context, in *RoomResponse, opts ...grpc.CallOption) (*RoomRes, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) CreateHotel(ctx context.Context, in *HotelRequest, opts ...grpc.CallOption) (*HotelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HotelResponse)
	err := c.cc.Invoke(ctx, HotelService_CreateHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetbyIdHotel(ctx context.Context, in *HotelResponse, opts ...grpc.CallOption) (*Hotel, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hotel)
	err := c.cc.Invoke(ctx, HotelService_GetbyIdHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetAllHotels(ctx context.Context, in *HotelEmpty, opts ...grpc.CallOption) (*ListHotels, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHotels)
	err := c.cc.Invoke(ctx, HotelService_GetAllHotels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*HotelRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HotelRes)
	err := c.cc.Invoke(ctx, HotelService_UpdateHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteHotel(ctx context.Context, in *HotelResponse, opts ...grpc.CallOption) (*HotelRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HotelRes)
	err := c.cc.Invoke(ctx, HotelService_DeleteHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) CreateRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, HotelService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetbyIdRoom(ctx context.Context, in *RoomResponse, opts ...grpc.CallOption) (*Room, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Room)
	err := c.cc.Invoke(ctx, HotelService_GetbyIdRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetAllRooms(ctx context.Context, in *HotelEmpty, opts ...grpc.CallOption) (*ListRooms, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRooms)
	err := c.cc.Invoke(ctx, HotelService_GetAllRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*RoomRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomRes)
	err := c.cc.Invoke(ctx, HotelService_UpdateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteRoom(ctx context.Context, in *RoomResponse, opts ...grpc.CallOption) (*RoomRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomRes)
	err := c.cc.Invoke(ctx, HotelService_DeleteRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility
type HotelServiceServer interface {
	CreateHotel(context.Context, *HotelRequest) (*HotelResponse, error)
	GetbyIdHotel(context.Context, *HotelResponse) (*Hotel, error)
	GetAllHotels(context.Context, *HotelEmpty) (*ListHotels, error)
	UpdateHotel(context.Context, *Hotel) (*HotelRes, error)
	DeleteHotel(context.Context, *HotelResponse) (*HotelRes, error)
	CreateRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	GetbyIdRoom(context.Context, *RoomResponse) (*Room, error)
	GetAllRooms(context.Context, *HotelEmpty) (*ListRooms, error)
	UpdateRoom(context.Context, *Room) (*RoomRes, error)
	DeleteRoom(context.Context, *RoomResponse) (*RoomRes, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelServiceServer struct {
}

func (UnimplementedHotelServiceServer) CreateHotel(context.Context, *HotelRequest) (*HotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetbyIdHotel(context.Context, *HotelResponse) (*Hotel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetbyIdHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetAllHotels(context.Context, *HotelEmpty) (*ListHotels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHotels not implemented")
}
func (UnimplementedHotelServiceServer) UpdateHotel(context.Context, *Hotel) (*HotelRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHotel not implemented")
}
func (UnimplementedHotelServiceServer) DeleteHotel(context.Context, *HotelResponse) (*HotelRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotel not implemented")
}
func (UnimplementedHotelServiceServer) CreateRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedHotelServiceServer) GetbyIdRoom(context.Context, *RoomResponse) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetbyIdRoom not implemented")
}
func (UnimplementedHotelServiceServer) GetAllRooms(context.Context, *HotelEmpty) (*ListRooms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRooms not implemented")
}
func (UnimplementedHotelServiceServer) UpdateRoom(context.Context, *Room) (*RoomRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedHotelServiceServer) DeleteRoom(context.Context, *RoomResponse) (*RoomRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_CreateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CreateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CreateHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CreateHotel(ctx, req.(*HotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetbyIdHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetbyIdHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetbyIdHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetbyIdHotel(ctx, req.(*HotelResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetAllHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetAllHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetAllHotels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetAllHotels(ctx, req.(*HotelEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hotel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_UpdateHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateHotel(ctx, req.(*Hotel))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_DeleteHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteHotel(ctx, req.(*HotelResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CreateRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetbyIdRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetbyIdRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetbyIdRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetbyIdRoom(ctx, req.(*RoomResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetAllRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetAllRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetAllRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetAllRooms(ctx, req.(*HotelEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateRoom(ctx, req.(*Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_DeleteRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteRoom(ctx, req.(*RoomResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHotel",
			Handler:    _HotelService_CreateHotel_Handler,
		},
		{
			MethodName: "GetbyIdHotel",
			Handler:    _HotelService_GetbyIdHotel_Handler,
		},
		{
			MethodName: "GetAllHotels",
			Handler:    _HotelService_GetAllHotels_Handler,
		},
		{
			MethodName: "UpdateHotel",
			Handler:    _HotelService_UpdateHotel_Handler,
		},
		{
			MethodName: "DeleteHotel",
			Handler:    _HotelService_DeleteHotel_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _HotelService_CreateRoom_Handler,
		},
		{
			MethodName: "GetbyIdRoom",
			Handler:    _HotelService_GetbyIdRoom_Handler,
		},
		{
			MethodName: "GetAllRooms",
			Handler:    _HotelService_GetAllRooms_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _HotelService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _HotelService_DeleteRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotelproto/hotel.proto",
}
