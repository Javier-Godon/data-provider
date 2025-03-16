package grpc

import (
	"context"

	pb "data-provider/usecases/cpu/get_cpu_system_usage/grpc/proto"
)

type Service struct {
	pb.UnimplementedGetCpuSystemUsageServiceServer
}

func (s *Service) GetCpuSystemUsage(ctx context.Context, req *pb.GetCpuSystemUsageRequest) (*pb.GetCpuSystemUsageResponse, error) {
	return &pb.GetCpuSystemUsageResponse{
		Cpu:      "Intel Xeon",
		AvgUsage: 45.6,
		MaxUsage: 88.9,
		MinUsage: 10.2,
	}, nil
}
