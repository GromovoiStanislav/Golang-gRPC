package main

import (
	"grpc-example/application/grpc"
	"grpc-example/application/http"
)

func main() {
	webserver := http.NewWebServer()
	go webserver.Serve(":8080")
	grpc.StartGrpcServer()
}