package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "main/protos"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	stream, err := client.CalculateSum(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Отправляем числа в виде стрима
	for _, number := range []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		if err := stream.Send(&pb.SumRequest{Number: number}); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Закрываем стрим и получаем результат
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Sum of numbers: %d\n", response.Result)
}