package mediator

import (
	"log"

	"data-provider/framework"
	"data-provider/repositoyimpl"
	"data-provider/usecases/cpu/get_cpu_system_usage"
)

func init() {
	err := framework.Register(
		get_cpu_system_usage.NewGetCpuSystemUsageHandler(repositoyimpl.New()),
	)
	if err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}
}

func Send(command get_cpu_system_usage.GetCpuSystemUsageQuery) get_cpu_system_usage.GetCpuSystemUsageResult {
	result, err := framework.Send[get_cpu_system_usage.GetCpuSystemUsageQuery, get_cpu_system_usage.GetCpuSystemUsageResult](command)
	if err != nil {
		log.Fatalf("Could not execute command: %v", err)
	}
	return result
}
