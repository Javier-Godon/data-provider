package main

import (
	"log"
	"net"
	"data-provider/framework"

	// gRPC services
	getCpuSystemUsage "data-provider/usecases/cpu/get_cpu_system_usage/grpc"
	getCpuUserUsage "data-provider/usecases/cpu/get_cpu_user_usage/grpc"

	// generated Protobuf packages
	pbCpuSystemUsage "data-provider/usecases/cpu/get_cpu_system_usage/grpc/proto"
	pbCpuUserUsage "data-provider/usecases/cpu/get_cpu_user_usage/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	framework.ReadConfig()

	framework.InitDatabase()

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
    pbCpuUserUsage.RegisterGetCpuUserUsageServiceServer(grpcServer, &getCpuUserUsage.Service{})

    log.Println("gRPC Server is running on port", serverPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
