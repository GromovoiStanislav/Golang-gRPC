package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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


	
	wg.Add(1)
	go func(cities *pb.CityQuery_Result) {
		defer wg.Done()
		getForecast(client, cities)
	}(cities)

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

func getWeather(client pb.WeatherClient, req *pb.GetTemperature) {
	ctx := context.Background()
	
	stream, err := client.Get(ctx, req)
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
        fmt.Printf("Get: code: %s current: %d\n", data.GetCode(), data.GetCurrent())
    }

}


func getForecast(client pb.WeatherClient, cities *pb.CityQuery_Result) {
	ctx := context.Background()

	stream, err := client.Forecast(ctx)
	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})

	// Отправка запроса в стрим
	go func() {
		defer close(waitc)
		for _, city := range cities.Cities {

			if err := stream.Send(&pb.Forecast{Code: city.GetCode(), Date: time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")}); err != nil {
				log.Fatal(err)
			}
		}	
		// Закрытие отправки данных в стрим после завершения
		if err := stream.CloseSend(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
        data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
        fmt.Printf("Forecast: code: %s current: %d\n", data.Temperature.GetCode(), data.Temperature.GetCurrent())
    }

	// Ожидание завершения отправки и приема данных
	<-waitc
}