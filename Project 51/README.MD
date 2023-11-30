## gRPC simple example of Geometry service

### Setup:

```
go mod init grpc-example
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
or:
protoc -Iproto/ --go_out=. --go_opt=module=grpc-example --go-grpc_out=. --go-grpc_opt=module=grpc-example proto/*.proto
```

### Start the server:

```
go run cmd/server/main.go
```

### Run the client

```
go run cmd/client/main.go
```
