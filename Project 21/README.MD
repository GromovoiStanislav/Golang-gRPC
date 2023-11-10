## gRPC CRUDAPI example

setup:

```
go mod init moviesapp
go get -u google.golang.org/grpc
```

generate code

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
```

start:

```
go run server/main.go
go run client/main.go
```
