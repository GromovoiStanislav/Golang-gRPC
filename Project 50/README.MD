## gRPC simple example of Profile service

### Setup:

```
go mod init grpc_example
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/profile.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/profile.proto
```

### Start the server:

```
go run ./server
```

### Run the client

```
go run ./client
```
