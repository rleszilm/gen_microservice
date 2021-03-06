// Package greeter is generated by protoc-gen-go-genms. *DO NOT EDIT*
package greeter

import (
	context "context"

	service "github.com/rleszilm/genms/service"
	graphql "github.com/rleszilm/genms/service/graphql"
	grpc "github.com/rleszilm/genms/service/grpc"
	grpc1 "google.golang.org/grpc"
)

// WithGraphQLServerService implements WithGraphQLService
type WithGraphQLServerService struct {
	service.Dependencies

	impl       WithGraphQLServer
	grpcServer *grpc.Server

	graphqlServer *graphql.Server
	proxy         *grpc.Proxy
}

// Initialize implements service.Service.Initialize
func (s *WithGraphQLServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc1.Server) {
		RegisterWithGraphQLServer(server, s.impl)
	})

	if err := s.graphqlServer.WithGrpcProxy(ctx, s.proxy, RegisterWithGraphQLGraphqlWithOptions); err != nil {
		return err
	}
	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *WithGraphQLServerService) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of the service
func (s *WithGraphQLServerService) NameOf() string {
	return "with-graph-ql"
}

// String returns the string name of the service
func (s *WithGraphQLServerService) String() string {
	return s.NameOf()
}

// NewWithGraphQLServerService returns a new WithGraphQLServerService
func NewWithGraphQLServerService(impl WithGraphQLServer, grpcServer *grpc.Server, graphqlServer *graphql.Server, proxy *grpc.Proxy) *WithGraphQLServerService {
	server := &WithGraphQLServerService{
		impl:          impl,
		grpcServer:    grpcServer,
		graphqlServer: graphqlServer,

		proxy: proxy,
	}

	if asService, ok := impl.(service.Service); ok {
		server.WithDependencies(asService)
	}

	grpcServer.WithDependencies(server)

	graphqlServer.WithDependencies(server)
	return server
}
