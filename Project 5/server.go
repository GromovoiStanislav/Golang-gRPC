package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "main/protos"
)

type server struct{
	pb.UnimplementedSumServiceServer
}

func (s *server) CalculateSum(stream pb.SumService_CalculateSumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		sum += req.Number
	}
	return stream.SendAndClose(&pb.SumResponse{Result: sum})
}


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})

	log.Println("Server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}