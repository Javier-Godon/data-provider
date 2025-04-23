package main

import (
	"log"
	"net"

	"github.com/Javier-Godon/data-provider/framework"

	// gRPC services
	getCpuSystemUsage "github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage/grpc"
	getCpuUserUsage "github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_user_usage/grpc"
	getFullPrometheusData "github.com/Javier-Godon/data-provider/usecases/prometheus/get_full_prometheus_data/grpc"

	// generated Protobuf packages
	pbCpuSystemUsage "github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage/grpc/proto"
	pbCpuUserUsage "github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_user_usage/grpc/proto"
	pbFullPrometheusData "github.com/Javier-Godon/data-provider/proto/get_full_prometheus_data"

	"google.golang.org/grpc"
)

func main() {
	framework.ReadConfig()

	framework.InitDatabase()
	if framework.DB == nil {
		log.Fatal("Database pool not initialized (nil DB)")
	}

	defer framework.DB.Close()

	log.Println("Application started successfully")
	

	serverPort := "50051"
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register services
	pbCpuSystemUsage.RegisterGetCpuSystemUsageServiceServer(grpcServer, &getCpuSystemUsage.Service{})
	pbCpuUserUsage.RegisterGetCpuUserUsageServiceServer(grpcServer, getCpuUserUsage.NewService())
	pbFullPrometheusData.RegisterGetFullPrometheusDataServiceServer(grpcServer, &getFullPrometheusData.Service{})

	log.Println("gRPC Server is running on port", serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
