package get_cpu_user_usage

import (
	"data-provider/persistence"
)

type GetCpuUserUsageHandler struct {
	Repository persistence.Repository
}

func NewGetCpuUserUsageHandler(repository persistence.Repository) *GetCpuUserUsageHandler {
	return &GetCpuUserUsageHandler{
		Repository: repository,
	}
}

func (handler GetCpuUserUsageHandler) Handle(query GetCpuUserUsageQuery) (GetCpuUserUsageResult, error) {
	cpuUsages, err := handler.Repository.GetCpuUserUsage(query.DateFrom, query.DateTo)
	if err != nil {
		return GetCpuUserUsageResult{}, err
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

	return GetCpuUserUsageResult{Usages: usages}, nil
}
