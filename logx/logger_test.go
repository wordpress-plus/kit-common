package logx

import (
	"context"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTraceLog(t *testing.T) {
	tracer := sdktrace.NewTracerProvider().Tracer("test")
	ctx, span := tracer.Start(
		context.Background(),
		"foo",
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
		oteltrace.WithAttributes(semconv.HTTPClientAttributesFromHTTPRequest(httptest.NewRequest(http.MethodGet, "/", nil))...),
	)
	defer span.End()

	Logger.WithContext(ctx)

	Logger.Warna("sasa")
}
