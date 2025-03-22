package grpc

import (
	"context"

	"data-provider/usecases/cpu/get_cpu_system_usage"
	pb "data-provider/usecases/cpu/get_cpu_system_usage/grpc/proto"
	"data-provider/usecases/cpu/get_cpu_system_usage/mediator"
)

type Service struct {
	pb.UnimplementedGetCpuSystemUsageServiceServer
}

func (s *Service) GetCpuSystemUsage(ctx context.Context, req *pb.GetCpuSystemUsageRequest) (*pb.GetCpuSystemUsageResponse, error) {
	command := fromRequestToQuery(req)
	result := mediator.Send(command)
	return fromResultToResponse(result), nil
}

func fromRequestToQuery(req *pb.GetCpuSystemUsageRequest) get_cpu_system_usage.GetCpuSystemUsageQuery {
	return get_cpu_system_usage.GetCpuSystemUsageQuery{
		DateFrom: req.DateFrom,
		DateTo:   req.DateTo,
	}
}

func fromResultToResponse(result get_cpu_system_usage.GetCpuSystemUsageResult) *pb.GetCpuSystemUsageResponse {
	var usages []*pb.CpuUsage
	for _, usage := range result.Usages {
		usages = append(usages, &pb.CpuUsage{
			Cpu:      usage.Cpu,
			AvgUsage: usage.AvgUsage,
			MaxUsage: usage.MaxUsage,
			MinUsage: usage.MinUsage,
		})
	}

	return &pb.GetCpuSystemUsageResponse{
		Usages: usages,
	}
}
