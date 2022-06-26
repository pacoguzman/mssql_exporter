package internal

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

func returnMetrics(db *sql.DB) []prometheus.Metric {
	var metrics []prometheus.Metric
	metrics = append(metrics, getCurrentUserSessions(db)...)
	metrics = append(metrics, getSuspendedSessions(db)...)
	return metrics
}

func returnMetric(name, desc, labelDesc, label string, value float64) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			name,
			desc,
			[]string{labelDesc}, nil),
		prometheus.GaugeValue,
		value,
		[]string{label}...,
	)
}