## Protocol Buffer

setup:

```
go mod init main
go get github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

generate code

```
protoc --go_out=. proto/*.proto
```

start:

```
go run main.go
```
