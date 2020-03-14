package entry

import (
	"fmt"
	"github.com/bmsandoval/ekube/configs"
	"github.com/bmsandoval/ekube/db"
	"github.com/bmsandoval/ekube/library/appcontext"
	"github.com/bmsandoval/ekube/servers"
	"github.com/bmsandoval/ekube/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Entry() {
	// Get Configs
	config, err := configs.Configure()
	if err != nil {
		panic(err) }

	// Setup Database
	connection, err := db.Start(*config);
	if err != nil {
		panic(err) }
	defer func() {
		if err := db.Stop(); err != nil {
			panic(err) }
	}()

	// Build Context
	ctx := appcontext.Context{
		Config: *config,
		DB: *connection,
		// Redis
	}

	// Bundle Services
	serviceBundle, err := services.NewBundle(ctx)
	if err != nil {
		panic(err) }

	// Bundle Servers
	s := grpc.NewServer()
	servers.BundleAll(s, ctx, *serviceBundle)

	// Start Server
	log.Println("Starting Server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.SrvPort))
	if err != nil {
		panic(err) }
	if err := s.Serve(lis); err != nil {
		panic(err) }
}
