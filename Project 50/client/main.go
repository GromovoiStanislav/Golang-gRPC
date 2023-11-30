package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"grpc_example/profile"
)

func main() {
	var conn *grpc.ClientConn
    
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := profile.NewProfileServiceClient(conn)

	createReq := profile.CreateRequest{Name: "Hakan ENSEP", Id:123, IsVerified: true}
	response, err := c.Create(context.Background(), &createReq)
	if err != nil {
		log.Fatalf("Error Profile Create: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}