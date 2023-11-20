package main

import (
	"context"
	//"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//"google.golang.org/grpc/credentials"

	pb "grpcalc/pb"
)

func main() {
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address in the format of host:port",
	)
	flag.Parse()

	//creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{
		//grpc.WithTransportCredentials(creds),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	//conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	
	client := pb.NewCalculatorClient(conn)

	add(client, 10, 4)
	subtract(client, 10, 4)
	multiply(client, 10, 4)
	divide(client, 6, 0)
	divide(client, 6, 3)
	sum(client, []int64{10, 10, 10, 10, 10})
}



func add(client pb.CalculatorClient, a, b int64) {
	res, err := client.Add(context.Background(), &pb.CalculationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Println("error sending request:", err)
	} else {
		fmt.Println("Add result:", res.Result)
	}
}

func subtract(client pb.CalculatorClient, a, b int64) {
	res, err := client.Subtract(context.Background(), &pb.CalculationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Println("error sending request:", err)
	} else {
		fmt.Println("Subtract result:", res.Result)
	}
}

func multiply(client pb.CalculatorClient, a, b int64) {
	res, err := client.Multiply(context.Background(), &pb.CalculationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Println("error sending request:", err)
	} else {
		fmt.Println("Multiply result:", res.Result)
	}
}


func divide(client pb.CalculatorClient, a, b int64) {
	res, err := client.Divide(context.Background(), &pb.CalculationRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Println("error sending request:", err)
	} else {
		fmt.Println("Divide result:", res.Result)
	}
}

func sum(client pb.CalculatorClient, numbers []int64) {
	res, err := client.Sum(context.Background(), &pb.NumbersRequest{
		Numbers: numbers,
	})
	if err != nil {
		log.Println("error sending request:", err)
	} else {
		fmt.Println("Sum result:", res.Result)
	}	
}