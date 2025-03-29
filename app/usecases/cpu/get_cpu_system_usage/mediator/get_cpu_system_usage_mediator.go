package mediator

import (
	"log"

	"github.com/Javier-Godon/data-provider/framework"
	"github.com/Javier-Godon/data-provider/repositoryimpl"
	"github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_system_usage"
)

func init() {
	err := framework.Register(
		get_cpu_system_usage.NewGetCpuSystemUsageHandler(&repositoryimpl.DataProviderRepositoryImpl{}))
	if err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}
}

func Send(query get_cpu_system_usage.GetCpuSystemUsageQuery) get_cpu_system_usage.GetCpuSystemUsageResult {
	result, err := framework.Send[get_cpu_system_usage.GetCpuSystemUsageQuery, get_cpu_system_usage.GetCpuSystemUsageResult](query)
	if err != nil {
		log.Fatalf("Could not execute command: %v", err)
	}
	return result
}
