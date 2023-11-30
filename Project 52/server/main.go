package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-example/gravatar"
)

const port = ":50051"


type Server struct {
	pb.GravatarServiceServer
}


func (s *Server) Generate(ctx context.Context, in *pb.GravatarRequest) (*pb.GravatarResponse, error) {
	log.Printf("Received email %v with size %v", in.Email, in.Size)
	return &pb.GravatarResponse{Url: gravatar(in.Email, in.Size)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGravatarServiceServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to start server!", err)
	}
}