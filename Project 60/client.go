package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"

	pb "grpc-example/weather"
)



func getGRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial("localhost:9090", opts...)

	if err != nil {
		log.Fatal("Fail to dial %v", err)
	}

	return conn
}

func main() {
	conn := getGRPCClient()
	defer conn.Close()

	client := pb.NewWeatherClient(conn)

	cities := citiesWeather(client);
	fmt.Println(cities.Cities)

	for _, city := range cities.Cities {
		fmt.Println(city.GetCode())
		fmt.Println(city.GetName())
		fmt.Println()
    }



	var wg sync.WaitGroup

	for _, city := range cities.Cities {
		wg.Add(1)
		go func(cityCode string) {
			defer wg.Done()
			getWeather(client, &pb.GetTemperature{Code: cityCode})
		}(city.GetCode())
	}

	wg.Wait()


}

func citiesWeather(client pb.WeatherClient) *pb.CityQuery_Result {
	ctx := context.Background()
	
	res, err := client.Cities(ctx, &pb.CityQuery{})
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func getWeather(client pb.WeatherClient, city *pb.GetTemperature) {
	ctx := context.Background()
	
	stream, err := client.Get(ctx, city)
	if err != nil {
		log.Fatal(err)
	}

	for {
        data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
        fmt.Printf("code: %s current: %d\n", data.GetCode(), data.GetCurrent())
    }

}