## gRPC with TLS Greet example

Setup:

```
go mod init grpc-go-example
go get -u google.golang.org/grpc

```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=.  --go-grpc_out=. greeter.proto
```

Start:

```
go run server.go
```

```
go run client.go <name>
```
