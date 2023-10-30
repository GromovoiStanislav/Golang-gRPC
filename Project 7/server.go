package main

import (
	"log"
	"net"

	"google.golang.org/grpc"


	pb "main/protos"
)

type server struct{
	pb.UnimplementedMathServiceServer
}

func (s *server) CalculateSquare(stream pb.MathService_CalculateSquareServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		number := req.GetNumber()
		square := number * number

		res := &pb.NumberResponse{Square: square}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterMathServiceServer(s, &server{})


	log.Println("Server is running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}