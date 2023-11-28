package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"

    pb "grpc_example/calculator"
)

type server struct{
    pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
    result := req.Num1 + req.Num2
    return &pb.AddResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
    result := req.Num1 - req.Num2
    return &pb.SubtractResponse{Result: result}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        fmt.Printf("Failed to listen: %v\n", err)
        return
    }

    s := grpc.NewServer()
    pb.RegisterCalculatorServer(s, &server{})

    fmt.Println("Server is listening on port 50051")
    if err := s.Serve(lis); err != nil {
        fmt.Printf("Failed to serve: %v\n", err)
    }
}