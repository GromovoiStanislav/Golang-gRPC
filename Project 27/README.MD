## gRPC with TLS example

Setup:

```
go mod init grpcexample
go get -u google.golang.org/grpc
```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=./pingpong  --go-grpc_out=./pingpong pingpong/service.proto
```

Start:

```
go run server.go
```

```
go run client.go
```
