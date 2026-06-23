package core

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockSpan captures tags set during a request for test assertions.
type mockSpan struct {
	tags     map[string]interface{}
	finished bool
}

func (s *mockSpan) SetTag(key string, value interface{}) {
	s.tags[key] = value
}

func (s *mockSpan) Finish() {
	s.finished = true
}

type mockTracer struct {
	spans []*mockSpan
}

func newMockTracer() *mockTracer {
	return &mockTracer{}
}

func (m *mockTracer) StartSpanFromContext(_ context.Context, _, _, _ string) Span {
	s := &mockSpan{tags: make(map[string]interface{})}
	m.spans = append(m.spans, s)
	return s
}

// noOpHttpClient satisfies HttpClient without making real network calls.
type noOpHttpClient struct{}

func (c *noOpHttpClient) Get(_ interface{}, _ string, _ ...RequestOption) error { return nil }
func (c *noOpHttpClient) GetStream(_ BinaryResponse, _ string, _ ...RequestOption) error {
	return nil
}
func (c *noOpHttpClient) Post(_ interface{}, _ string, _ interface{}, _ ...RequestOption) error {
	return nil
}
func (c *noOpHttpClient) PostWithFiles(_ interface{}, _ string, _ interface{}, _ []File, _ ...RequestOption) error {
	return nil
}
func (c *noOpHttpClient) Request(_ interface{}, _ string, _ string, _ ...RequestOption) error {
	return nil
}
func (c *noOpHttpClient) RequestBinary(_ BinaryResponse, _ string, _ string, _ ...RequestOption) error {
	return nil
}
func (c *noOpHttpClient) SetRequester(_ HttpClientRequester) {}

type mockPayload struct {
	I int
	S string
}

const tracingTestOperationName = "test.GlideRequests"
const tracingTestServiceName = "test_service"

func newTestDecorator(ctx context.Context, tracer *mockTracer) HttpClient {
	return NewTracingHttpClientDecorator(ctx, &noOpHttpClient{}, tracer, TracingHttpClientConfig{
		OperationName: tracingTestOperationName,
		ServiceName:   tracingTestServiceName,
	})
}

func TestTracingHttpClientDecorator(t *testing.T) {
	const payloadJSON = "{\n  \"I\": 4,\n  \"S\": \"16\"\n}"

	tests := []struct {
		name         string
		act          func(httpClient HttpClient) error
		wantSpanTags map[string]interface{}
	}{
		{
			name: "GET request",
			act: func(httpClient HttpClient) error {
				return httpClient.Get(nil, "someUrl")
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "GET",
				"Endpoint":   "someUrl",
				"Response":   "(nil)",
			},
		},
		{
			name: "GET stream",
			act: func(httpClient HttpClient) error {
				return httpClient.GetStream(nil, "someUrl")
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "GET",
				"Endpoint":   "someUrl",
				"Response":   "(not-json-response)",
			},
		},
		{
			name: "POST request",
			act: func(httpClient HttpClient) error {
				return httpClient.Post(nil, "someUrl", mockPayload{I: 4, S: "16"})
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "POST",
				"Endpoint":   "someUrl",
				"Payload":    payloadJSON,
			},
		},
		{
			name: "POST without payload",
			act: func(httpClient HttpClient) error {
				return httpClient.Post(nil, "someUrl", nil)
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "POST",
				"Endpoint":   "someUrl",
			},
		},
		{
			name: "PostWithFiles",
			act: func(httpClient HttpClient) error {
				return httpClient.PostWithFiles(nil, "someUrl", mockPayload{I: 4, S: "16"}, nil)
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "POST",
				"Endpoint":   "someUrl",
				"Payload":    payloadJSON,
				"Response":   "(nil)",
			},
		},
		{
			name: "Request (PUT)",
			act: func(httpClient HttpClient) error {
				return httpClient.Request(nil, "PUT", "someUrl")
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "PUT",
				"Endpoint":   "someUrl",
			},
		},
		{
			name: "RequestBinary",
			act: func(httpClient HttpClient) error {
				return httpClient.RequestBinary(nil, "GET", "someUrl")
			},
			wantSpanTags: map[string]interface{}{
				"HttpMethod": "GET",
				"Endpoint":   "someUrl",
				"Response":   "(not-json-response)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := newMockTracer()
			httpClient := newTestDecorator(context.Background(), mt)

			err := tt.act(httpClient)
			assert.NoError(t, err)

			assert.Len(t, mt.spans, 1)
			span := mt.spans[0]
			assert.True(t, span.finished, "span should be finished")

			_, hasDuration := span.tags["Request Duration"]
			assert.True(t, hasDuration, "Request Duration tag should be set")

			for tag, expectedValue := range tt.wantSpanTags {
				assert.Equal(t, expectedValue, span.tags[tag], "tag %q mismatch", tag)
			}

			_, hasPayload := span.tags["Payload"]
			if _, wantsPayload := tt.wantSpanTags["Payload"]; !wantsPayload {
				assert.False(t, hasPayload, "Payload tag should not be set when no payload")
			}
		})
	}
}

