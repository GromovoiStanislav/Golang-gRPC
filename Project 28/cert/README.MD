### Step 1: Generate ca.key

```
openssl genrsa -out ca.key 4096
```

### Step 2: Generate certificate

```
openssl req -new -x509 -key ca.key -sha256 -subj "/C=SE/ST=HL/O=Example, INC." -days 365 -out ca.cert
```

### Step 3: Generate the server key

```
openssl genrsa -out server.key 4096
```

### Step 4: Generate the csr

```
openssl req -new -key server.key -out server.csr -config certificate.conff
```

### Step 5: Sign the certificate

```
openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile certificate.conf -extensions req_ext
```
