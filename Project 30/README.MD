## Protocol Buffer

Setup:

```
go mod init protobufexample
go get github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Generate code:

```
protoc --go_out=. proto/*.proto
```

Start:

```
go run main.go
```
