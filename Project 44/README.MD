## gRPC official example of Encryption mTLS

The example for encryption includes example for mTLS encryption mechanism

### mTLS

In mutual TLS (mTLS), the client and the server authenticate each other. gRPC
allows users to configure mutual TLS at the connection level.

In normal TLS, the server is only concerned with presenting the server
certificate for clients to verify. In mutual TLS, the server also loads in a
list of trusted CA files for verifying client presented certificates with.
This is done via setting
[`tls.Config.ClientCAs`](https://pkg.go.dev/crypto/tls#Config.ClientCAs)
to the list of trusted CA files,
and setting [`tls.config.ClientAuth`](https://pkg.go.dev/crypto/tls#Config.ClientAuth)
to [`tls.RequireAndVerifyClientCert`](https://pkg.go.dev/crypto/tls#RequireAndVerifyClientCert).

In normal TLS, the client is only concerned with authenticating the server by
using one or more trusted CA file. In mutual TLS, the client also presents its
client certificate to the server for authentication. This is done via setting
[`tls.Config.Certificates`](https://pkg.go.dev/crypto/tls#Config.Certificates).

### Setup:

```
go mod init grpc_example
go get -u google.golang.org/grpc
```

### Generate code:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Start the server:

```
go run ./server
```

### Run the client

```
go run ./client
```
