package metricsmw

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/metrics"
	"google.golang.org/grpc"
)

type GrpcMiddleware struct {
	metrics metrics.MetricerGRPC
}

func NewGrpcMiddleware(metrics metrics.MetricerGRPC) *GrpcMiddleware {
	return &GrpcMiddleware{
		metrics: metrics,
	}
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)

	if err != nil {
		m.metrics.IncreaseErr(info.FullMethod)
	}

	m.metrics.AddDurationToHistogram(info.FullMethod, tm)

	m.metrics.AddDurationToSummary("", info.FullMethod, tm)

	m.metrics.IncreaseHits(info.FullMethod)

	return h, err

}
