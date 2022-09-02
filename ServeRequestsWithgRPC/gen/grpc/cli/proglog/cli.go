// Code generated by goa v3.8.4, DO NOT EDIT.
//
// proglog gRPC client CLI support package
//
// Command:
// $ goa gen github.com/ikawaha/proglog/design

package cli

import (
	"flag"
	"fmt"
	"os"

	proglogc "github.com/ikawaha/proglog/gen/grpc/prog_log/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `prog-log (procedure|consume|procedure-stream|consume-stream)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` prog-log procedure --message '{
      "record": {
         "offset": 10189682183573662085,
         "value": "Vm9sdXB0YXRlIHF1YWVyYXQgcXVpIGRvbG9yIGxhdWRhbnRpdW0gbWFpb3Jlcy4="
      }
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		progLogFlags = flag.NewFlagSet("prog-log", flag.ContinueOnError)

		progLogProcedureFlags       = flag.NewFlagSet("procedure", flag.ExitOnError)
		progLogProcedureMessageFlag = progLogProcedureFlags.String("message", "", "")

		progLogConsumeFlags       = flag.NewFlagSet("consume", flag.ExitOnError)
		progLogConsumeMessageFlag = progLogConsumeFlags.String("message", "", "")

		progLogProcedureStreamFlags = flag.NewFlagSet("procedure-stream", flag.ExitOnError)

		progLogConsumeStreamFlags = flag.NewFlagSet("consume-stream", flag.ExitOnError)
	)
	progLogFlags.Usage = progLogUsage
	progLogProcedureFlags.Usage = progLogProcedureUsage
	progLogConsumeFlags.Usage = progLogConsumeUsage
	progLogProcedureStreamFlags.Usage = progLogProcedureStreamUsage
	progLogConsumeStreamFlags.Usage = progLogConsumeStreamUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "prog-log":
			svcf = progLogFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "prog-log":
			switch epn {
			case "procedure":
				epf = progLogProcedureFlags

			case "consume":
				epf = progLogConsumeFlags

			case "procedure-stream":
				epf = progLogProcedureStreamFlags

			case "consume-stream":
				epf = progLogConsumeStreamFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "prog-log":
			c := proglogc.NewClient(cc, opts...)
			switch epn {
			case "procedure":
				endpoint = c.Procedure()
				data, err = proglogc.BuildProcedurePayload(*progLogProcedureMessageFlag)
			case "consume":
				endpoint = c.Consume()
				data, err = proglogc.BuildConsumePayload(*progLogConsumeMessageFlag)
			case "procedure-stream":
				endpoint = c.ProcedureStream()
				data = nil
			case "consume-stream":
				endpoint = c.ConsumeStream()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// prog-logUsage displays the usage of the prog-log command and its subcommands.
func progLogUsage() {
	fmt.Fprintf(os.Stderr, `Service is the ProgLog service interface.
Usage:
    %[1]s [globalflags] prog-log COMMAND [flags]

COMMAND:
    procedure: Procedure implements Procedure.
    consume: Consume implements Consume.
    procedure-stream: ProcedureStream implements ProcedureStream.
    consume-stream: ConsumeStream implements ConsumeStream.

Additional help:
    %[1]s prog-log COMMAND --help
`, os.Args[0])
}
func progLogProcedureUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] prog-log procedure -message JSON

Procedure implements Procedure.
    -message JSON: 

Example:
    %[1]s prog-log procedure --message '{
      "record": {
         "offset": 10189682183573662085,
         "value": "Vm9sdXB0YXRlIHF1YWVyYXQgcXVpIGRvbG9yIGxhdWRhbnRpdW0gbWFpb3Jlcy4="
      }
   }'
`, os.Args[0])
}

func progLogConsumeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] prog-log consume -message JSON

Consume implements Consume.
    -message JSON: 

Example:
    %[1]s prog-log consume --message '{
      "offset": 8097310025797036250
   }'
`, os.Args[0])
}

func progLogProcedureStreamUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] prog-log procedure-stream

ProcedureStream implements ProcedureStream.

Example:
    %[1]s prog-log procedure-stream
`, os.Args[0])
}

func progLogConsumeStreamUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] prog-log consume-stream

ConsumeStream implements ConsumeStream.

Example:
    %[1]s prog-log consume-stream
`, os.Args[0])
}
