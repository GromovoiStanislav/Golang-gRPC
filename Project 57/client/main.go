package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	posts "grpc-example/client/posts"
)

var (
	serverURL = "localhost:10000"
)

func getGRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial(serverURL, opts...)

	if err != nil {
		log.Fatal("Fail to dial %v", err)
	}

	return conn
}

func main() {
	conn := getGRPCClient()

	defer conn.Close()

	client := posts.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	posts, err := client.GetPosts(ctx, &posts.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.GetPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Text)
	}
}