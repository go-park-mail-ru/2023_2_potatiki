package metricsmw

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/metrics"
	"google.golang.org/grpc"
	"time"
)

type GrpcMiddleware struct {
	metrics metrics.MetricsGRPC
}

func NewGrpcMiddleware(metrics metrics.MetricsGRPC) *GrpcMiddleware {
	return &GrpcMiddleware{
		metrics: metrics,
	}
}

// Отзывы (гет, создать)
func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)

	m.metrics.IncreaseMetric("", info.FullMethod, tm.String())
	if err != nil {
		m.metrics.IncreaseErr(info.FullMethod)
	}
	m.metrics.IncreaseHits(info.FullMethod)
	//m.durations.With(prometheus.Labels{URL: info.FullMethod}).Observe(tm.Seconds())
	//
	//m.counter.With(prometheus.Labels{URL: info.FullMethod}).Inc()

	return h, err

}
