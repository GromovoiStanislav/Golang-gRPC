package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pingpong "grpcexample/pingpong"
)


type Server struct {
	pingpong.UnimplementedPingPongServer
}

func (s *Server) Ping(ctx context.Context, ping *pingpong.PingRequest) (*pingpong.PongResponse, error) {
	return &pingpong.PongResponse{
		Ok: true,
	}, nil
}


func main() {

	// We Generate a TLS grpc API
	apiserver, err := GenerateTLSApi("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	
	s := &Server{}
	
	pingpong.RegisterPingPongServer(apiserver, s)

	// Start serving
	log.Printf("gRPC server listening on " + "127.0.0.1:50051")
	if err := apiserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// GenerateTLSApi will load TLS certificates and key and create a grpc server with those.
func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.Creds(cred),
	)
	return s, nil
}

