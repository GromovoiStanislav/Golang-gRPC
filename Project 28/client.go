package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	hardwaremonitoring "grpcexample/proto"
)

func main() {
	ctx := context.Background()
	// Load our TLS certificate and use grpc/credentials to create new transport credentials
	creds := credentials.NewTLS(loadTLSCfg())
	// Create a new connection using the transport credentials
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	// Close connection when we are done
	defer conn.Close()

	// A new GRPC client to use
	client := hardwaremonitoring.NewHardwareMonitorClient(conn)

	// Call Monitor to receive the Stream of data
	// With an empty request
	emptyreq := &hardwaremonitoring.EmptyRequest{}
	// call Monitor function, this will return a stream of data
	stream, err := client.Monitor(ctx, emptyreq)
	if err != nil {
		panic(err)
	}
	
	// Create a timer to cancel
	stop := time.NewTicker(7 * time.Second)
	// Itterate stream
	for {
		select {
		case <-stop.C:
			// Tell the Server to close this Stream, used to clean up running on the server
			err := stream.CloseSend()
			if err != nil {
				log.Fatal("Failed to close stream: ", err.Error())
			}
			return
		default:
			// Recieve on the stream
			res, err := stream.Recv()
			if err != nil {
				panic(err)
			}
			log.Println("New Hardware state receieved")
			log.Printf("CPU Usage: %d%%\n", res.Cpu)
			log.Printf("Memory Used: %d bytes\n", res.MemoryUsed)
			log.Printf("Memory Free: %d bytes\n", res.MemoryFree)
			log.Printf("Memory Total: %d bytes\n", res.MemoryTotal)
			log.Printf("Uptime Seconds: %d\n", res.UptimeSeconds)
		}
	}
}

// loadTLSCfg will load a certificate and create a tls config
func loadTLSCfg() *tls.Config {
	b, _ := ioutil.ReadFile("./cert/server.crt")
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatal("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	return config
}