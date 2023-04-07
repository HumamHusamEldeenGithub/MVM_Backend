package core

import (
	"context"
	"fmt"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

// IService is the interface for all services
type IService interface {
	Register(srv grpc.ServiceRegistrar)
}

type application struct {
	opts []grpc.ServerOption

	grpcPort   uint32
	grpc       net.Listener
	grpcServer *grpc.Server
}

// New creates core application
func New(ctx context.Context, grpcPort uint32) *application {
	a := &application{
		grpcPort: grpcPort,
	}
	a.opts = make([]grpc.ServerOption, 0)

	return a
}

// Run starts core application
func (a *application) Run(services ...IService) error {

	a.listenGRPC(a.grpcPort)

	a.grpcServer = a.initGRPC()
	for _, service := range services {
		service.Register(a.grpcServer)

	}

	err := a.grpcServer.Serve(a.grpc)
	if err != nil {
		panic(err)
	}

	return nil
}

func (a *application) listenGRPC(port uint32) {
	grpc, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	a.grpc = grpc
}

func (a *application) initGRPC() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	}
	opts = append(opts, a.opts...)

	grpcServer := grpc.NewServer(
		opts...,
	)
	grpc_prometheus.Register(grpcServer)

	return grpcServer
}

// WithUnaryMW adds unary interceptor
func (a *application) WithUnaryMW(unary grpc.UnaryServerInterceptor) {
	a.opts = append(a.opts, grpc.ChainUnaryInterceptor(unary))
}

func (a *application) Close() {
	a.grpcServer.GracefulStop()
}
