package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "main/hello"
)

var port = 50051


type HelloService struct {
	pb.UnimplementedHelloServiceServer
}
func (h *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Req: &pb.Request{
			Name:    request.Name,
			Message: request.Message,
			IsTrue:  true,
		},
	}, nil
}

func main() {
	lisn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}
	log.Printf("listen: %d\n", port)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})
	grpcServer.Serve(lisn)
}