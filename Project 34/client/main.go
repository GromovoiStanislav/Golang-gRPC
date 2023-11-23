package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	
	pb "health/echo"
)

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

func callUnaryEcho(c pb.EchoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{})
	if err != nil {
		fmt.Println("UnaryEcho: _, ", err)
	} else {
		fmt.Println("UnaryEcho: ", r.GetMessage())
	}
}

func main() {
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: "localhost:50051"},
			{Addr: "localhost:50052"},
		},
	})

	address := fmt.Sprintf("%s:///unused", r.Scheme())

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	conn, err := grpc.Dial(address, options...)
	if err != nil {
		log.Fatalf("grpc.Dial(%q): %v", address, err)
	}
	defer conn.Close()

	echoClient := pb.NewEchoClient(conn)

	for {
		callUnaryEcho(echoClient)
		time.Sleep(time.Second)
	}
}