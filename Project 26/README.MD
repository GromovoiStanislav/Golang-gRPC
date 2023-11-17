## gRPC Hello example with using Cobra

Setup:

```
go mod init grpc-go-example
go get -u google.golang.org/grpc
go get github.com/spf13/cobra
```

Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Start:

```
go run main.go
go run main.go server
go run main.go client <Name>
```
