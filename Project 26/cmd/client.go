package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	
	"grpc-go-example/pb"
)

var Client = &cobra.Command{
	Use: "client",
	Short: "Run client",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteClient(args[0])
	},
}

func ExecuteClient(name string) {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: name,
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}