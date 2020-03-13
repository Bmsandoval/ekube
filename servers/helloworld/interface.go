package helloworld

import (
	"github.com/bmsandoval/ekube/servers"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	servers.ServerContext
}

//type Registerable struct {}
func init() {
	servers.Include(func(s *grpc.Server, ctx servers.ServerContext){
		RegisterGreeterServer(s, &Server{
			ServerContext: ctx,
		})
	})
}
