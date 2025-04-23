package models

import "github.com/jackc/pgx/v5/pgtype"

type CpuUsage struct {
	CPU      pgtype.Text
	AvgUsage pgtype.Float8
	MaxUsage pgtype.Float8
	MinUsage pgtype.Float8
}

type FullPrometheusData struct {
	Timestamp                             pgtype.Timestamp
	ProcessCpuUsage                       pgtype.Float8
	JvmMemoryMax                          pgtype.Int4
	ProcessRuntimeJvmMemoryUsage          pgtype.Int4
	ProcessRuntimeJvmThreadsCount         pgtype.Int2
	ProcessRuntimeJvmSystemCpuUtilization pgtype.Float8
	K8sPodName                            pgtype.Text
	K8sContainerName                      pgtype.Text
	K8sDeploymentName                     pgtype.Text
	OtlpExporterExported                  pgtype.Int2
}
