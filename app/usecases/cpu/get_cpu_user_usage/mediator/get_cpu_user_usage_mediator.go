package mediator

import (
	"log"

	"github.com/Javier-Godon/data-provider/framework"
	"github.com/Javier-Godon/data-provider/repositoryimpl"
	"github.com/Javier-Godon/data-provider/usecases/cpu/get_cpu_user_usage"
)

func init() {
	err := framework.Register(
		get_cpu_user_usage.NewGetCpuUserUsageHandler(&repositoryimpl.DataProviderRepositoryImpl{}),
	)
	if err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}
}

func Send(command get_cpu_user_usage.GetCpuUserUsageQuery) get_cpu_user_usage.GetCpuUserUsageResult {
	result, err := framework.Send[get_cpu_user_usage.GetCpuUserUsageQuery, get_cpu_user_usage.GetCpuUserUsageResult](command)
	if err != nil {
		log.Fatalf("Could not execute command: %v", err)
	}
	return result
}
