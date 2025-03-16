package mediator

import (
	"data-provider/app/persistence/adapter"
	"data-provider/usecases/get_cpu_system_usage"
	"data-provider/framework"
	"log"
)

func init() {
	err := framework.Register(get_cpu_system_usage.NewGetCpuSystemUsageHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(query get_cpu_system_usage.GetCpuSystemUsageQuery) get_cpu_system_usage.GetCategoryByIdResult {
	GetCategoryByIdResult, err := framework.Send[get_cpu_system_usage.GetCategoryByIdCommand, get_cpu_system_usage.GetCategoryByIdResult](query)
	if err != nil {
		log.Fatalf("Could not execute: %v", query)
	}
	return GetCategoryByIdResult
}
