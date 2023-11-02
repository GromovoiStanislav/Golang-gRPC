## gRPC simple example

setup:

```
go mod init main
go get -u google.golang.org/grpc

```

generate code

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=. --go-grpc_out=. proto/usermgmt.proto
#or
protoc --go_out=. --go-grpc_out=. proto/*.proto

```

start:

```
go run server.go
go run client.go
```
