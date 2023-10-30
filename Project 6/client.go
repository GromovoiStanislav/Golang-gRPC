package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "main/protos"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatalf("Usage: client <X>")
    }

    x, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatalf("Invalid input: %v", err)
    }

    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewNumberStreamClient(conn)

    stream, err := client.GenerateNumbers(context.Background(), &pb.NumberRequest{Limit: int32(x)})
    if err != nil {
        log.Fatalf("Failed to call GenerateNumbers: %v", err)
    }

    for {
        number, err := stream.Recv()
        if err != nil {
            break
        }
        log.Printf("Received: %d\n", number.Number)
    }
}