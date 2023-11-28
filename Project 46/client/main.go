package main

import (
	"context"
	"log"
	
	"google.golang.org/grpc"

	pb "grpc_example/invoicer"
)

func main() {
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInvoicerClient(conn)
	
	resp, err := client.Create(context.Background(), &pb.CreateRequest{
		From: "John",
		To: "Doe",
		Total: &pb.Amount{
			Currency: "EUR",
			Amount: 20,
		},
	})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("Response: %v", resp)
}