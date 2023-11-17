### Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)

```
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -sha256 -days 365 -key ca.key -out ca.crt -subj "/CN=localhost"
```

### Step 2: Generate the Server Private Key (server.key)

```
openssl genrsa -passout pass:1111 -des3 -out server.key 4096
```

### Step 3: Convert the server certificate to .pem format (server.pem) - usable by gRPC

```
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
```

### Step 4: Get a certificate signing request from the CA (server.csr)

```
openssl req -passin pass:1111 -new -sha256 -key server.key -out server.csr -subj "/CN=localhost" -config ssl.cnf
```

### Step 5: Sign the certificate with the CA we created (it's called self signing) - server.crt

```
openssl x509 -req -passin pass:1111 -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions req_ext -extfile ssl.cnf
```
