// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/weather.proto

package weather

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
	Weather_Cities_FullMethodName   = "/Weather/cities"
	Weather_Get_FullMethodName      = "/Weather/get"
	Weather_Ping_FullMethodName     = "/Weather/ping"
	Weather_Forecast_FullMethodName = "/Weather/forecast"
)

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherClient interface {
	Cities(ctx context.Context, in *CityQuery, opts ...grpc.CallOption) (*CityQuery_Result, error)
	Get(ctx context.Context, in *GetTemperature, opts ...grpc.CallOption) (Weather_GetClient, error)
	// Unsupported at the moment
	Ping(ctx context.Context, opts ...grpc.CallOption) (Weather_PingClient, error)
	Forecast(ctx context.Context, opts ...grpc.CallOption) (Weather_ForecastClient, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) Cities(ctx context.Context, in *CityQuery, opts ...grpc.CallOption) (*CityQuery_Result, error) {
	out := new(CityQuery_Result)
	err := c.cc.Invoke(ctx, Weather_Cities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherClient) Get(ctx context.Context, in *GetTemperature, opts ...grpc.CallOption) (Weather_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Weather_ServiceDesc.Streams[0], Weather_Get_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Weather_GetClient interface {
	Recv() (*Temperature, error)
	grpc.ClientStream
}

type weatherGetClient struct {
	grpc.ClientStream
}

func (x *weatherGetClient) Recv() (*Temperature, error) {
	m := new(Temperature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherClient) Ping(ctx context.Context, opts ...grpc.CallOption) (Weather_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Weather_ServiceDesc.Streams[1], Weather_Ping_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherPingClient{stream}
	return x, nil
}

type Weather_PingClient interface {
	Send(*Ping) error
	CloseAndRecv() (*Ping_Ack, error)
	grpc.ClientStream
}

type weatherPingClient struct {
	grpc.ClientStream
}

func (x *weatherPingClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *weatherPingClient) CloseAndRecv() (*Ping_Ack, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Ping_Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherClient) Forecast(ctx context.Context, opts ...grpc.CallOption) (Weather_ForecastClient, error) {
	stream, err := c.cc.NewStream(ctx, &Weather_ServiceDesc.Streams[2], Weather_Forecast_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherForecastClient{stream}
	return x, nil
}

type Weather_ForecastClient interface {
	Send(*Forecast) error
	Recv() (*Forecast_Result, error)
	grpc.ClientStream
}

type weatherForecastClient struct {
	grpc.ClientStream
}

func (x *weatherForecastClient) Send(m *Forecast) error {
	return x.ClientStream.SendMsg(m)
}

func (x *weatherForecastClient) Recv() (*Forecast_Result, error) {
	m := new(Forecast_Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServer is the server API for Weather service.
// All implementations must embed UnimplementedWeatherServer
// for forward compatibility
type WeatherServer interface {
	Cities(context.Context, *CityQuery) (*CityQuery_Result, error)
	Get(*GetTemperature, Weather_GetServer) error
	// Unsupported at the moment
	Ping(Weather_PingServer) error
	Forecast(Weather_ForecastServer) error
	mustEmbedUnimplementedWeatherServer()
}

// UnimplementedWeatherServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (UnimplementedWeatherServer) Cities(context.Context, *CityQuery) (*CityQuery_Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cities not implemented")
}
func (UnimplementedWeatherServer) Get(*GetTemperature, Weather_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWeatherServer) Ping(Weather_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWeatherServer) Forecast(Weather_ForecastServer) error {
	return status.Errorf(codes.Unimplemented, "method Forecast not implemented")
}
func (UnimplementedWeatherServer) mustEmbedUnimplementedWeatherServer() {}

// UnsafeWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServer will
// result in compilation errors.
type UnsafeWeatherServer interface {
	mustEmbedUnimplementedWeatherServer()
}

func RegisterWeatherServer(s grpc.ServiceRegistrar, srv WeatherServer) {
	s.RegisterService(&Weather_ServiceDesc, srv)
}

func _Weather_Cities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).Cities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Weather_Cities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).Cities(ctx, req.(*CityQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weather_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTemperature)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServer).Get(m, &weatherGetServer{stream})
}

type Weather_GetServer interface {
	Send(*Temperature) error
	grpc.ServerStream
}

type weatherGetServer struct {
	grpc.ServerStream
}

func (x *weatherGetServer) Send(m *Temperature) error {
	return x.ServerStream.SendMsg(m)
}

func _Weather_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WeatherServer).Ping(&weatherPingServer{stream})
}

type Weather_PingServer interface {
	SendAndClose(*Ping_Ack) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type weatherPingServer struct {
	grpc.ServerStream
}

func (x *weatherPingServer) SendAndClose(m *Ping_Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *weatherPingServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Weather_Forecast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WeatherServer).Forecast(&weatherForecastServer{stream})
}

type Weather_ForecastServer interface {
	Send(*Forecast_Result) error
	Recv() (*Forecast, error)
	grpc.ServerStream
}

type weatherForecastServer struct {
	grpc.ServerStream
}

func (x *weatherForecastServer) Send(m *Forecast_Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *weatherForecastServer) Recv() (*Forecast, error) {
	m := new(Forecast)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Weather_ServiceDesc is the grpc.ServiceDesc for Weather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "cities",
			Handler:    _Weather_Cities_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "get",
			Handler:       _Weather_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ping",
			Handler:       _Weather_Ping_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "forecast",
			Handler:       _Weather_Forecast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/weather.proto",
}