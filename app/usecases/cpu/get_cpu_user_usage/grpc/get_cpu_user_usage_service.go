package grpc

import (
	"context"

	pb "data-provider/usecases/cpu/get_cpu_user_usage/grpc/proto"
)

type Service struct {
	pb.UnimplementedGetCpuUserUsageServiceServer
}

func (s *Service) GetCpuUserUsage(ctx context.Context, req *pb.GetCpuUserUsageRequest) (*pb.GetCpuUserUsageResponse, error) {
	return &pb.GetCpuUserUsageResponse{
		Cpu:      "Intel Xeon",
		AvgUsage: 45.6,
		MaxUsage: 88.9,
		MinUsage: 10.2,
	}, nil
}
