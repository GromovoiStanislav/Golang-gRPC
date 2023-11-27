## gRPC official example of Metadata

This example shows how to set and read metadata in RPC headers and trailers.
Please see
[grpc-metadata.md](./grpc-metadata.md)
for more information.

#### Setup:

```
go mod init grpc_example
go get -u google.golang.org/grpc
```

#### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

#### Start the server:

```
go run server/main.go
```

#### Run the client

```
go run client/main.go
```
