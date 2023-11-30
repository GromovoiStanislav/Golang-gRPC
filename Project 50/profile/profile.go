package profile

import (
	"log"
	"context"
)

type Server struct {
	UnimplementedProfileServiceServer
}

func (s *Server) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	log.Printf("Receive message body from client: %s", req.GetName())
	return &CreateResponse{Message: "Profile Created!"}, nil
}