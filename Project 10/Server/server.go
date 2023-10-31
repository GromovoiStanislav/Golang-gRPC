package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "main/gen/proto"
)

type testAPIServer struct {
	pb.UnimplementedTestAPIServer
}

func (s testAPIServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s testAPIServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Id: req.Id, Name: "John", Age: 25, Email: "john@mail.com"}, nil
}

func main() {
	go func(){
		//mux
		mux := runtime.NewServeMux()
		//regester
		pb.RegisterTestAPIHandlerServer(context.Background(), mux, &testAPIServer{})
		//http server
		log.Fatalln(http.ListenAndServe(":8081", mux))
	}()



	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterTestAPIServer(grpcServer, &testAPIServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}