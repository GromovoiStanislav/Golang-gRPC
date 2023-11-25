package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	ecpb "grpc_example/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
const fallbackToken = "some-secret-token"


// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
func logger(format string, a ...any) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}


// unaryInterceptor is an example unary interceptor.
func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//До вызова
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["authorization"])==0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+fallbackToken)
	}
	
	start := time.Now()
 	
	// Вызов
	err := invoker(ctx, method, req, reply, cc, opts...)

	//После вызова
	end := time.Now()
	logger("RPC: %s, start time: %s, end time: %s, err: %v", method, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}

// wrappedStream  wraps around the embedded grpc.ClientStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m any) error {
	logger("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m any) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// streamInterceptor is an example stream interceptor.
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//До вызова
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["authorization"])==0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+fallbackToken)
	}



	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}


func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}


func callBidiStreamingEcho(client ecpb.EchoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := client.BidirectionalStreamingEcho(ctx)
	if err != nil {
		return
	}
	for i := 0; i < 5; i++ {
		if err := c.Send(&ecpb.EchoRequest{Message: fmt.Sprintf("Request %d", i+1)}); err != nil {
			log.Fatalf("failed to send request due to error: %v", err)
		}
	}
	c.CloseSend()

	for {
		resp, err := c.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response due to error: %v", err)
		}
		fmt.Println("BidiStreaming Echo: ", resp.Message)
	}
}


func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor), grpc.WithStreamInterceptor(streamInterceptor))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Make a echo client and send RPCs.
	rgc := ecpb.NewEchoClient(conn)

	callUnaryEcho(rgc, "hello world")
	time.Sleep(time.Second*1)
	callBidiStreamingEcho(rgc)
}