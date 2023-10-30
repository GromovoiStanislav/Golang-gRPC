package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "main/protos"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	
	// Create trailer in defer to record function return time.
	defer func() {
		trailer := metadata.Pairs("server-trailer","END")
		grpc.SetTrailer(ctx, trailer)
	}()
	
	
	// Получение метаданных от клиента
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println("Received metadata from client:")
		for key, values := range md {
			log.Printf("%s: %v\n", key, values)
		}
	}

	// Отправка метаданных от сервера
	md = metadata.Pairs("server-metadata", "Server Value")
	grpc.SendHeader(ctx, md)

	// Отправка основного ответа от сервера
	greeting := "Hello, " + req.Name
	return &pb.HelloResponse{Greeting: greeting}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})


	log.Println("Server is running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}