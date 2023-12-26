package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "grpc-example/currency"
)

type server struct{
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
    // In a real application, you'd fetch exchange rates from an external source.
    // For this example, we'll use simple conversion rates.
    rates := map[string]float32{
        "USD": 1.0,
        "EUR": 0.85,
        "JPY": 110.0,
    }

    fromCurrency := req.GetFromCurrency()
    toCurrency := req.GetToCurrency()
    amount := req.GetAmount()

    fromRate, ok := rates[fromCurrency]
    if !ok {
        return nil, grpc.Errorf(codes.InvalidArgument, "unknown from currency")
    }

    toRate, ok := rates[toCurrency]
    if !ok {
        return nil, grpc.Errorf(codes.InvalidArgument, "unknown to currency")
    }

    convertedAmount := amount * (fromRate / toRate)

    return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterCurrencyConverterServer(s, &server{})

    log.Println("gRPC server started on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}