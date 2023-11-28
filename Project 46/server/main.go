package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"


	pb "grpc_example/invoicer"
)
// should implement the interface myPkgName.InvoicerServer
type myGRPCServer struct {
    pb.UnimplementedInvoicerServer
}

func (m *myGRPCServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
    log.Println("Create called")
    return &pb.CreateResponse{Pdf: []byte("TODO")}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":3333")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    myInvoicerServer := &myGRPCServer{}
    pb.RegisterInvoicerServer(s, myInvoicerServer)
    log.Printf("server listening at %v", lis.Addr())

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}