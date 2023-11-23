package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	ServiceAuthName = "auth"
)

type MetricsGRPC interface {
	IncreaseMetric(string, string, string)
	IncreaseHits(string)
	IncreaseErr(string)
	AddDurationToHistogram(string, time.Duration)
	AddDurationToSummary(string, string, time.Duration)
}

type MetricGRPC struct {
	metric            *prometheus.GaugeVec
	totalHits         *prometheus.CounterVec
	totalErrors       *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
	durationSummary   *prometheus.SummaryVec
	name              string
}

func NewGRPCMetrics(serverName string) *MetricGRPC {
	labelGauge := []string{"path", "service_name", "method", "full_time"}
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: serverName,
		Help: fmt.Sprintf("SLO for service %s", serverName),
	}, labelGauge)
	prometheus.MustRegister(gauge)

	labelHits := []string{"path", "service_name"}
	totalHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "number_of_all_requests",
	}, labelHits)
	prometheus.MustRegister(totalHits)

	labelErrors := []string{"path", "service_name"}
	totalErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_errors_total",
		Help: "number_of_all_errors",
	}, labelErrors)
	prometheus.MustRegister(totalErrors)

	labelHistogram := []string{"path", "service_name"}
	durationHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "durations_stats_histogram",
		Help: "durations_stats_histogram",
		// Buckets: prometheus.LinearBuckets(0, 1, 10),
	}, labelHistogram)
	prometheus.MustRegister(durationHistogram)

	labelSummary := []string{"status_code", "path", "service_name"}
	durationSummary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "durations_stats_summary",
		Help: "durations_stats_summary",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.8:  0.1,
			0.9:  0.1,
			0.95: 0.1,
			0.99: 0.1,
			1:    0.1,
		}}, labelSummary)
	prometheus.MustRegister(durationSummary)

	return &MetricGRPC{
		totalHits:         totalHits,
		totalErrors:       totalErrors,
		durationHistogram: durationHistogram,
		durationSummary:   durationSummary,
		name:              serverName,
	}
}

func (m *MetricGRPC) IncreaseMetric(path, method, duration string) {
	m.metric.WithLabelValues(path, m.name, method, duration).Inc()
}
func (m *MetricGRPC) IncreaseHits(path string) {
	m.totalHits.WithLabelValues(path, m.name).Inc()
}

func (m *MetricGRPC) IncreaseErr(path string) {
	// m.totalHits.WithLabelValues(path).Inc()

	m.totalErrors.WithLabelValues(path, m.name).Inc()
}

func (m *MetricGRPC) AddDurationToHistogram(path string, duration time.Duration) {
	m.durationHistogram.WithLabelValues(path, m.name).Observe(duration.Seconds())
}

func (m *MetricGRPC) AddDurationToSummary(statusCode string, path string, duration time.Duration) {
	m.durationSummary.WithLabelValues(statusCode, path, m.name).Observe(duration.Seconds())
}
