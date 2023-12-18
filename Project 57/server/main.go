package main

import (
	"fmt"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"

	posts "grpc-example/server/posts"
)

var (
	serverURL  = "localhost:10000"
	portServer = 10000
)

type Server struct {
	posts.UnimplementedPostServiceServer
}

func (s *Server) GetPosts(ctx context.Context, in *posts.Empty) (*posts.PostList, error) {
	log.Println("GetPosts")

	return &posts.PostList{
		Posts: []*posts.Post{
			&posts.Post{
				Id:    1,
				Title: "Titulo 1",
				Text:  "bla bla bla",
			},
			&posts.Post{
				Id:    2,
				Title: "Titulo 2",
				Text:  "bla bla bla",
			},
		},
	}, nil
}

func main() {
	fmt.Println("Server running on", serverURL)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portServer))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	posts.RegisterPostServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}