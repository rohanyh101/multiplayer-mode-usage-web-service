package api

import (
	"fmt"
	"log"
	"net"

	"github.com/roohanyh/lila_p1/cache"
	"github.com/roohanyh/lila_p1/config"
	"github.com/roohanyh/lila_p1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer() {
	cache.Init()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Env.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMultiplayerServiceServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	fmt.Printf("gRPC server is running on port :%s\n", config.Env.PORT)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
