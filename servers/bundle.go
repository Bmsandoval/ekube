package servers

import (
	"github.com/bmsandoval/ekube/configs"
	"github.com/bmsandoval/ekube/servers/helloworld"
	"google.golang.org/grpc"
)

func BundleAll(server *grpc.Server, config *configs.Configuration) {
	helloworld.RegisterServer(server, config)
}