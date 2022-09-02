// Code generated by goa v3.8.4, DO NOT EDIT.
//
// ProgLog gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/ikawaha/proglog/design

package client

import (
	"context"

	prog_logpb "github.com/ikawaha/proglog/gen/grpc/prog_log/pb"
	proglog "github.com/ikawaha/proglog/gen/prog_log"
	proglogviews "github.com/ikawaha/proglog/gen/prog_log/views"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildProcedureFunc builds the remote method to invoke for "ProgLog" service
// "Procedure" endpoint.
func BuildProcedureFunc(grpccli prog_logpb.ProgLogClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Procedure(ctx, reqpb.(*prog_logpb.ProcedureRequest), opts...)
		}
		return grpccli.Procedure(ctx, &prog_logpb.ProcedureRequest{}, opts...)
	}
}

// EncodeProcedureRequest encodes requests sent to ProgLog Procedure endpoint.
func EncodeProcedureRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*proglog.ProduceRequest)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ProgLog", "Procedure", "*proglog.ProduceRequest", v)
	}
	return NewProtoProcedureRequest(payload), nil
}

// DecodeProcedureResponse decodes responses from the ProgLog Procedure
// endpoint.
func DecodeProcedureResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*prog_logpb.ProcedureResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ProgLog", "Procedure", "*prog_logpb.ProcedureResponse", v)
	}
	res := NewProcedureResult(message)
	vres := &proglogviews.Produceresponse{Projected: res, View: view}
	if err := proglogviews.ValidateProduceresponse(vres); err != nil {
		return nil, err
	}
	return proglog.NewProduceresponse(vres), nil
}

// BuildConsumeFunc builds the remote method to invoke for "ProgLog" service
// "Consume" endpoint.
func BuildConsumeFunc(grpccli prog_logpb.ProgLogClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Consume(ctx, reqpb.(*prog_logpb.ConsumeRequest), opts...)
		}
		return grpccli.Consume(ctx, &prog_logpb.ConsumeRequest{}, opts...)
	}
}

// EncodeConsumeRequest encodes requests sent to ProgLog Consume endpoint.
func EncodeConsumeRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*proglog.ConsumeRequest)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ProgLog", "Consume", "*proglog.ConsumeRequest", v)
	}
	return NewProtoConsumeRequest(payload), nil
}

// DecodeConsumeResponse decodes responses from the ProgLog Consume endpoint.
func DecodeConsumeResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*prog_logpb.ConsumeResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ProgLog", "Consume", "*prog_logpb.ConsumeResponse", v)
	}
	res := NewConsumeResult(message)
	vres := &proglogviews.Consumeresponse{Projected: res, View: view}
	if err := proglogviews.ValidateConsumeresponse(vres); err != nil {
		return nil, err
	}
	return proglog.NewConsumeresponse(vres), nil
}

// BuildProcedureStreamFunc builds the remote method to invoke for "ProgLog"
// service "ProcedureStream" endpoint.
func BuildProcedureStreamFunc(grpccli prog_logpb.ProgLogClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.ProcedureStream(ctx, opts...)
		}
		return grpccli.ProcedureStream(ctx, opts...)
	}
}

// DecodeProcedureStreamResponse decodes responses from the ProgLog
// ProcedureStream endpoint.
func DecodeProcedureStreamResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	return &ProcedureStreamClientStream{
		stream: v.(prog_logpb.ProgLog_ProcedureStreamClient),
		view:   view,
	}, nil
}

// BuildConsumeStreamFunc builds the remote method to invoke for "ProgLog"
// service "ConsumeStream" endpoint.
func BuildConsumeStreamFunc(grpccli prog_logpb.ProgLogClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.ConsumeStream(ctx, opts...)
		}
		return grpccli.ConsumeStream(ctx, opts...)
	}
}

// DecodeConsumeStreamResponse decodes responses from the ProgLog ConsumeStream
// endpoint.
func DecodeConsumeStreamResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	return &ConsumeStreamClientStream{
		stream: v.(prog_logpb.ProgLog_ConsumeStreamClient),
		view:   view,
	}, nil
}
