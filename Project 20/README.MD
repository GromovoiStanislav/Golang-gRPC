## gRPC Greet examples

We have implemented a simple gRPC server and client with the following functionality:

- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

setup:

```bash
go mod init grpc-demo
go get -u google.golang.org/grpc
```

generate code

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Run the server:

```bash
cd server
go run .
```

Run the client:

```bash
cd client
go run .
```