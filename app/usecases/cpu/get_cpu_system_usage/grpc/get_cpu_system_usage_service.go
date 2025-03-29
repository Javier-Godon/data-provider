package grpc

import (
	"context"
	"log"
	
	"github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage"
	pb "github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage/grpc/proto"
	"github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage/mediator"
)

type Service struct {
	pb.UnimplementedGetCpuSystemUsageServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetCpuSystemUsage(ctx context.Context, req *pb.GetCpuSystemUsageRequest) (*pb.GetCpuSystemUsageResponse, error) {
	log.Printf("GetCpuSystemUsage called with dateFrom=%d, dateTo=%d", req.DateFrom, req.DateTo)
	query := fromRequestToQuery(req)
	result := mediator.Send(query)
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
