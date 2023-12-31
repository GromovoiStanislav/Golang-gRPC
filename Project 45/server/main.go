package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "bookshop/server/pb/inventory"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	
	return &pb.GetBookListResponse{Books: getSampleBooks()}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleBooks() []*pb.Book {

	sampleBooks := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
			Language: proto.String("eng"),
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}