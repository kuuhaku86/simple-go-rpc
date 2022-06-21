package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/kuuhaku86/simple-go-rpc/dataserver"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8600")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDataServerServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}