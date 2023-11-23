package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type MetricsGRPC interface {
	IncreaseHits(string)
	IncreaseErr(string, string)
	AddDurationToHistogram(string, time.Duration)
	AddDurationToSummary(string, string, time.Duration)
}

type MetricGRPC struct {
	totalHits         *prometheus.CounterVec
	totalErrors       *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
	durationSummary   *prometheus.SummaryVec
	name              string
}

func NewGRPCMetrics() *MetricGRPC {
	labelHits := []string{"path"}
	totalHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "number_of_all_requests",
	}, labelHits)
	prometheus.MustRegister(totalHits)

	labelErrors := []string{"status_code", "path"}
	totalErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_errors_total",
		Help: "number_of_all_errors",
	}, labelErrors)
	prometheus.MustRegister(totalErrors)

	labelHistogram := []string{"path"}
	durationHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "durations_stats_histogram",
		Help: "durations_stats_histogram",
		// Buckets: prometheus.LinearBuckets(0, 1, 10),
	}, labelHistogram)
	prometheus.MustRegister(durationHistogram)

	labelSummary := []string{"status_code", "path"}
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
	}
}

func (m *MetricGRPC) IncreaseHits(path string) {
	m.totalHits.WithLabelValues(path).Inc()
}

func (m *MetricGRPC) IncreaseErr(statusCode, path string) {
	// m.totalHits.WithLabelValues(path).Inc()

	m.totalErrors.WithLabelValues(statusCode, path).Inc()
}

func (m *MetricGRPC) AddDurationToHistogram(path string, duration time.Duration) {
	m.durationHistogram.WithLabelValues(path).Observe(duration.Seconds())
}

func (m *MetricGRPC) AddDurationToSummary(statusCode string, path string, duration time.Duration) {
	m.durationSummary.WithLabelValues(statusCode, path).Observe(duration.Seconds())
}
