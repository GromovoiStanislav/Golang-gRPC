package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "main/protos"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Отправка метаданных на сервер
	//md := metadata.Pairs("authorization", "Bearer token123", "user-agent", "gRPC Client","client-data","777")
	md := metadata.New(map[string]string{"authorization": "Bearer token123", "user-agent": "gRPC Client", "client-data":"777"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Отправка запроса на сервер
	req := &pb.HelloRequest{Name: "John"}
	var headers, trailer metadata.MD
	resp, err := client.SayHello(ctx, req, grpc.Header(&headers), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}


	// Получение метаданных от сервера
	fmt.Println("Received server headers:", headers)
	if t, ok := headers["server-metadata"]; ok {
		fmt.Printf("server-metadata from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Println("server-metadata expected but doesn't exist in header")
	}

	// Получение ответа от сервера
	log.Printf("Server response: %s\n", resp.Greeting)


	// Получение метаданных от сервера
	fmt.Println("Received server trailer:", trailer)
	if t, ok := trailer["server-trailer"]; ok {
		fmt.Printf("server-trailer from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Println("server-trailer expected but doesn't exist in trailer")
	}
}