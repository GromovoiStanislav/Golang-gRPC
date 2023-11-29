package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
    
	apiv1 "grpc-example/usersvcapi/v1"
)

func main() {
	// Flags.
	fs := flag.NewFlagSet("", flag.ExitOnError)
	grpcAddr := fs.String("grpc-addr", ":6565", "grpc address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := apiv1.NewUserServiceClient(conn)

	// Example: Create a new user.
	createUserResponse, err := client.CreateUser(context.Background(), &apiv1.CreateUserRequest{
		User: &apiv1.UserWrite{
			Name: "John Doe",
			Type: apiv1.UserType_USER_TYPE_ADMIN,
		},
	})
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("Created user with ID: %s\n", createUserResponse.Id)

	// Example: Get user details.
	getUserResponse, err := client.GetUser(context.Background(), &apiv1.GetUserRequest{
		Id: createUserResponse.Id,
	})
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	fmt.Printf("User details: %+v\n", getUserResponse.User)
}