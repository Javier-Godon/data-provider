package get_full_prometheus_data

type GetFullPrometheusDataResult struct {
	PrometheusData []PrometheusDataResult
}
type PrometheusDataResult struct {
	Timestamp                             int64
	ProcessCpuUsage                       float64
	JvmMemoryMax                          int32
	ProcessRuntimeJvmMemoryUsage          int32
	ProcessRuntimeJvmThreadsCount         int8
	ProcessRuntimeJvmSystemCpuUtilization int16
	K8sPodName                            string
	K8sContainerName                      string
	K8sDeploymentName                     string
	OtlpExporterExported                  int8
}
