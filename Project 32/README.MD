## gRPC official example of Error handling

This example demonstrates basic RPC error handling in gRPC.

Setup:

```
go mod init error_handling
go get -u google.golang.org/grpc

```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=. --go-grpc_out=. proto/helloworld.proto

```

Run the server, which returns an error if the RPC request's Name field is empty:

```
go run ./server/main.go
```

Then run the client in another terminal, which does two requests: one with an empty Name field and one with it populated with the current username provided by os/user:

```
go run ./client/main.go
```

It should print the status codes it received from the server.
