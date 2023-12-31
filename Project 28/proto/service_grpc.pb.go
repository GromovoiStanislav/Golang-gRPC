// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/service.proto

package hardwaremonitoring

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
	HardwareMonitor_Monitor_FullMethodName = "/hardwaremonitoring.HardwareMonitor/Monitor"
)

// HardwareMonitorClient is the client API for HardwareMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HardwareMonitorClient interface {
	Monitor(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (HardwareMonitor_MonitorClient, error)
}

type hardwareMonitorClient struct {
	cc grpc.ClientConnInterface
}

func NewHardwareMonitorClient(cc grpc.ClientConnInterface) HardwareMonitorClient {
	return &hardwareMonitorClient{cc}
}

func (c *hardwareMonitorClient) Monitor(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (HardwareMonitor_MonitorClient, error) {
	stream, err := c.cc.NewStream(ctx, &HardwareMonitor_ServiceDesc.Streams[0], HardwareMonitor_Monitor_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hardwareMonitorMonitorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HardwareMonitor_MonitorClient interface {
	Recv() (*HardwareStats, error)
	grpc.ClientStream
}

type hardwareMonitorMonitorClient struct {
	grpc.ClientStream
}

func (x *hardwareMonitorMonitorClient) Recv() (*HardwareStats, error) {
	m := new(HardwareStats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HardwareMonitorServer is the server API for HardwareMonitor service.
// All implementations must embed UnimplementedHardwareMonitorServer
// for forward compatibility
type HardwareMonitorServer interface {
	Monitor(*EmptyRequest, HardwareMonitor_MonitorServer) error
	mustEmbedUnimplementedHardwareMonitorServer()
}

// UnimplementedHardwareMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedHardwareMonitorServer struct {
}

func (UnimplementedHardwareMonitorServer) Monitor(*EmptyRequest, HardwareMonitor_MonitorServer) error {
	return status.Errorf(codes.Unimplemented, "method Monitor not implemented")
}
func (UnimplementedHardwareMonitorServer) mustEmbedUnimplementedHardwareMonitorServer() {}

// UnsafeHardwareMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HardwareMonitorServer will
// result in compilation errors.
type UnsafeHardwareMonitorServer interface {
	mustEmbedUnimplementedHardwareMonitorServer()
}

func RegisterHardwareMonitorServer(s grpc.ServiceRegistrar, srv HardwareMonitorServer) {
	s.RegisterService(&HardwareMonitor_ServiceDesc, srv)
}

func _HardwareMonitor_Monitor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HardwareMonitorServer).Monitor(m, &hardwareMonitorMonitorServer{stream})
}

type HardwareMonitor_MonitorServer interface {
	Send(*HardwareStats) error
	grpc.ServerStream
}

type hardwareMonitorMonitorServer struct {
	grpc.ServerStream
}

func (x *hardwareMonitorMonitorServer) Send(m *HardwareStats) error {
	return x.ServerStream.SendMsg(m)
}

// HardwareMonitor_ServiceDesc is the grpc.ServiceDesc for HardwareMonitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HardwareMonitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hardwaremonitoring.HardwareMonitor",
	HandlerType: (*HardwareMonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Monitor",
			Handler:       _HardwareMonitor_Monitor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
