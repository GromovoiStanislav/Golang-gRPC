## gRPC simple example of Bookshop service

### Setup:

```
go mod init bookshop
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client
```

### Start the server:

```
go run server/main.go
```

### Run the client

```
go run client/main.go
```
