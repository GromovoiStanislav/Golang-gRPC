package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "grpc-example/user"
)

type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
    log.Printf("Received: %d", in.GetId())
    // здест к примеру логика получения пользователя из базы данных или другого сервиса
    return &pb.UserResponse{Id: in.GetId(), Name: "John", Email: "john@example.com"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}