## gRPC simple example of Calculator service

### Setup:

```
go mod init grpc_example
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/calculator.proto
```

### Start the server:

```
go run ./server
```

### Run the client

```
go run ./client
```
