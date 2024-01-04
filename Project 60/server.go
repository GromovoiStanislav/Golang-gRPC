package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "grpc-example/weather"
)

type NWeatherServer struct {
	*pb.UnimplementedWeatherServer
}


func (*NWeatherServer) Cities(context.Context, *pb.CityQuery) (*pb.CityQuery_Result, error) {
	return &pb.CityQuery_Result{
		Cities: []*pb.City{
			{
				Code: "TR_ANTALYA",
				Name: "Antalya",
			},
			{
				Code: "CA_VANCOUVER",
				Name: "Vancouver",
			},
		},
	}, nil
}

func Rand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func (*NWeatherServer) Get(get *pb.GetTemperature, resp pb.Weather_GetServer) error {
	temp := int(70)
	for i := 0; i < Rand(10, 30); i++ {
		temp = temp - Rand(0, 4)
		if err := resp.Send(&pb.Temperature{
			Code:    get.GetCode(),
			Current: int32(temp),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(Rand(100, 500)))
	}
	return nil
}

func (*NWeatherServer) Forecast(stream pb.Weather_ForecastServer) error  {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		cod := req.GetCode()
		//date := req.GetDate()

		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.Forecast_Result{Temperature: &pb.Temperature{
				Code:    cod,
				Current: int32(Rand(10, 30)),
			}}); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * time.Duration(Rand(100, 500)))
		}
	
	}
	
}


func main() {
    port := 9090
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterWeatherServer(s, &NWeatherServer{})


    log.Printf("gRPC server started on port: %d", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}