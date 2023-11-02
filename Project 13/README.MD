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
or
protoc -I=proto --go_out=. proto/addressbook.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

start:

```
go run list_people_write.go address-book
go run list_people_read.go address-book
```
