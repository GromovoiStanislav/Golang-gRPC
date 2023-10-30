package main

import (
	"log"
	"net"

	"google.golang.org/grpc"


	pb "main/protos"
)

type server struct{
	pb.UnimplementedNumberStreamServer
}

func (s *server) GenerateNumbers(req *pb.NumberRequest, stream pb.NumberStream_GenerateNumbersServer) error {
    for i := 1; i <= int(req.Limit); i++ {
        if err := stream.Send(&pb.NumberResponse{Number: int32(i)}); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterNumberStreamServer(s, &server{})


	log.Println("Server is running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}