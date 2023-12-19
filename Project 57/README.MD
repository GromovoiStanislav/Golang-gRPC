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

protoc --go_out=./server --go-grpc_out=./server proto/posts_service.proto
protoc --go_out=./client --go-grpc_out=./client proto/posts_service.proto
```

### Start the server:

```
go run ./server
```

### Run the client

```
go run ./client
```
