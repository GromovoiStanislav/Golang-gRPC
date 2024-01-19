## Protocol Buffer\*

#### Setup:

```
go mod init protobuf-example

go get github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### Generate code

```bash
protoc -I . --go_out=. protos/simple/simple.proto
protoc -I . --go_out=. protos/enum_example/enum_example.proto
protoc -I . --go_out=. protos/complex/complex.proto
protoc -I . --go_out=. protos/addressbook/addressbook.proto
```

#### Start:

```
go run main.go
```
