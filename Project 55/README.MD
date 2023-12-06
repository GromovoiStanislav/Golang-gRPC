## gRPC and Echo/v4 REST API of Products service

### Setup:

```
go mod init grpc-example
go get -u google.golang.org/grpc
go get github.com/labstack/echo/v4
go get github.com/satori/go.uuid
```

### Generate code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=./application/grpc --go-grpc_out=./application/grpc  application/grpc/proto/*.proto
```

### Start the server:

```
go run main.go
```

### Start the gRPC-client:

```
go run grpc-client.go
```
