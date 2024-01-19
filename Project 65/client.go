package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-example/user"
)

const (
    address     = "localhost:50051"
)

func main() {
    // установка соединения с сервером
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewUserServiceClient(conn)

    // установим контекст с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // запрс на получение пользователя
    r, err := c.GetUser(ctx, &pb.UserRequest{Id: 123})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("User: %d %s %s", r.GetId(), r.GetName(), r.GetEmail())
}