package persistence

import "data-provider/persistence/models"

type Repository interface {
	GetCpuSystemUsage(dateFrom int64, dateTo int64) ([]models.CpuUsage, error)
	GetCpuUserUsage(dateFrom int64, dateTo int64) ([]models.CpuUsage, error)
}
