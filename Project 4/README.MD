## gRPC Hello example

setup:

```
go mod init main
go get -u google.golang.org/grpc

```

generate code

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=./hello --go-grpc_out=./hello hello/hello.proto

```

start:

```
go run server.go
go run client.go
```
