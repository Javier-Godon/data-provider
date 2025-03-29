package repositoryimpl

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Javier-Godon/data-provider/framework"
	"github.com/Javier-Godon/data-provider/persistence/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataProviderRepositoryImpl struct {
	db *pgxpool.Pool
}

func New() *DataProviderRepositoryImpl {
	return &DataProviderRepositoryImpl{
		db: framework.DB,
	}
}

func (r DataProviderRepositoryImpl) GetCpuSystemUsage(dateFrom int64, dateTo int64) ([]models.CpuUsage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dateFromMicro := dateFrom * 1_000_000
	dateToMicro := dateTo * 1_000_000

	if r.db == nil {
		log.Println("Database connection is nil")
		return nil, errors.New("database connection is not initialized")
	}

	rows, err := r.db.Query(ctx, `
	SELECT cpu, 
	       COALESCE(avg(usage_system), 0) AS avg_usage, 
	       COALESCE(max(usage_system), 0) AS max_usage, 
	       COALESCE(min(usage_system), 0) AS min_usage 
	FROM cpu
	WHERE cpu IS NOT NULL
	  AND timestamp >= CAST($1 AS TIMESTAMP)
	  AND timestamp <= CAST($2 AS TIMESTAMP)
	GROUP BY cpu;
`, dateFromMicro, dateToMicro)
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

	dateFromMicro := dateFrom * 1_000_000
	dateToMicro := dateTo * 1_000_000

	rows, err := r.db.Query(ctx, `
	SELECT cpu, 
	       COALESCE(avg(usage_user), 0) AS avg_usage, 
	       COALESCE(max(usage_user), 0) AS max_usage, 
	       COALESCE(min(usage_user), 0) AS min_usage 
	FROM cpu
	WHERE cpu IS NOT NULL
	  AND timestamp >= CAST($1 AS TIMESTAMP)
	  AND timestamp <= CAST($2 AS TIMESTAMP)
	GROUP BY cpu;
`, dateFromMicro, dateToMicro)
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
