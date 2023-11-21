## gRPC official example of Multiplex

This example demonstrates basic RPC error handling in gRPC.

Setup:

```
go mod init multiplex
go get -u google.golang.org/grpc

```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Run:

```
go run ./server/main.go
```

```
go run ./client/main.go
```
