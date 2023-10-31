## Creating REST API based on gRPC with Gateway

setup:

```
go mod init main
go get -u google.golang.org/grpc

```

### Creating gRPC API based

```
setup:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate code:
protoc --proto_path=proto proto/*.proto  --go_out=gen/ --go-grpc_out=gen/
```

start:

```
go run .\Server\server.go
go run .\Client\client.go
```

### Creating REST API based on gRPC with Gateway

https://github.com/grpc-ecosystem/grpc-gateway

```
setup:
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/runtime

https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto
https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

generate code:
protoc -I . --grpc-gateway_out ./gen/ --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/test.proto

go list -f '{{.Target}}' github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```
