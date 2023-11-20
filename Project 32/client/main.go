package main

import (
	"context"
	"flag"
	"log"
	"time"
	"os/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "error_handling/helloworld"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

func main() {
	flag.Parse()

	name := "unknown"
	if u, err := user.Current(); err == nil && u.Username != "" {
		name = u.Username
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, reqName := range []string{"", name} {
		log.Printf("Calling SayHello with Name:%q", reqName)
		
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: reqName})
		if err != nil {
			if status.Code(err) != codes.InvalidArgument {
				log.Printf("Received unexpected error: %v", err)
				continue
			}
			log.Printf("Received error: %v", err)
			continue
		}
		
		log.Printf("Received response: %s", r.Message)
	}
}