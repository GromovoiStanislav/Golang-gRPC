package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc_example/profile"
)

func main() {
	fmt.Println("Go gRPC!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	profile.RegisterProfileServiceServer(s, &profile.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}