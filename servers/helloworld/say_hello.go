package helloworld

import (
	"context"
	"log"
)

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}
