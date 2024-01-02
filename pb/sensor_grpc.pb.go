// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: sensor.proto

package sensors

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
	SensorServices_InsertSensor_FullMethodName = "/sensors.SensorServices/InsertSensor"
	SensorServices_GetSensors_FullMethodName   = "/sensors.SensorServices/GetSensors"
	SensorServices_EditSensor_FullMethodName   = "/sensors.SensorServices/EditSensor"
	SensorServices_DeleteSensor_FullMethodName = "/sensors.SensorServices/DeleteSensor"
)

// SensorServicesClient is the client API for SensorServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensorServicesClient interface {
	InsertSensor(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*MessageReturn, error)
	GetSensors(ctx context.Context, in *PayloadGetSensors, opts ...grpc.CallOption) (*Sensors, error)
	EditSensor(ctx context.Context, in *PayloadUpdateSensors, opts ...grpc.CallOption) (*MessageReturn, error)
	DeleteSensor(ctx context.Context, in *PayloadGetSensors, opts ...grpc.CallOption) (*MessageReturn, error)
}

type sensorServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorServicesClient(cc grpc.ClientConnInterface) SensorServicesClient {
	return &sensorServicesClient{cc}
}

func (c *sensorServicesClient) InsertSensor(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*MessageReturn, error) {
	out := new(MessageReturn)
	err := c.cc.Invoke(ctx, SensorServices_InsertSensor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServicesClient) GetSensors(ctx context.Context, in *PayloadGetSensors, opts ...grpc.CallOption) (*Sensors, error) {
	out := new(Sensors)
	err := c.cc.Invoke(ctx, SensorServices_GetSensors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServicesClient) EditSensor(ctx context.Context, in *PayloadUpdateSensors, opts ...grpc.CallOption) (*MessageReturn, error) {
	out := new(MessageReturn)
	err := c.cc.Invoke(ctx, SensorServices_EditSensor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServicesClient) DeleteSensor(ctx context.Context, in *PayloadGetSensors, opts ...grpc.CallOption) (*MessageReturn, error) {
	out := new(MessageReturn)
	err := c.cc.Invoke(ctx, SensorServices_DeleteSensor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorServicesServer is the server API for SensorServices service.
// All implementations should embed UnimplementedSensorServicesServer
// for forward compatibility
type SensorServicesServer interface {
	InsertSensor(context.Context, *PayloadRequest) (*MessageReturn, error)
	GetSensors(context.Context, *PayloadGetSensors) (*Sensors, error)
	EditSensor(context.Context, *PayloadUpdateSensors) (*MessageReturn, error)
	DeleteSensor(context.Context, *PayloadGetSensors) (*MessageReturn, error)
}

// UnimplementedSensorServicesServer should be embedded to have forward compatible implementations.
type UnimplementedSensorServicesServer struct {
}

func (UnimplementedSensorServicesServer) InsertSensor(context.Context, *PayloadRequest) (*MessageReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertSensor not implemented")
}
func (UnimplementedSensorServicesServer) GetSensors(context.Context, *PayloadGetSensors) (*Sensors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensors not implemented")
}
func (UnimplementedSensorServicesServer) EditSensor(context.Context, *PayloadUpdateSensors) (*MessageReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSensor not implemented")
}
func (UnimplementedSensorServicesServer) DeleteSensor(context.Context, *PayloadGetSensors) (*MessageReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSensor not implemented")
}

// UnsafeSensorServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensorServicesServer will
// result in compilation errors.
type UnsafeSensorServicesServer interface {
	mustEmbedUnimplementedSensorServicesServer()
}

func RegisterSensorServicesServer(s grpc.ServiceRegistrar, srv SensorServicesServer) {
	s.RegisterService(&SensorServices_ServiceDesc, srv)
}

func _SensorServices_InsertSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServicesServer).InsertSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensorServices_InsertSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServicesServer).InsertSensor(ctx, req.(*PayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorServices_GetSensors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadGetSensors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServicesServer).GetSensors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensorServices_GetSensors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServicesServer).GetSensors(ctx, req.(*PayloadGetSensors))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorServices_EditSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadUpdateSensors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServicesServer).EditSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensorServices_EditSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServicesServer).EditSensor(ctx, req.(*PayloadUpdateSensors))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorServices_DeleteSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadGetSensors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServicesServer).DeleteSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SensorServices_DeleteSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServicesServer).DeleteSensor(ctx, req.(*PayloadGetSensors))
	}
	return interceptor(ctx, in, info, handler)
}

// SensorServices_ServiceDesc is the grpc.ServiceDesc for SensorServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SensorServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensors.SensorServices",
	HandlerType: (*SensorServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertSensor",
			Handler:    _SensorServices_InsertSensor_Handler,
		},
		{
			MethodName: "GetSensors",
			Handler:    _SensorServices_GetSensors_Handler,
		},
		{
			MethodName: "EditSensor",
			Handler:    _SensorServices_EditSensor_Handler,
		},
		{
			MethodName: "DeleteSensor",
			Handler:    _SensorServices_DeleteSensor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sensor.proto",
}