## gRPC Currency Converter service

### Setup:

```
go mod init grpc-example
go get -u google.golang.org/grpc
```

### Generate code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Start the server:

```
go run server.go
```

### Run the client:

```
go run client.go
```
