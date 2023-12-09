package metricsmw

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/metrics"
	"google.golang.org/grpc"
)

var (
	ClientError = status.Error(codes.InvalidArgument, "invalid ID, fail to cast uuid")
	ServerError = status.Error(codes.Internal, "internal server error")
)

type GrpcMiddleware struct {
	metrics metrics.MetricsGRPC
}

func NewGrpcMiddleware(metrics metrics.MetricsGRPC) *GrpcMiddleware {
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
		if errors.Is(err, ClientError) {
			m.metrics.IncreaseErr("400", info.FullMethod)
			m.metrics.AddDurationToSummary("400", info.FullMethod, tm)
		}
		if errors.Is(err, ServerError) {
			m.metrics.IncreaseErr("429", info.FullMethod)
			m.metrics.AddDurationToSummary("429", info.FullMethod, tm)
		}
	} else {
		m.metrics.AddDurationToSummary("200", info.FullMethod, tm)
	}

	m.metrics.AddDurationToHistogram(info.FullMethod, tm)

	m.metrics.IncreaseHits(info.FullMethod)

	return h, err

}
