package proglogapi

import (
	"context"
	"log"

	pb "proglog/gen/grpc/prog_log/pb"
	proglog "proglog/gen/prog_log"
)

const (
	objectWildcard = "*"
	produceAction  = "produce"
	consumeAction  = "consume"
)

// ProgLog service example implementation.
// The example methods log the requests and return zero values.
type progLogsrvc struct {
	logger *log.Logger
	*Config
}

// NewProgLog returns the ProgLog service implementation.
func NewProgLog(logger *log.Logger, config *Config) proglog.Service {
	return &progLogsrvc{logger: logger, Config: config}
}

type CommitLog interface {
	Append(*pb.Record) (uint64, error)
	Read(uint64) (*pb.Record, error)
}

type Config struct {
	CommitLog  CommitLog
	Authorizer *Authorizer
}

// Produce implements Produce.
func (s *progLogsrvc) Produce(ctx context.Context, p *proglog.ProduceRequest) (*proglog.Produceresponse, error) {
	if err := s.Authorizer.Authorize(subject(ctx), objectWildcard, produceAction); err != nil {
		return nil, err
	}
	offset, err := s.CommitLog.Append(&pb.Record{
		Value:  p.Record.Value,
		Offset: p.Record.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &proglog.Produceresponse{Offset: offset}, nil
}

// Consume implements Consume.
func (s *progLogsrvc) Consume(ctx context.Context, p *proglog.ConsumeRequest) (*proglog.Consumeresponse, error) {
	if err := s.Authorizer.Authorize(subject(ctx), objectWildcard, consumeAction); err != nil {
		return nil, err
	}
	record, err := s.CommitLog.Read(p.Offset)
	if err != nil {
		return nil, err
	}
	return &proglog.Consumeresponse{Record: &proglog.Record{
		Value:  record.Value,
		Offset: record.Offset,
	}}, nil
}

func (s *progLogsrvc) ProduceStream(ctx context.Context, stream proglog.ProduceStreamServerStream) (err error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		res, err := s.Produce(ctx, req)
		if err != nil {
			return err
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

// ConsumeStream implements ConsumeStream.
func (s *progLogsrvc) ConsumeStream(ctx context.Context, p *proglog.ConsumeRequest, stream proglog.ConsumeStreamServerStream) (err error) {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			res, err := s.Consume(ctx, p)
			switch err.(type) {
			case nil:
			case proglog.OffsetOutOfRange:
				continue
			default:
				return err
			}
			if err = stream.Send(res); err != nil {
				return err
			}
			p.Offset++
		}
	}
}
