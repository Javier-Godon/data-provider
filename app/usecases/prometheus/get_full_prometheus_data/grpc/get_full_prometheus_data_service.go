package grpc

import (
	"context"
	"log"

	pb "github.com/Javier-Godon/data-provider/proto/get_full_prometheus_data"
	"github.com/Javier-Godon/data-provider/usecases/prometheus/get_full_prometheus_data"
	"github.com/Javier-Godon/data-provider/usecases/prometheus/get_full_prometheus_data/mediator"
)

type Service struct {
	pb.UnimplementedGetFullPrometheusDataServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetFullPrometheusData(ctx context.Context, req *pb.GetFullPrometheusDataRequest) (*pb.GetFullPrometheusDataResponse, error) {
	log.Printf("GetFullPrometheusData called with dateFrom=%d, dateTo=%d", req.DateFrom, req.DateTo)
	query := fromRequestToQuery(req)
	result := mediator.Send(query)
	return fromResultToResponse(result), nil
}

func fromRequestToQuery(req *pb.GetFullPrometheusDataRequest) get_full_prometheus_data.GetFullPrometheusDataQuery {
	return get_full_prometheus_data.GetFullPrometheusDataQuery{
		DateFrom: req.DateFrom,
		DateTo:   req.DateTo,
	}
}

func fromResultToResponse(result get_full_prometheus_data.GetFullPrometheusDataResult) *pb.GetFullPrometheusDataResponse {
	var prometheusData []*pb.FullPrometheusData
	for _, data := range result.PrometheusData {
		prometheusData = append(prometheusData, &pb.FullPrometheusData{
			Timestamp:                             data.Timestamp,
			ProcessCpuUsage:                       data.ProcessCpuUsage,
			JvmMemoryMax:                          data.JvmMemoryMax,
			ProcessRuntimeJvmMemoryUsage:          data.ProcessRuntimeJvmMemoryUsage,
			ProcessRuntimeJvmThreadsCount:         float64(data.ProcessRuntimeJvmThreadsCount),
			ProcessRuntimeJvmSystemCpuUtilization: float64(data.ProcessRuntimeJvmSystemCpuUtilization),
			K8SPodName:                            data.K8sPodName,
			K8SContainerName:                      data.K8sContainerName,
			K8SDeploymentName:                     data.K8sDeploymentName,
			OtlpExporterExported:                  int32(data.OtlpExporterExported),
		})
	}

	return &pb.GetFullPrometheusDataResponse{
		PrometheusData: prometheusData,
	}
}
