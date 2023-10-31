package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "main/gen/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTestAPIClient(conn)



	response, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello everyone!"})
	if err != nil {
		log.Fatalf("Create request failed: %v", err)
	}

	log.Println(response)
	log.Println(response.Msg)
	

	

}