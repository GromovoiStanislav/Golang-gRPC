package main

import (
	"context"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"

	pb "main/invoicer"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInvoicerClient(conn)

	request := &pb.CreateRequest{
		Amount: &pb.Amount{
			Amount:   1000,
			Currency: "USD",
		},
		From:      "Sender",
		To:        "Receiver",
		VATNumber: "123456789",
	}

	response, err := client.Create(context.Background(), request)
	if err != nil {
		log.Fatalf("Create request failed: %v", err)
	}

	log.Printf("Received PDF: %v\n", response.Pdf)
	log.Printf("Received PDF: %v\n", string(response.Pdf))
	log.Printf("Received DOCX: %v\n", response.Docx)
	log.Printf("Received DOCX: %v\n", string(response.Docx))


	// Сохранение полученных файлов
	err = ioutil.WriteFile("output.pdf", response.Pdf, 0644)
	if err != nil {
		log.Fatalf("Failed to write PDF to file: %v", err)
	}

	err = ioutil.WriteFile("output.docx", response.Docx, 0644)
	if err != nil {
		log.Fatalf("Failed to write DOCX to file: %v", err)
	}

}