package main

import (
	"log"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/proto"
)

const (
	port = ":8080"
)


func main() {
	//conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}