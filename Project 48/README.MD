## gRPC simple example of Notes service

This demo gRPC application is a simple note-taking application. You can use the client to:

- Save a note by providing a title and a content
- Load a note by searching a keyword

### Setup:

```
go mod init go-grpc-demo
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative notes/notes.proto
```

### Start the server:

```
go run ./server
```

### Save a note:

```
go run ./client save -title test -content "Lorem ipsum dolor sit amet, fgslasl"
```

### Save a large note:

```
go run ./client savelarge -title lorem -content "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
```

### Load a note:

```
go run ./client load -keyword adipiscing
```
