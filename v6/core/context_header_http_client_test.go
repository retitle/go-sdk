package core

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// capturingHttpClient records the RequestOptions of every call so tests can
// assert which headers were injected.
type capturingHttpClient struct {
	lastOpts RequestOptionsImpl
}

func (c *capturingHttpClient) capture(opts []RequestOption) {
	c.lastOpts = RequestOptionsImpl{}
	for _, opt := range opts {
		opt(&c.lastOpts)
	}
}

func (c *capturingHttpClient) Get(_ interface{}, _ string, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) GetStream(_ BinaryResponse, _ string, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) Post(_ interface{}, _ string, _ interface{}, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) PostWithFiles(_ interface{}, _ string, _ interface{}, _ []File, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) Request(_ interface{}, _ string, _ string, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) RequestBinary(_ BinaryResponse, _ string, _ string, opts ...RequestOption) error {
	c.capture(opts)
	return nil
}
func (c *capturingHttpClient) SetRequester(_ HttpClientRequester) {}

// traceExtractor is a simple HeaderExtractor that returns a fixed trace ID header.
func traceExtractor(traceID string) HeaderExtractor {
	return func(_ context.Context) map[string]string {
		if traceID == "" {
			return nil
		}
		return map[string]string{"X-Compass-Trace-Id": traceID}
	}
}

func TestContextHeaderHttpClient_InjectsHeadersOnAllMethods(t *testing.T) {
	const traceID = "crafting-test-trace-abc123"

	methods := []struct {
		name string
		act  func(client HttpClient)
	}{
		{"Get", func(c HttpClient) { _ = c.Get(nil, "http://example.com") }},
		{"GetStream", func(c HttpClient) { _ = c.GetStream(nil, "http://example.com") }},
		{"Post", func(c HttpClient) { _ = c.Post(nil, "http://example.com", nil) }},
		{"PostWithFiles", func(c HttpClient) { _ = c.PostWithFiles(nil, "http://example.com", nil, nil) }},
		{"Request", func(c HttpClient) { _ = c.Request(nil, "PUT", "http://example.com") }},
		{"RequestBinary", func(c HttpClient) { _ = c.RequestBinary(nil, "GET", "http://example.com") }},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			inner := &capturingHttpClient{}
			client := NewContextHeaderHttpClient(context.Background(), inner, traceExtractor(traceID))

			tt.act(client)

			assert.Equal(t, http.Header{"X-Compass-Trace-Id": []string{traceID}}, inner.lastOpts.GetHeaders())
		})
	}
}

func TestContextHeaderHttpClient_NoHeadersWhenExtractorReturnsEmpty(t *testing.T) {
	emptyExtractor := func(_ context.Context) map[string]string { return nil }

	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(context.Background(), inner, emptyExtractor)

	_ = client.Get(nil, "http://example.com")

	// No headers should have been injected.
	assert.Empty(t, inner.lastOpts.GetHeaders())
}

func TestContextHeaderHttpClient_MultipleExtractorsMerged(t *testing.T) {
	extA := func(_ context.Context) map[string]string {
		return map[string]string{"X-Compass-Trace-Id": "trace-123"}
	}
	extB := func(_ context.Context) map[string]string {
		return map[string]string{"X-Compass-Service-Id": "svc-abc"}
	}

	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(context.Background(), inner, extA, extB)

	_ = client.Get(nil, "http://example.com")

	headers := inner.lastOpts.GetHeaders()
	assert.Equal(t, "trace-123", headers.Get("X-Compass-Trace-Id"))
	assert.Equal(t, "svc-abc", headers.Get("X-Compass-Service-Id"))
}

func TestContextHeaderHttpClient_CallerOptOverridesExtractedHeader(t *testing.T) {
	// The extractor supplies a trace ID, but the caller passes its own value
	// for the same header. The caller's value should win because caller opts
	// are appended after the extracted opts (prepend puts extracted opts first).
	extractor := func(_ context.Context) map[string]string {
		return map[string]string{"X-Compass-Trace-Id": "extracted-value"}
	}

	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(context.Background(), inner, extractor)

	_ = client.Get(nil, "http://example.com", WithHeader("X-Compass-Trace-Id", "caller-value"))

	// withHeaders merges and the last write wins; caller opts come after extracted
	// opts so the caller's value takes precedence.
	assert.Equal(t, "caller-value", inner.lastOpts.GetHeaders().Get("X-Compass-Trace-Id"))
}

func TestContextHeaderHttpClient_ContextPassedToExtractor(t *testing.T) {
	type ctxKey struct{}
	ctx := context.WithValue(context.Background(), ctxKey{}, "my-service")

	var capturedCtx context.Context
	extractor := func(c context.Context) map[string]string {
		capturedCtx = c
		return nil
	}

	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(ctx, inner, extractor)
	_ = client.Get(nil, "http://example.com")

	assert.Equal(t, "my-service", capturedCtx.Value(ctxKey{}))
}

func TestContextHeaderHttpClient_NoExtractors(t *testing.T) {
	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(context.Background(), inner)

	_ = client.Get(nil, "http://example.com")

	assert.Empty(t, inner.lastOpts.GetHeaders())
}

func TestContextHeaderHttpClient_SetRequesterDelegates(t *testing.T) {
	inner := &capturingHttpClient{}
	client := NewContextHeaderHttpClient(context.Background(), inner, traceExtractor("t"))

	// Should not panic — SetRequester is delegated to the inner client.
	client.SetRequester(nil)
}

func TestContextHeaderHttpClient_ComposedWithTracingDecorator(t *testing.T) {
	// Verify the decorator can be stacked inside TracingHttpClientDecorator,
	// which is the expected usage in production client factories.
	const traceID = "crafting-composed-trace"

	inner := &capturingHttpClient{}
	withHeaders := NewContextHeaderHttpClient(context.Background(), inner, traceExtractor(traceID))

	mt := newMockTracer()
	composed := NewTracingHttpClientDecorator(
		context.Background(),
		withHeaders,
		mt,
		TracingHttpClientConfig{OperationName: "test.op", ServiceName: "test_svc"},
	)

	_ = composed.Get(nil, "http://example.com")

	// Tracing decorator created a span.
	assert.Len(t, mt.spans, 1)
	// Header was propagated through to the inner client.
	assert.Equal(t, traceID, inner.lastOpts.GetHeaders().Get("X-Compass-Trace-Id"))
}
