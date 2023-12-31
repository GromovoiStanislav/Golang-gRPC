package main

import (
	"fmt"
	"log"
	"net"
	
	"google.golang.org/grpc"
	
	"grpc-go-example/hash"
)

func main() {
	fmt.Println("Go gRPC hash server")

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("Failed to listen on port 6000: %v", err)
	}

	s := hash.Server{}

	grpcServer := grpc.NewServer()

	hash.RegisterHashServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 6000: %v", err)
	}
}