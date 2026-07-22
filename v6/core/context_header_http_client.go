package core

import "context"

// HeaderExtractor is a function that derives HTTP headers to inject into outbound
// Glide SDK requests from a context. Returning an empty or nil map is safe.
//
// Callers can provide multiple extractors; all results are merged before the
// request is made. If two extractors return the same header name the last one
// wins (consistent with the behaviour of WithHeader).
//
// Example — propagating a Compass trace ID:
//
//	func CompassHeaders(ctx context.Context) map[string]string {
//	    if tid := traceid.FromContext(ctx); tid != "" {
//	        return map[string]string{"X-Compass-Trace-Id": string(tid)}
//	    }
//	    return nil
//	}
type HeaderExtractor func(ctx context.Context) map[string]string

type contextHeaderHttpClient struct {
	ctx        context.Context
	httpClient HttpClient
	extractors []HeaderExtractor
}

var _ HttpClient = (*contextHeaderHttpClient)(nil)

// NewContextHeaderHttpClient wraps httpClient so that every outbound request
// automatically includes the headers returned by each extractor.
//
// ctx is captured at construction time and forwarded to each extractor on
// every request. Like TracingHttpClientDecorator, a fresh instance must be
// constructed for each request-context scope — do not share one instance
// across requests with different contexts.
//
// The injected headers are prepended to the opts slice, so explicit
// per-call WithHeader options and the Authorization header appended by
// ClientImpl.request always take precedence.
//
// To add new propagated headers in the future, simply add another
// HeaderExtractor (or extend an existing one) — no other code needs to change.
func NewContextHeaderHttpClient(
	ctx context.Context,
	httpClient HttpClient,
	extractors ...HeaderExtractor,
) HttpClient {
	return &contextHeaderHttpClient{
		ctx:        ctx,
		httpClient: httpClient,
		extractors: extractors,
	}
}

// headerOpts builds the RequestOption slice for the headers extracted from ctx.
func (c *contextHeaderHttpClient) headerOpts() []RequestOption {
	var opts []RequestOption
	for _, extract := range c.extractors {
		for name, value := range extract(c.ctx) {
			if name != "" && value != "" {
				opts = append(opts, WithHeader(name, value))
			}
		}
	}
	return opts
}

// prepend returns the extracted header options followed by the caller's opts,
// so that any explicit caller-supplied header for the same name wins.
func (c *contextHeaderHttpClient) prepend(opts []RequestOption) []RequestOption {
	headerOpts := c.headerOpts()
	if len(headerOpts) == 0 {
		return opts
	}
	return append(headerOpts, opts...)
}

func (c *contextHeaderHttpClient) Get(res interface{}, url string, opts ...RequestOption) error {
	return c.httpClient.Get(res, url, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) GetStream(res BinaryResponse, url string, opts ...RequestOption) error {
	return c.httpClient.GetStream(res, url, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) Post(res interface{}, url string, payload interface{}, opts ...RequestOption) error {
	return c.httpClient.Post(res, url, payload, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) PostWithFiles(res interface{}, url string, payload interface{}, files []File, opts ...RequestOption) error {
	return c.httpClient.PostWithFiles(res, url, payload, files, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) Request(res interface{}, requestMethod string, url string, opts ...RequestOption) error {
	return c.httpClient.Request(res, requestMethod, url, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) RequestBinary(res BinaryResponse, requestMethod string, url string, opts ...RequestOption) error {
	return c.httpClient.RequestBinary(res, requestMethod, url, c.prepend(opts)...)
}

func (c *contextHeaderHttpClient) SetRequester(requester HttpClientRequester) {
	c.httpClient.SetRequester(requester)
}
