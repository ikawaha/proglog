package design

import (
	. "goa.design/goa/v3/dsl"
)

var Record = Type("Record", func() {
	Field(1, "value", Bytes)
	Field(2, "offset", UInt64)
	Required("value", "offset")
})

var ProduceRequest = Type("ProduceRequest", func() {
	Field(1, "record", Record)
	Required("record")
})

var ProduceResponse = ResultType("ProduceResponse", func() {
	Field(1, "offset", UInt64)
	Required("offset")
})

var ConsumeRequest = Type("ConsumeRequest", func() {
	Field(1, "offset", UInt64)
	Required("offset")
})

var ConsumeResponse = ResultType("ConsumeResponse", func() {
	Field(1, "record", Record)
	Required("record")
})

var _ = Service("ProgLog", func() {
	Method("Procedure", func() {
		Payload(ProduceRequest)
		Result(ProduceResponse)
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("Consume", func() {
		Payload(ConsumeRequest)
		Result(ConsumeResponse)
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("ProcedureStream", func() {
		StreamingPayload(ProduceRequest)
		Result(ProduceResponse)
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("ConsumeStream", func() {
		StreamingPayload(ConsumeRequest)
		StreamingResult(ConsumeResponse)
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
