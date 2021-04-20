package main

import (
	"context"
	pingpong "github.com/memeoAmazonas/grpc-react/pingpong"
)

type Server struct {
	pingpong.UnimplementedPingPongServer
}

// Ping fullfills the requirement for PingPong Server interface
func (s *Server) Ping(ctx context.Context, ping *pingpong.PingRequest) (*pingpong.PongResponse, error) {
	return &pingpong.PongResponse{
		Ok: true,
	}, nil
}
