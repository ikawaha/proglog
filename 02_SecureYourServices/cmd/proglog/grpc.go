package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"

	"google.golang.org/grpc/credentials"
	proglogapi "proglog"
	"proglog/config"
	prog_logpb "proglog/gen/grpc/prog_log/pb"
	proglogsvr "proglog/gen/grpc/prog_log/server"
	proglog "proglog/gen/prog_log"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, progLogEndpoints *proglog.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		progLogServer *proglogsvr.Server
	)
	{
		progLogServer = proglogsvr.New(progLogEndpoints, nil, nil)
	}

	// Initialize gRPC server with the middleware.
	logger.Printf("gRPC server listening on %q", u.Host)
	lis, err := net.Listen("tcp", u.Host)
	if err != nil {
		errc <- err
	}
	serverTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.ServerCertFile,
		KeyFile:       config.ServerKeyFile,
		CAFile:        config.CAFile,
		ServerAddress: lis.Addr().String(),
		Server:        true,
	})
	if err != nil {
		errc <- err
	}
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
			grpcauth.UnaryServerInterceptor(proglogapi.Authenticate),
		),
		grpcmiddleware.WithStreamServerChain(
			grpcmdlwr.StreamRequestID(),
			grpcmdlwr.StreamServerLog(adapter),
			grpcauth.StreamServerInterceptor(proglogapi.Authenticate),
		),
		grpc.Creds(credentials.NewTLS(serverTLSConfig)),
	)

	// Register the servers.
	prog_logpb.RegisterProgLogServer(srv, progLogServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
