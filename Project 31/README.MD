## gRPC official example of Error details

This example demonstrates the use of status details in grpc errors.

Setup:

```
go mod init error_details
go get -u google.golang.org/grpc

```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=. --go-grpc_out=. proto/helloworld.proto

```

Run the server:

```
go run ./server/main.go
```

Then run the client in another terminal:

```
go run ./client/main.go
```

It should succeed and print the greeting it received from the server. Then run the client again:

```
go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
