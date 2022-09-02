// Code generated by goa v3.8.4, DO NOT EDIT.
//
// ProgLog gRPC client types
//
// Command:
// $ goa gen proglog/design

package client

import (
	prog_logpb "proglog/gen/grpc/prog_log/pb"
	proglog "proglog/gen/prog_log"
	proglogviews "proglog/gen/prog_log/views"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoProcedureRequest builds the gRPC request type from the payload of
// the "Procedure" endpoint of the "ProgLog" service.
func NewProtoProcedureRequest(payload *proglog.ProduceRequest) *prog_logpb.ProcedureRequest {
	message := &prog_logpb.ProcedureRequest{}
	if payload.Record != nil {
		message.Record = svcProglogRecordToProgLogpbRecord(payload.Record)
	}
	return message
}

// NewProcedureResult builds the result type of the "Procedure" endpoint of the
// "ProgLog" service from the gRPC response type.
func NewProcedureResult(message *prog_logpb.ProcedureResponse) *proglogviews.ProduceresponseView {
	result := &proglogviews.ProduceresponseView{
		Offset: &message.Offset,
	}
	return result
}

// NewProtoConsumeRequest builds the gRPC request type from the payload of the
// "Consume" endpoint of the "ProgLog" service.
func NewProtoConsumeRequest(payload *proglog.ConsumeRequest) *prog_logpb.ConsumeRequest {
	message := &prog_logpb.ConsumeRequest{
		Offset: payload.Offset,
	}
	return message
}

// NewConsumeResult builds the result type of the "Consume" endpoint of the
// "ProgLog" service from the gRPC response type.
func NewConsumeResult(message *prog_logpb.ConsumeResponse) *proglogviews.ConsumeresponseView {
	result := &proglogviews.ConsumeresponseView{}
	if message.Record != nil {
		result.Record = protobufProgLogpbRecordToProglogviewsRecordView(message.Record)
	}
	return result
}

func NewProduceresponseView(v *prog_logpb.ProcedureStreamResponse) *proglogviews.ProduceresponseView {
	vresult := &proglogviews.ProduceresponseView{
		Offset: &v.Offset,
	}
	return vresult
}

func NewProtoProcedureStreamStreamingRequest(spayload *proglog.ProduceRequest) *prog_logpb.ProcedureStreamStreamingRequest {
	v := &prog_logpb.ProcedureStreamStreamingRequest{}
	if spayload.Record != nil {
		v.Record = svcProglogRecordToProgLogpbRecord(spayload.Record)
	}
	return v
}

func NewConsumeresponseView(v *prog_logpb.ConsumeStreamResponse) *proglogviews.ConsumeresponseView {
	vresult := &proglogviews.ConsumeresponseView{}
	if v.Record != nil {
		vresult.Record = protobufProgLogpbRecordToProglogviewsRecordView(v.Record)
	}
	return vresult
}

func NewProtoConsumeStreamStreamingRequest(spayload *proglog.ConsumeRequest) *prog_logpb.ConsumeStreamStreamingRequest {
	v := &prog_logpb.ConsumeStreamStreamingRequest{
		Offset: spayload.Offset,
	}
	return v
}

// ValidateRecord runs the validations defined on Record.
func ValidateRecord(record *prog_logpb.Record) (err error) {
	if record.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "record"))
	}
	return
}

// ValidateConsumeResponse runs the validations defined on ConsumeResponse.
func ValidateConsumeResponse(message *prog_logpb.ConsumeResponse) (err error) {
	if message.Record == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("record", "message"))
	}
	return
}

// ValidateConsumeStreamResponse runs the validations defined on
// ConsumeStreamResponse.
func ValidateConsumeStreamResponse(stream *prog_logpb.ConsumeStreamResponse) (err error) {
	if stream.Record == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("record", "stream"))
	}
	return
}

// protobufProgLogpbRecordToProglogRecord builds a value of type
// *proglog.Record from a value of type *prog_logpb.Record.
func protobufProgLogpbRecordToProglogRecord(v *prog_logpb.Record) *proglog.Record {
	res := &proglog.Record{
		Value:  v.Value,
		Offset: v.Offset,
	}

	return res
}

// svcProglogRecordToProgLogpbRecord builds a value of type *prog_logpb.Record
// from a value of type *proglog.Record.
func svcProglogRecordToProgLogpbRecord(v *proglog.Record) *prog_logpb.Record {
	res := &prog_logpb.Record{
		Value:  v.Value,
		Offset: v.Offset,
	}

	return res
}

// svcProglogviewsRecordViewToProgLogpbRecord builds a value of type
// *prog_logpb.Record from a value of type *proglogviews.RecordView.
func svcProglogviewsRecordViewToProgLogpbRecord(v *proglogviews.RecordView) *prog_logpb.Record {
	res := &prog_logpb.Record{
		Value: v.Value,
	}
	if v.Offset != nil {
		res.Offset = *v.Offset
	}

	return res
}

// protobufProgLogpbRecordToProglogviewsRecordView builds a value of type
// *proglogviews.RecordView from a value of type *prog_logpb.Record.
func protobufProgLogpbRecordToProglogviewsRecordView(v *prog_logpb.Record) *proglogviews.RecordView {
	res := &proglogviews.RecordView{
		Value:  v.Value,
		Offset: &v.Offset,
	}

	return res
}
