package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ecpb "multiplex/echo"
	hwpb "multiplex/helloworld"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")


func callSayHello(client hwpb.GreeterClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &hwpb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("client.SayHello(_) = _, %v", err)
	}
	fmt.Println("Greeting: ", r.Message)
}

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("--- calling helloworld.Greeter/SayHello ---")
	// Make a greeter client and send an RPC.
	hwc := hwpb.NewGreeterClient(conn)
	callSayHello(hwc, "multiplex")

	fmt.Println()
	
	fmt.Println("--- calling echo.Echo/UnaryEcho ---")
	// Make a routeguild client with the same ClientConn.
	rgc := ecpb.NewEchoClient(conn)
	callUnaryEcho(rgc, "this is examples/multiplex")
}