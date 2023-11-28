package main

import (
	"context"
	"log"
	
	"google.golang.org/grpc"

	pb "bookshop/client/pb/inventory"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	
	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("book list: %v", bookList)
}