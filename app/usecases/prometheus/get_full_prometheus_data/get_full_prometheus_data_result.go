package get_full_prometheus_data

type GetFullPrometheusDataResult struct {
	PrometheusData []PrometheusDataResult
}
type PrometheusDataResult struct {
	Timestamp                             int64
	ProcessCpuUsage                       float64
	JvmMemoryMax                          float64
	ProcessRuntimeJvmMemoryUsage          float64
	ProcessRuntimeJvmThreadsCount         float64
	ProcessRuntimeJvmSystemCpuUtilization float64
	K8sPodName                            string
	K8sContainerName                      string
	K8sDeploymentName                     string
	OtlpExporterExported                  int8
}
