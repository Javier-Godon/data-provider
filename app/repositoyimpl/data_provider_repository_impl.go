package repositoyimpl

import (
	"context"
	"log"
	"time"

	"data-provider/framework"
	"data-provider/persistence/models"
)

type DataProviderRepositoryImpl struct{}

func New() *DataProviderRepositoryImpl {
	return &DataProviderRepositoryImpl{}
}

func (r *DataProviderRepositoryImpl) GetCpuSystemUsage(dateFrom int64, dateTo int64) ([]models.CpuUsage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := framework.DB.Query(ctx, `
		SELECT cpu, 
		       COALESCE(avg(usage_system), 0), 
		       COALESCE(max(usage_system), 0), 
		       COALESCE(min(usage_system), 0) 
		FROM cpu
		WHERE cpu IS NOT NULL
		GROUP BY cpu;
	`)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.CpuUsage
	for rows.Next() {
		var usage models.CpuUsage
		if err := rows.Scan(&usage.CPU, &usage.AvgUsage, &usage.MaxUsage, &usage.MinUsage); err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		results = append(results, usage)
	}

	return results, nil
}

func (r *DataProviderRepositoryImpl) GetCpuUserUsage(dateFrom int64, dateTo int64) ([]models.CpuUsage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := framework.DB.Query(ctx, `
		SELECT cpu, 
		       COALESCE(avg(usage_user), 0), 
		       COALESCE(max(usage_user), 0), 
		       COALESCE(min(usage_user), 0) 
		FROM cpu
		WHERE cpu IS NOT NULL
		GROUP BY cpu;
	`)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.CpuUsage
	for rows.Next() {
		var usage models.CpuUsage
		if err := rows.Scan(&usage.CPU, &usage.AvgUsage, &usage.MaxUsage, &usage.MinUsage); err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		results = append(results, usage)
	}

	return results, nil
}
