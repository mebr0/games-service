package server

import (
	"context"
	"fmt"
	"github.com/mebr0/grpc-server/internal/config"
	pb "github.com/mebr0/grpc-server/pkg/api/v1"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(gameService pb.GameServiceServer) *Server {
	grpcServer := grpc.NewServer()

	pb.RegisterGameServiceServer(grpcServer, gameService)

	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Start(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))

	if err != nil {
		return err
	}

	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	// timeout?
	s.grpcServer.GracefulStop()

	return nil
}
