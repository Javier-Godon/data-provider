package grpc

import (
	"context"

	"data-provider/usecases/cpu/get_cpu_user_usage"
	pb "data-provider/usecases/cpu/get_cpu_user_usage/grpc/proto"
	"data-provider/usecases/cpu/get_cpu_user_usage/mediator"
)

type Service struct {
	pb.UnimplementedGetCpuUserUsageServiceServer
}

func (s *Service) GetCpuUserUsage(ctx context.Context, req *pb.GetCpuUserUsageRequest) (*pb.GetCpuUserUsageResponse, error) {
	command := fromRequestToQuery(req)
	result := mediator.Send(command)
	return fromResultToResponse(result), nil
}

func fromRequestToQuery(req *pb.GetCpuUserUsageRequest) get_cpu_user_usage.GetCpuUserUsageQuery {
	return get_cpu_user_usage.GetCpuUserUsageQuery{
		DateFrom: req.DateFrom,
		DateTo:   req.DateTo,
	}
}

func fromResultToResponse(result get_cpu_user_usage.GetCpuUserUsageResult) *pb.GetCpuUserUsageResponse {
	var usages []*pb.CpuUsage
	for _, usage := range result.Usages {
		usages = append(usages, &pb.CpuUsage{
			Cpu:      usage.Cpu,
			AvgUsage: usage.AvgUsage,
			MaxUsage: usage.MaxUsage,
			MinUsage: usage.MinUsage,
		})
	}

	return &pb.GetCpuUserUsageResponse{
		Usages: usages,
	}
}
