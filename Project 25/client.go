package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "grpc-go-example/hello"
)

func main() {
	addr := "localhost:50051"
	cred, err := credentials.NewClientTLSFromFile(
		"./tls/server.crt",//or ca.crt
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	//conn, err := grpc.Dial(addr, grpc.WithInsecure())
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(cred),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := os.Args[1]
	ctx := context.Background()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}