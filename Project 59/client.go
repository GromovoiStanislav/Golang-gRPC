package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "grpc-example/currency"
)



func getGRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial("localhost:50051", opts...)

	if err != nil {
		log.Fatal("Fail to dial %v", err)
	}

	return conn
}

func main() {
	conn := getGRPCClient()
	defer conn.Close()

	client := pb.NewCurrencyConverterClient(conn)

	convertCurrency(client, "USD", "EUR", 100);
}

func convertCurrency(client pb.CurrencyConverterClient, fromCurrency, toCurrency string, amount int)  {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	res, err := client.Convert(ctx, &pb.ConvertRequest{
		FromCurrency: fromCurrency,
		ToCurrency: toCurrency,
		Amount: 100,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Converted",amount,fromCurrency,"to",res.ConvertedAmount,fromCurrency)

}