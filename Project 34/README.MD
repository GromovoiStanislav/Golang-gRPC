## gRPC official example of Health

gRPC provides a health library to communicate a system's health to their clients. It works by providing a service definition via the health/v1 api.

By using the health library, clients can gracefully avoid using servers as they encounter issues. Most languages provide an implementation out of box, making it interoperable between systems.

Setup:

```
go mod init health
go get -u google.golang.org/grpc
```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Try it:

```
go run server/main.go -port=50051 -sleep=5s
go run server/main.go -port=50052 -sleep=10s
```

```
go run client/main.go
```
