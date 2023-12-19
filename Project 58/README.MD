## gRPC Reflection

Adding this service to your server will allow clients such as postman to dynamically load the API specification from your running application rather than needing to pass around and load proto files manually.

<image src="example.gif" alt="preview">

### Setup:

```
go mod init grpc-example
go get -u google.golang.org/grpc
```

### Generate code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Start the server:

```
go run .\main.go
```
