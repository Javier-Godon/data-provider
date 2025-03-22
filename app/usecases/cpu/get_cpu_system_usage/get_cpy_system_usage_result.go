package get_cpu_system_usage

type GetCpuSystemUsageResult struct {
	Usages []CpuUsageResult
}
type CpuUsageResult struct {
	Cpu      string
	AvgUsage float64
	MaxUsage float64
	MinUsage float64
}
