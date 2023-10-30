## gRPC Server stream simple example

setup:

```
go mod init main
go get -u google.golang.org/grpc

```

generate cod

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=./protos --go-grpc_out=./protos protos/number_stream.proto

```

start:

```
go run server.go
go run client.go <Number>
```
