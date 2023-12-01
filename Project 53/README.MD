## gRPC example of Pokemon application

This is a demo of gRPC application using MongoDB

### Setup:

```
go mod init grpc-pokemon
go get -u google.golang.org/grpc
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo
```

### Generate code:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pokemon/pokemon.proto
```

### Start the server:

```
go run ./server
```

### Run the client

```
go run ./client
```
