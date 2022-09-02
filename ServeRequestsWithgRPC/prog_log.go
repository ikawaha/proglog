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

// Procedure implements Procedure.
func (s *progLogsrvc) Procedure(ctx context.Context, p *proglog.ProduceRequest) (*proglog.Produceresponse, error) {
	res := &proglog.Produceresponse{}
	s.logger.Print("progLog.Procedure")
	return res, nil
}

// Consume implements Consume.
func (s *progLogsrvc) Consume(ctx context.Context, p *proglog.ConsumeRequest) (*proglog.Consumeresponse, error) {
	res := &proglog.Consumeresponse{}
	s.logger.Print("progLog.Consume")
	return res, nil
}

// ProcedureStream implements ProcedureStream.
func (s *progLogsrvc) ProcedureStream(ctx context.Context, stream proglog.ProcedureStreamServerStream) error {
	s.logger.Print("progLog.ProcedureStream")
	return nil
}

// ConsumeStream implements ConsumeStream.
func (s *progLogsrvc) ConsumeStream(ctx context.Context, stream proglog.ConsumeStreamServerStream) error {
	s.logger.Print("progLog.ConsumeStream")
	return nil
}
