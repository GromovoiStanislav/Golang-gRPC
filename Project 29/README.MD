## gRPC Calculator Service example

setup:

```
go mod init grpcalc
go get -u google.golang.org/grpc

```

generate code

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/calculator.proto
```

start:

```
go run .\server\main.go
go run .\client\main.go
```
