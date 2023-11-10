## gRPC heartbeat with MongoDB

Setup:

```
go mod init heartbeat
go get -u google.golang.org/grpc
```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Run the server:

```
go run server/main.go
```

Run the client:

```
go run client/main.go
```
