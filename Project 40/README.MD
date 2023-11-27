## gRPC official example of Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the context passed to the RPC call. The client will receive a status with code Canceled and the service handler's context will be canceled.

### Setup:

```
go mod init grpc_example
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Start the server:

```
go run server/main.go
```

### Run the client

```
go run client/main.go
```