func TestTracingHttpClientDecoratorSetsErrorTag(t *testing.T) {
	requestErr := errors.New("connection refused")
	errClient := &errorHttpClient{err: requestErr}

	mt := newMockTracer()
	httpClient := NewTracingHttpClientDecorator(context.Background(), errClient, mt, TracingHttpClientConfig{
		OperationName: tracingTestOperationName,
		ServiceName:   tracingTestServiceName,
	})

	err := httpClient.Get(nil, "someUrl")

	assert.Equal(t, requestErr, err)
	assert.Len(t, mt.spans, 1)
	assert.Equal(t, requestErr, mt.spans[0].tags["error"])
}

type errorHttpClient struct {
	err error
}

func (c *errorHttpClient) Get(_ interface{}, _ string, _ ...RequestOption) error        { return c.err }
func (c *errorHttpClient) GetStream(_ BinaryResponse, _ string, _ ...RequestOption) error { return c.err }
func (c *errorHttpClient) Post(_ interface{}, _ string, _ interface{}, _ ...RequestOption) error {
	return c.err
}
func (c *errorHttpClient) PostWithFiles(_ interface{}, _ string, _ interface{}, _ []File, _ ...RequestOption) error {
	return c.err
}
func (c *errorHttpClient) Request(_ interface{}, _ string, _ string, _ ...RequestOption) error {
	return c.err
}
func (c *errorHttpClient) RequestBinary(_ BinaryResponse, _ string, _ string, _ ...RequestOption) error {
	return c.err
}
func (c *errorHttpClient) SetRequester(_ HttpClientRequester) {}

func TestTracingHttpClientDecoratorForwardsSpanMetadata(t *testing.T) {
	mt := newMockTracer()

	capturedArgs := struct {
		operationName string
		serviceName   string
		resourceName  string
	}{}

	capturingTracer := &capturingTracerImpl{
		onStart: func(opName, svcName, resName string) {
			capturedArgs.operationName = opName
			capturedArgs.serviceName = svcName
			capturedArgs.resourceName = resName
		},
		delegate: mt,
	}

	httpClient := NewTracingHttpClientDecorator(context.Background(), &noOpHttpClient{}, capturingTracer, TracingHttpClientConfig{
		OperationName: tracingTestOperationName,
		ServiceName:   tracingTestServiceName,
	})

	_ = httpClient.Get(nil, "https://api.example.com/transactions/123/documents")

	assert.Equal(t, tracingTestOperationName, capturedArgs.operationName)
	assert.Equal(t, tracingTestServiceName, capturedArgs.serviceName)
	assert.Equal(t, "/transactions/{num}/documents", capturedArgs.resourceName)
}

type capturingTracerImpl struct {
	onStart  func(operationName, serviceName, resourceName string)
	delegate *mockTracer
}

func (c *capturingTracerImpl) StartSpanFromContext(ctx context.Context, operationName, serviceName, resourceName string) Span {
	c.onStart(operationName, serviceName, resourceName)
	return c.delegate.StartSpanFromContext(ctx, operationName, serviceName, resourceName)
}

func TestNormalizeURLForTracing(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "URL with single numeric ID",
			input:    "https://api.example.com/users/123",
			expected: "/users/{num}",
		},
		{
			name:     "URL with multiple numeric IDs",
			input:    "https://api.example.com/users/123/posts/456",
			expected: "/users/{num}/posts/{num}",
		},
		{
			name:     "URL with numeric ID in middle",
			input:    "https://api.example.com/users/123/profile",
			expected: "/users/{num}/profile",
		},
		{
			name:     "URL with no numeric IDs",
			input:    "https://api.example.com/users/profile",
			expected: "/users/profile",
		},
		{
			name:     "URL with query parameters",
			input:    "https://api.example.com/users/123?param=value",
			expected: "/users/{num}",
		},
		{
			name:     "URL with fragment",
			input:    "https://api.example.com/users/123#section",
			expected: "/users/{num}",
		},
		{
			name:     "URL with numeric ID and trailing slash",
			input:    "https://api.example.com/users/123/",
			expected: "/users/{num}/",
		},
		{
			name:     "URL with only numeric segments",
			input:    "https://api.example.com/123/456/789",
			expected: "/{num}/{num}/{num}",
		},
		{
			name:     "URL with port and numeric ID",
			input:    "https://api.example.com:8080/users/123",
			expected: "/users/{num}",
		},
		{
			name:     "Complex URL with multiple numeric IDs and parameters",
			input:    "https://api.example.com/v1/users/123/projects/456/tasks/789?limit=10&offset=20",
			expected: "/v1/users/{num}/projects/{num}/tasks/{num}",
		},
		{
			name:     "URL with negative numbers (should not be replaced)",
			input:    "https://api.example.com/values/-123",
			expected: "/values/-123",
		},
		{
			name:     "HTTP URL with numeric ID",
			input:    "http://api.example.com/users/123",
			expected: "/users/{num}",
		},
		{
			name:     "Relative URL with numeric ID",
			input:    "/users/123/profile",
			expected: "/users/{num}/profile",
		},
		{
			name:     "URL with zero as ID",
			input:    "https://api.example.com/users/0",
			expected: "/users/{num}",
		},
		{
			name:     "URL with large numeric ID",
			input:    "https://api.example.com/users/9876543210",
			expected: "/users/{num}",
		},
		{
			name:     "Invalid URL (unparseable) returns raw input",
			input:    "%invalid-percent-encoding",
			expected: "%invalid-percent-encoding",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeURLForTracing(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
