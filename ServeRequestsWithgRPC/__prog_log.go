package proglogapi

import (
	"context"
	"log"
	proglog "proglog/gen/prog_log"
)

// ProgLog service example implementation.
// The example methods log the requests and return zero values.
type progLogsrvc struct {
	logger *log.Logger
}

// NewProgLog returns the ProgLog service implementation.
func NewProgLog(logger *log.Logger) proglog.Service {
	return &progLogsrvc{logger}
}

// Produce implements Produce.
func (s *progLogsrvc) Produce(ctx context.Context, p *proglog.ProduceRequest) (res *proglog.Produceresponse, err error) {
	res = &proglog.Produceresponse{}
	s.logger.Print("progLog.Produce")
	return
}

// Consume implements Consume.
func (s *progLogsrvc) Consume(ctx context.Context, p *proglog.ConsumeRequest) (res *proglog.Consumeresponse, err error) {
	res = &proglog.Consumeresponse{}
	s.logger.Print("progLog.Consume")
	return
}

// ProduceStream implements ProduceStream.
func (s *progLogsrvc) ProduceStream(ctx context.Context, stream proglog.ProduceStreamServerStream) (err error) {
	s.logger.Print("progLog.ProduceStream")
	return
}

// ConsumeStream implements ConsumeStream.
func (s *progLogsrvc) ConsumeStream(ctx context.Context, p *proglog.ConsumeRequest, stream proglog.ConsumeStreamServerStream) (err error) {
	s.logger.Print("progLog.ConsumeStream")
	return
}
