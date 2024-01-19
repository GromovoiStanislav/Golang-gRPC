## Protocol Buffer

#### Setup:

```
go mod init protobuf-example

go get github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### Generate code

```bash
protoc -I . --go_out=. src/phonebook/phonebook.proto
```

#### Start:

```
go run main.go
```
