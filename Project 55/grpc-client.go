package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "grpc-example/application/grpc/pb"
)

func main() {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Создаем gRPC клиент
	client := pb.NewProductServiceClient(conn)

	// Вызываем методы сервера:
	createProduct(client, "New Product")
	time.Sleep(time.Second*1)
	callListMethod(client)
}


func createProduct(client pb.ProductServiceClient, productName string) (*pb.ProductResult, error) {
	// Вызываем метод CreateProduct
	res, err := client.CreateProduct(context.Background(), &pb.Product{
		Name: productName,
	})
	if err != nil {
		log.Println("error while calling CreateProduct: %v", err)
		return nil, fmt.Errorf("error while calling CreateProduct: %v", err)
	}

	fmt.Printf("CreateProduct Response: ID=%v, Name=%v\n", res.GetId(), res.GetName())
	return res, nil
}


func callListMethod(client pb.ProductServiceClient) error {
	// Вызываем метод List
	listStream, err := client.List(context.Background(), &pb.Empty{})
	
	if err != nil {
		log.Println("error while calling List: %v", err)
		return err
	}

	for {
		// Читаем данные из стрима
		productResult, err := listStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error while receiving from List: %v", err)
			return err
		}

		// Выводим результат
		fmt.Printf("List Response: ID=%v, Name=%v\n", productResult.Id, productResult.Name)
	}
}