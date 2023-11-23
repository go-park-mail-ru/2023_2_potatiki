package metrcismw

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/metrics"
	"google.golang.org/grpc"
	"time"
)

type GrpcMiddleware struct {
	metrics metrics.MetricsGRPC // интерфейс
}

func NewGrpcMiddleware(metrics metrics.MetricsGRPC) *GrpcMiddleware {
	return &GrpcMiddleware{
		metrics: metrics,
	}
}

// конструктор
// Отзывы (гет, создать)
// TODO: засунуть сюда mt := NewMetrics
func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)

	//m.metric.With(prometheus.Labels{
	//	URL:         "",
	//	ServiceName: m.name,
	//	StatusCode:  "OK",
	//	Method:      info.FullMethod,
	//	FullTime:    tm.String(),
	//}).Inc()
	m.metrics.IncreaseHits("/path")
	//m.durations.With(prometheus.Labels{URL: info.FullMethod}).Observe(tm.Seconds())
	//
	//m.counter.With(prometheus.Labels{URL: info.FullMethod}).Inc()

	return h, err

}
