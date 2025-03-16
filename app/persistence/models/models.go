package models

import "github.com/jackc/pgx/v5/pgtype"

type CpuUsage struct {
	CPU      pgtype.Text
	AvgUsage pgtype.Float8
	MaxUsage pgtype.Float8
	MinUsage pgtype.Float8
}
