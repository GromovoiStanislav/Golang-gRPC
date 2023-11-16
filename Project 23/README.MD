## gRPC with TLS example

We have implemented a simple gRPC server and client with the following functionality:

- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

Run the server:

```bash
cd server
go run main.go -tls=true
```

Run the client:

```bash
cd client
go run main.go -tls=true
```
