// Code generated by goa v3.8.4, DO NOT EDIT.
//
// ProgLog client
//
// Command:
// $ goa gen proglog/design

package proglog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ProgLog" service client.
type Client struct {
	ProduceEndpoint       goa.Endpoint
	ConsumeEndpoint       goa.Endpoint
	ProduceStreamEndpoint goa.Endpoint
	ConsumeStreamEndpoint goa.Endpoint
}

// NewClient initializes a "ProgLog" service client given the endpoints.
func NewClient(produce, consume, produceStream, consumeStream goa.Endpoint) *Client {
	return &Client{
		ProduceEndpoint:       produce,
		ConsumeEndpoint:       consume,
		ProduceStreamEndpoint: produceStream,
		ConsumeStreamEndpoint: consumeStream,
	}
}

// Produce calls the "Produce" endpoint of the "ProgLog" service.
func (c *Client) Produce(ctx context.Context, p *ProduceRequest) (res *Produceresponse, err error) {
	var ires interface{}
	ires, err = c.ProduceEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Produceresponse), nil
}

// Consume calls the "Consume" endpoint of the "ProgLog" service.
func (c *Client) Consume(ctx context.Context, p *ConsumeRequest) (res *Consumeresponse, err error) {
	var ires interface{}
	ires, err = c.ConsumeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Consumeresponse), nil
}

// ProduceStream calls the "ProduceStream" endpoint of the "ProgLog" service.
func (c *Client) ProduceStream(ctx context.Context) (res ProduceStreamClientStream, err error) {
	var ires interface{}
	ires, err = c.ProduceStreamEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(ProduceStreamClientStream), nil
}

// ConsumeStream calls the "ConsumeStream" endpoint of the "ProgLog" service.
// ConsumeStream may return the following errors:
//   - "OffsetOutOfRange" (type OffsetOutOfRange)
//   - error: internal error
func (c *Client) ConsumeStream(ctx context.Context, p *ConsumeRequest) (res ConsumeStreamClientStream, err error) {
	var ires interface{}
	ires, err = c.ConsumeStreamEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(ConsumeStreamClientStream), nil
}
