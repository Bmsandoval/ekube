package helloworld

import (
	"github.com/bmsandoval/ekube/configs"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	//pb.UnimplementedGreeterServer
	config configs.Configuration
}

func RegisterServer(s *grpc.Server, config *configs.Configuration) {
	RegisterGreeterServer(s, &Server{
		config: *config,
	})
}
