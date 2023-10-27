package main

import (
    "context"
    "log"
    "os"

    "google.golang.org/grpc"
    pb "main/protos"
)

func main() {
    // Устанавливаем подключение к серверу
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    // Создаем клиент
    client := pb.NewHelloServiceClient(conn)


     // Определяем имя из аргументов командной строки или используем "Hello" по умолчанию
     name := "World"
     if len(os.Args) > 1 {
         name = os.Args[1]
     }

    // Отправляем запрос на сервер
    response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("Error calling SayHello: %v", err)
    }

    // Выводим ответ от сервера
    log.Printf("Response from server: %s", response.Message)
}
