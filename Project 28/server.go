package main

import (
	"log"
	"net"
	"time"

	// "github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	hardwaremonitoring "grpcexample/proto"
)

// Server is our struct that will handle the Hardware monitoring Logic
// It will fulfill the gRPC interface generated
type Server struct {
	hardwaremonitoring.UnimplementedHardwareMonitorServer
}

// Monitor is used to start a stream of HardwareStats
func (s *Server) Monitor(req *hardwaremonitoring.EmptyRequest, stream hardwaremonitoring.HardwareMonitor_MonitorServer) error {
	// Start a ticker that executes each 2 seconds
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		// Exit on stream context done
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			// Grab stats and output
			hwStats, err := s.GetStats()
			if err != nil {
				log.Println(err.Error())
			}

			// Send the Hardware stats on the stream
			err = stream.Send(hwStats)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

// GetStats will extract system stats and output a Hardware Object, or an error
// if extraction fails
func (s *Server) GetStats() (*hardwaremonitoring.HardwareStats, error) {
	// // Extract CPU stats
	// cpu, err := cpu.Get()
	// if err != nil {
	// 	return nil, err
	// }
	
	// Extarcyt Memory statas
	mem, err := memory.Get()
	if err != nil {
		return nil, err
	}
	
	// Extarcyt uptime statas
	upt, err := uptime.Get()
	if err != nil {
		return nil, err
	}

	// Create our response object
	hwStats := &hardwaremonitoring.HardwareStats{
		// Cpu:        int32(cpu.Total),
		MemoryFree: int32(mem.Free),
		MemoryUsed: int32(mem.Used),
		MemoryTotal: int32(mem.Total),
		UptimeSeconds: int32(upt.Seconds()),
	}

	return hwStats, nil
}


func main() {
	log.Println("Welcome to streaming HW monitoring")

	// We Generate a TLS grpc API
	gRPCserver, err := GenerateTLSApi("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	
	// Create a server object of the type we created 
	s := &Server{}
	
	// Regiser our server as a gRPC server
	hardwaremonitoring.RegisterHardwareMonitorServer(gRPCserver, s)

	// Start serving
	log.Printf("gRPC server listening on " + "127.0.0.1:50051")
	if err := gRPCserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// GenerateTLSApi will load TLS certificates and key and create a grpc server with those.
func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Create a gRPC server
	gRPCserver := grpc.NewServer(
		grpc.Creds(cred),
	)
	return gRPCserver, nil
}

