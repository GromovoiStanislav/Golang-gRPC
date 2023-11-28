package main

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/grpc"

    pb "grpc_example/calculator"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewCalculatorClient(conn)

    addResp, err := c.Add(context.Background(), &pb.AddRequest{Num1: 10, Num2: 5})
    if err != nil {
        log.Fatalf("Error calling Add: %v", err)
    }
    fmt.Printf("Add result: %d\n", addResp.Result)

    subResp, err := c.Subtract(context.Background(), &pb.SubtractRequest{Num1: 10, Num2: 5})
    if err != nil {
        log.Fatalf("Error calling Subtract: %v", err)
    }
    fmt.Printf("Subtract result: %d\n", subResp.Result)
}