package core

import "context"

// Span represents an active tracing span with tag and lifecycle management.
type Span interface {
	SetTag(key string, value interface{})
	Finish()
}

// Tracer is an interface for distributed tracing that can be injected into the SDK.
// Implementations should start a new span derived from the provided context.
type Tracer interface {
	StartSpanFromContext(ctx context.Context, operationName, serviceName, resourceName string) Span
}
