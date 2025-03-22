package get_cpu_system_usage

import (
	"data-provider/persistence"
)

type GetCpuSystemUsageHandler struct {
	Repository persistence.Repository
}

func NewGetCpuSystemUsageHandler(repository persistence.Repository) *GetCpuSystemUsageHandler {
	return &GetCpuSystemUsageHandler{
		Repository: repository,
	}
}

func (handler GetCpuSystemUsageHandler) Handle(query GetCpuSystemUsageQuery) (GetCpuSystemUsageResult, error) {
	cpuUsages, err := handler.Repository.GetCpuSystemUsage(query.DateFrom, query.DateTo)
	if err != nil {
		return GetCpuSystemUsageResult{}, err
	}

	var usages []CpuUsageResult
	for _, usage := range cpuUsages {
		usages = append(usages, CpuUsageResult{
			Cpu:      usage.CPU.String,
			AvgUsage: usage.AvgUsage.Float64,
			MaxUsage: usage.MaxUsage.Float64,
			MinUsage: usage.MinUsage.Float64,
		})
	}

	return GetCpuSystemUsageResult{Usages: usages}, nil
}
