package get_cpu_user_usage

type GetCpuUserUsageResult struct {
	Usages []CpuUsageResult
}
type CpuUsageResult struct {
	Cpu      string
	AvgUsage float64
	MaxUsage float64
	MinUsage float64
}
