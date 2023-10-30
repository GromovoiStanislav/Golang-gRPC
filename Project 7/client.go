package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "main/protos"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMathServiceClient(conn)
	stream, err := client.CalculateSquare(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 5}

	go func() {
		for _, number := range numbers {
			req := &pb.NumberRequest{Number: number}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	for {
		res, err := stream.Recv()
		if err != nil {
			st, ok := status.FromError(err)
			if !ok || st.Code() != codes.Canceled {
                log.Println("stream Canceled")
			}
			break
		}
		log.Printf("Square: %d\n", res.GetSquare())
	}
}