package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"



	pb "main/protos"
)

type server struct{
	pb.UnimplementedHelloServiceServer
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: "Hello, " + in.Name}, nil
}


func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})



		fmt.Println("Server is running on port :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}