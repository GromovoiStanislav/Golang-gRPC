## gRPC official example of Encryption TLS

The example for encryption includes example for TLS encryption mechanism

### TLS

TLS is a commonly used cryptographic protocol to provide end-to-end
communication security. In the example, we show how to set up a server
authenticated TLS connection to transmit RPC.

In our `grpc/credentials` package, we provide several convenience methods to
create grpc
[`credentials.TransportCredentials`](https://godoc.org/google.golang.org/grpc/credentials#TransportCredentials)
base on TLS. Refer to the
[godoc](https://godoc.org/google.golang.org/grpc/credentials) for details.

In our example, we use the public/private keys created ahead:

- "server_cert.pem" contains the server certificate (public key).
- "server_key.pem" contains the server private key.
- "ca_cert.pem" contains the certificate (certificate authority)
  that can verify the server's certificate.

On server side, we provide the paths to "server.pem" and "server.key" to
configure TLS and create the server credential using
[`credentials.NewServerTLSFromFile`](https://godoc.org/google.golang.org/grpc/credentials#NewServerTLSFromFile).

On client side, we provide the path to the "ca_cert.pem" to configure TLS and create
the client credential using
[`credentials.NewClientTLSFromFile`](https://godoc.org/google.golang.org/grpc/credentials#NewClientTLSFromFile).
Note that we override the server name with "x.test.example.com", as the server
certificate is valid for \*.test.example.com but not localhost. It is solely for
the convenience of making an example.

Once the credentials have been created at both sides, we can start the server
with the just created server credential (by calling
[`grpc.Creds`](https://godoc.org/google.golang.org/grpc#Creds)) and let client dial
to the server with the created client credential (by calling
[`grpc.WithTransportCredentials`](https://godoc.org/google.golang.org/grpc#WithTransportCredentials))

And finally we make an RPC call over the created `grpc.ClientConn` to test the secure
connection based upon TLS is successfully up.

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
go run server/main.go
```

### Run the client

```
go run client/main.go
```
