// Code generated by goa v3.8.4, DO NOT EDIT.
//
// ProgLog gRPC client CLI support package
//
// Command:
// $ goa gen proglog/design

package client

import (
	"encoding/json"
	"fmt"
	prog_logpb "proglog/gen/grpc/prog_log/pb"
	proglog "proglog/gen/prog_log"
)

// BuildProducePayload builds the payload for the ProgLog Produce endpoint from
// CLI flags.
func BuildProducePayload(progLogProduceMessage string) (*proglog.ProduceRequest, error) {
	var err error
	var message prog_logpb.ProduceRequest
	{
		if progLogProduceMessage != "" {
			err = json.Unmarshal([]byte(progLogProduceMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"record\": {\n         \"offset\": 10189682183573662085,\n         \"value\": \"Vm9sdXB0YXRlIHF1YWVyYXQgcXVpIGRvbG9yIGxhdWRhbnRpdW0gbWFpb3Jlcy4=\"\n      }\n   }'")
			}
		}
	}
	v := &proglog.ProduceRequest{}
	if message.Record != nil {
		v.Record = protobufProgLogpbRecordToProglogRecord(message.Record)
	}

	return v, nil
}

// BuildConsumePayload builds the payload for the ProgLog Consume endpoint from
// CLI flags.
func BuildConsumePayload(progLogConsumeMessage string) (*proglog.ConsumeRequest, error) {
	var err error
	var message prog_logpb.ConsumeRequest
	{
		if progLogConsumeMessage != "" {
			err = json.Unmarshal([]byte(progLogConsumeMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"offset\": 8097310025797036250\n   }'")
			}
		}
	}
	v := &proglog.ConsumeRequest{
		Offset: message.Offset,
	}

	return v, nil
}

// BuildConsumeStreamPayload builds the payload for the ProgLog ConsumeStream
// endpoint from CLI flags.
func BuildConsumeStreamPayload(progLogConsumeStreamMessage string) (*proglog.ConsumeRequest, error) {
	var err error
	var message prog_logpb.ConsumeStreamRequest
	{
		if progLogConsumeStreamMessage != "" {
			err = json.Unmarshal([]byte(progLogConsumeStreamMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"offset\": 17539627402234799792\n   }'")
			}
		}
	}
	v := &proglog.ConsumeRequest{
		Offset: message.Offset,
	}

	return v, nil
}
