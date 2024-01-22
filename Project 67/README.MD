## Protocol Buffer (google/protobuf/timestamp)

#### Setup:

```
go mod init protobuf-example

go get github.com/golang/protobuf/proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### Generate code

```bash
protoc --go_out=. proto/*.proto
```

#### Start:

```
go run main.go
```
