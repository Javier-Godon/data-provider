package get_cpu_system_usage

import (
	"github.com/Javier-Godon/data-provider/persistence"
	"github.com/Javier-Godon/data-provider/repositoryimpl"
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

	repository := repositoryimpl.New()
	cpuUsages, err := repository.GetCpuSystemUsage(query.DateFrom, query.DateTo)
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
