// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: goadesign_goagen_prog_log.proto

package prog_logpb

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

// ProgLogClient is the client API for ProgLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgLogClient interface {
	// Produce implements Produce.
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	// Consume implements Consume.
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	// ProduceStream implements ProduceStream.
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (ProgLog_ProduceStreamClient, error)
	// ConsumeStream implements ConsumeStream.
	ConsumeStream(ctx context.Context, in *ConsumeStreamRequest, opts ...grpc.CallOption) (ProgLog_ConsumeStreamClient, error)
}

type progLogClient struct {
	cc grpc.ClientConnInterface
}

func NewProgLogClient(cc grpc.ClientConnInterface) ProgLogClient {
	return &progLogClient{cc}
}

func (c *progLogClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/prog_log.ProgLog/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progLogClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, "/prog_log.ProgLog/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progLogClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (ProgLog_ProduceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProgLog_ServiceDesc.Streams[0], "/prog_log.ProgLog/ProduceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &progLogProduceStreamClient{stream}
	return x, nil
}

type ProgLog_ProduceStreamClient interface {
	Send(*ProduceStreamStreamingRequest) error
	Recv() (*ProduceStreamResponse, error)
	grpc.ClientStream
}

type progLogProduceStreamClient struct {
	grpc.ClientStream
}

func (x *progLogProduceStreamClient) Send(m *ProduceStreamStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *progLogProduceStreamClient) Recv() (*ProduceStreamResponse, error) {
	m := new(ProduceStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *progLogClient) ConsumeStream(ctx context.Context, in *ConsumeStreamRequest, opts ...grpc.CallOption) (ProgLog_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProgLog_ServiceDesc.Streams[1], "/prog_log.ProgLog/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &progLogConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProgLog_ConsumeStreamClient interface {
	Recv() (*ConsumeStreamResponse, error)
	grpc.ClientStream
}

type progLogConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *progLogConsumeStreamClient) Recv() (*ConsumeStreamResponse, error) {
	m := new(ConsumeStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProgLogServer is the server API for ProgLog service.
// All implementations must embed UnimplementedProgLogServer
// for forward compatibility
type ProgLogServer interface {
	// Produce implements Produce.
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	// Consume implements Consume.
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	// ProduceStream implements ProduceStream.
	ProduceStream(ProgLog_ProduceStreamServer) error
	// ConsumeStream implements ConsumeStream.
	ConsumeStream(*ConsumeStreamRequest, ProgLog_ConsumeStreamServer) error
	mustEmbedUnimplementedProgLogServer()
}

// UnimplementedProgLogServer must be embedded to have forward compatible implementations.
type UnimplementedProgLogServer struct {
}

func (UnimplementedProgLogServer) Produce(context.Context, *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedProgLogServer) Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedProgLogServer) ProduceStream(ProgLog_ProduceStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProduceStream not implemented")
}
func (UnimplementedProgLogServer) ConsumeStream(*ConsumeStreamRequest, ProgLog_ConsumeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
}
func (UnimplementedProgLogServer) mustEmbedUnimplementedProgLogServer() {}

// UnsafeProgLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgLogServer will
// result in compilation errors.
type UnsafeProgLogServer interface {
	mustEmbedUnimplementedProgLogServer()
}

func RegisterProgLogServer(s grpc.ServiceRegistrar, srv ProgLogServer) {
	s.RegisterService(&ProgLog_ServiceDesc, srv)
}

func _ProgLog_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgLogServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prog_log.ProgLog/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgLogServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgLog_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgLogServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prog_log.ProgLog/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgLogServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgLog_ProduceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProgLogServer).ProduceStream(&progLogProduceStreamServer{stream})
}

type ProgLog_ProduceStreamServer interface {
	Send(*ProduceStreamResponse) error
	Recv() (*ProduceStreamStreamingRequest, error)
	grpc.ServerStream
}

type progLogProduceStreamServer struct {
	grpc.ServerStream
}

func (x *progLogProduceStreamServer) Send(m *ProduceStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *progLogProduceStreamServer) Recv() (*ProduceStreamStreamingRequest, error) {
	m := new(ProduceStreamStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProgLog_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProgLogServer).ConsumeStream(m, &progLogConsumeStreamServer{stream})
}

type ProgLog_ConsumeStreamServer interface {
	Send(*ConsumeStreamResponse) error
	grpc.ServerStream
}

type progLogConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *progLogConsumeStreamServer) Send(m *ConsumeStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProgLog_ServiceDesc is the grpc.ServiceDesc for ProgLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prog_log.ProgLog",
	HandlerType: (*ProgLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _ProgLog_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _ProgLog_Consume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProduceStream",
			Handler:       _ProgLog_ProduceStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ConsumeStream",
			Handler:       _ProgLog_ConsumeStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_prog_log.proto",
}
