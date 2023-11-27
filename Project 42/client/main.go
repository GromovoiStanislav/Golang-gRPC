package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "grpc_example/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func unaryCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
	// Creates a context with a one second deadline for the RPC.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.EchoRequest{Message: message}

	_, err := c.UnaryEcho(ctx, req)
	got := status.Code(err)
	fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
}

func streamingCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
	// Creates a context with a one second deadline for the RPC.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Printf("Stream err: %v", err)
		return
	}

	err = stream.Send(&pb.EchoRequest{Message: message})
	if err != nil {
		log.Printf("Send error: %v", err)
		return
	}
	stream.CloseSend()


	for {
		_, err = stream.Recv()
		if err == io.EOF {
			break
		}

		got := status.Code(err)
		fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
		if err != nil {
			break 
		}
	}
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	// A successful request
	//unaryCall(c, 1, "world", codes.OK)

	// Exceeds deadline
	//unaryCall(c, 2, "delay", codes.DeadlineExceeded)

	// A successful request with propagated deadline
	//unaryCall(c, 3, "[propagate me]world", codes.OK)

	// Exceeds propagated deadline
	//unaryCall(c, 4, "[propagate me][propagate me]world", codes.DeadlineExceeded)



	// Receives a response from the stream successfully
	streamingCall(c, 5, "world", codes.OK)

	// Exceeds deadline before receiving a response
	streamingCall(c, 6, "delay", codes.DeadlineExceeded)

	// Receives a response from the stream successfully
	streamingCall(c, 7, "[propagate me]world", codes.OK)

	// Exceeds propagated deadline before receiving a response
	streamingCall(c, 8, "[propagate me][propagate me]world", codes.DeadlineExceeded)
}