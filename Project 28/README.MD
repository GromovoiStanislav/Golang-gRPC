## gRPC with TLS example of using Server side streaming\*!!!

Setup:

```
go mod init grpcexample
go get -u google.golang.org/grpc

go get github.com/mackerelio/go-osstat
```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=./proto  --go-grpc_out=./proto proto/*.proto
```

Start:

```
go run server.go
```

```
go run client.go
```
