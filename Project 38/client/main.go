package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc_example/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")


func callUnaryEcho(ctx context.Context, client pb.EchoClient) {
	resp, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message: "hello world"})
	if err != nil {
		log.Fatalf("UnaryEcho: %v", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}


func callBidiStreamingEcho(ctx context.Context, client pb.EchoClient) {
	c, err := client.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("BidiStreamingEcho: %v", err)
	}

	if err := c.Send(&pb.EchoRequest{Message: "hello world"}); err != nil {
		log.Fatalf("Sending echo request: %v", err)
	}
	c.CloseSend()

	for {
		resp, err := c.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Receiving echo response: %v", err)
		}
		fmt.Println("BidiStreaming Echo: ", resp.Message)
	}
}


func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial(%q): %v", *addr, err)
	}
	defer conn.Close()

	ec := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	callUnaryEcho(ctx, ec)

	callBidiStreamingEcho(ctx, ec)
}