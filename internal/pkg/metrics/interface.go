package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Increase()
	IncreaseWithErr(string, string)
}

type Metric struct {
	total       prometheus.Counter
	totalErrors *prometheus.CounterVec
}

func NewMetrics() *Metric {
	total := prometheus.NewCounter(prometheus.CounterOpts{Name: "HTTP_Requests_Total"})
	prometheus.MustRegister(total)

	label := []string{"status_code", "path"}
	totalErrors := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "HTTP_Requests_Total_Errors"}, label)
	prometheus.MustRegister(totalErrors)

	return &Metric{
		total:       total,
		totalErrors: totalErrors,
	}
}

func (m *Metric) Increase() {
	m.total.Inc()
}

func (m *Metric) IncreaseWithErr(statusCode, path string) {
	m.total.Inc()

	m.totalErrors.WithLabelValues(statusCode, path).Inc()
}
