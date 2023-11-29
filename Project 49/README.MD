## Creating REST API with gRPC Gateway

This repository contains a simple application that demonstrates how to use grpc and the grpc gateway with golang.

### Setup:

```
go mod init grpc-example
go get -u google.golang.org/grpc

```

### Creating gRPC API based:

```
Setup:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Generate code:
protoc --proto_path=proto proto/*.proto  --go_out=. --go-grpc_out=.
```

### Creating REST API based on gRPC with Gateway

https://github.com/grpc-ecosystem/grpc-gateway

```
Setup:
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/runtime

https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto
https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

Generate code:

protoc -I ./proto --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true proto/user_service.proto
or:
protoc --proto_path=proto --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true proto/user_service.proto

go list -f '{{.Target}}' github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

### Start:

```
go run cmd/main.go
```
