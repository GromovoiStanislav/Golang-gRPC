## Protocol Buffer

setup:

```
go mod init main
go get -u google.golang.org/grpc
```

generate code

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto
```

start:

```
go run add_person.go address-book

```
