package core

import (
	"context"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"regexp"
	"time"
)

// TracingHttpClientConfig holds the tracing metadata for a decorator instance.
type TracingHttpClientConfig struct {
	OperationName string
	ServiceName   string
}

type tracingHttpClientDecorator struct {
	httpClient    HttpClient
	ctx           context.Context
	tracer        Tracer
	operationName string
	serviceName   string
}

var _ HttpClient = (*tracingHttpClientDecorator)(nil)

// NewTracingHttpClientDecorator wraps httpClient with distributed tracing via the provided Tracer.
// Each request starts a span tagged with the HTTP method, endpoint, payload, response, and duration.
//
// ctx is captured at construction time and used as the parent for every span. This means the
// decorator must be constructed fresh for each request context scope — do not construct it once
// and reuse it across requests with different contexts, or all spans will be rooted under the
// trace that was active at construction time rather than the trace of each individual request.
func NewTracingHttpClientDecorator(ctx context.Context, httpClient HttpClient, tracer Tracer, config TracingHttpClientConfig) HttpClient {
	return &tracingHttpClientDecorator{
		httpClient:    httpClient,
		ctx:           ctx,
		tracer:        tracer,
		operationName: config.OperationName,
		serviceName:   config.ServiceName,
	}
}

func (c *tracingHttpClientDecorator) Get(res interface{}, url string, opts ...RequestOption) error {
	return c.observeRequest("GET", url, opts, nil, true, func() (interface{}, error) {
		err := c.httpClient.Get(res, url, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) GetStream(res BinaryResponse, url string, opts ...RequestOption) error {
	return c.observeRequest("GET", url, opts, nil, false, func() (interface{}, error) {
		err := c.httpClient.GetStream(res, url, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) Post(res interface{}, url string, payload interface{}, opts ...RequestOption) error {
	return c.observeRequest("POST", url, opts, payload, true, func() (interface{}, error) {
		err := c.httpClient.Post(res, url, payload, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) PostWithFiles(res interface{}, url string, payload interface{}, files []File, opts ...RequestOption) error {
	return c.observeRequest("POST", url, opts, payload, true, func() (interface{}, error) {
		err := c.httpClient.PostWithFiles(res, url, payload, files, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) Request(res interface{}, requestMethod string, url string, opts ...RequestOption) error {
	return c.observeRequest(requestMethod, url, opts, nil, true, func() (interface{}, error) {
		err := c.httpClient.Request(res, requestMethod, url, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) RequestBinary(res BinaryResponse, requestMethod string, url string, opts ...RequestOption) error {
	return c.observeRequest(requestMethod, url, opts, nil, false, func() (interface{}, error) {
		err := c.httpClient.RequestBinary(res, requestMethod, url, opts...)
		return res, err
	})
}

func (c *tracingHttpClientDecorator) SetRequester(requester HttpClientRequester) {
	c.httpClient.SetRequester(requester)
}

func (c *tracingHttpClientDecorator) observeRequest(
	httpMethod string,
	url string,
	opts []RequestOption,
	payload interface{},
	isJsonResponse bool,
	requestMakingFunc func() (interface{}, error),
) error {
	reqOptions := RequestOptionsImpl{}
	for _, opt := range opts {
		opt(&reqOptions)
	}

	payloadBytes := reqOptions.GetPayload()
	if payload == nil && payloadBytes != nil {
		payload = payloadBytes.String()
	}

	span := c.tracer.StartSpanFromContext(c.ctx, c.operationName, c.serviceName, NormalizeURLForTracing(url))
	defer span.Finish()

	span.SetTag("HttpMethod", httpMethod)
	if payload != nil {
		span.SetTag("Payload", jsonPrettyString(payload))
	}
	span.SetTag("Endpoint", url)

	start := time.Now()
	res, err := requestMakingFunc()
	durationMs := time.Since(start).Milliseconds()
	span.SetTag("Request Duration", durationMs)
	if err != nil {
		span.SetTag("error", err)
	}
	if isJsonResponse {
		span.SetTag("Response", jsonPrettyString(res))
	} else {
		span.SetTag("Response", "(not-json-response)")
	}

	return err
}

func jsonPrettyString(i interface{}) string {
	if i == nil {
		return "(nil)"
	}
	jsonBytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return fmt.Sprintf("json-encoding-error: %v", err.Error())
	}
	return string(jsonBytes)
}

var numericPathSegmentRegex = regexp.MustCompile(`/\d+`)

// NormalizeURLForTracing extracts the path from rawURL and replaces numeric segments
// with {num} to aggregate similar requests in tracing systems.
func NormalizeURLForTracing(rawURL string) string {
	parsedURL, err := neturl.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	return numericPathSegmentRegex.ReplaceAllStringFunc(parsedURL.Path, func(match string) string {
		return `/{num}`
	})
}
