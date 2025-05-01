package get_full_prometheus_data

import (
	"github.com/Javier-Godon/data-provider/persistence"
	"github.com/Javier-Godon/data-provider/repositoryimpl"
)


type GetFullPrometheusDataHandler struct {
	Repository persistence.Repository
}

func NewGetFullPrometheusDataHandler(repository persistence.Repository) *GetFullPrometheusDataHandler {
	return &GetFullPrometheusDataHandler{
		Repository: repository,
	}
}

func (handler GetFullPrometheusDataHandler) Handle(query GetFullPrometheusDataQuery) (GetFullPrometheusDataResult, error) {

	repository := repositoryimpl.New()
	prometheusDataResultset, err := repository.GetFullPrometheusData(query.DateFrom, query.DateTo)
	if err != nil {
		return GetFullPrometheusDataResult{}, err
	}

	var prometheusData []PrometheusDataResult
	for _, data := range prometheusDataResultset {
		prometheusData = append(prometheusData, PrometheusDataResult{
			Timestamp:                             data.Timestamp.Time.Unix(),
			ProcessCpuUsage:                       data.ProcessCpuUsage.Float64,
			JvmMemoryMax:                          float64(data.JvmMemoryMax.Float32),
			ProcessRuntimeJvmMemoryUsage:          data.ProcessRuntimeJvmMemoryUsage.Float64,
			ProcessRuntimeJvmThreadsCount:         data.ProcessRuntimeJvmThreadsCount.Float64,
			ProcessRuntimeJvmSystemCpuUtilization: data.ProcessRuntimeJvmSystemCpuUtilization.Float64,
			K8sPodName:                            data.K8sPodName.String,
			K8sContainerName:                      data.K8sContainerName.String,
			K8sDeploymentName:                     data.K8sDeploymentName.String,
			OtlpExporterExported:                  int8(data.OtlpExporterExported.Int16),
		})
	}

	return GetFullPrometheusDataResult{PrometheusData: prometheusData}, nil
}
