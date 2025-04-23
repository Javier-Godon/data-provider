package mediator

import (
	"log"

	"github.com/Javier-Godon/data-provider/framework"
	"github.com/Javier-Godon/data-provider/repositoryimpl"
	"github.com/Javier-Godon/data-provider/usecases/prometheus/get_full_prometheus_data"
)

func init() {
	err := framework.Register(
		get_full_prometheus_data.NewGetFullPrometheusDataHandler(&repositoryimpl.DataProviderRepositoryImpl{}))
	if err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}
}

func Send(query get_full_prometheus_data.GetFullPrometheusDataQuery) get_full_prometheus_data.GetFullPrometheusDataResult {
	result, err := framework.Send[get_full_prometheus_data.GetFullPrometheusDataQuery, get_full_prometheus_data.GetFullPrometheusDataResult](query)
	if err != nil {
		log.Fatalf("Could not execute command: %v", err)
	}
	return result
}
