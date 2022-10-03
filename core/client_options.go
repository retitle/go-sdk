package core

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ClientOptions interface {
	GetProtocol() string
	GetHost() string
	GetUrl() string
	GetBasePath() string
	GetAudience() string
	SetProtocol(value string) ClientOptions
	SetHost(value string) ClientOptions
	SetUrl(value string) ClientOptions
	SetBasePath(value string) ClientOptions
	SetAudience(value string) ClientOptions
}

type ClientOptionsImpl struct {
	protocol string
	host     string
	url      string
	basePath string
	audience string
}

type ClientOption func(options ClientOptions)

func (co *ClientOptionsImpl) GetProtocol() string {
	return co.protocol
}

func (co *ClientOptionsImpl) GetHost() string {
	return co.host
}

func (co *ClientOptionsImpl) GetUrl() string {
	return co.url
}

func (co *ClientOptionsImpl) GetBasePath() string {
	return co.basePath
}

func (co *ClientOptionsImpl) GetAudience() string {
	return co.audience
}

func (co *ClientOptionsImpl) SetProtocol(value string) ClientOptions {
	co.protocol = value
	return co
}

func (co *ClientOptionsImpl) SetHost(value string) ClientOptions {
	co.host = value
	return co
}

func (co *ClientOptionsImpl) SetUrl(value string) ClientOptions {
	co.url = value
	return co
}

func (co *ClientOptionsImpl) SetBasePath(value string) ClientOptions {
	co.basePath = value
	return co
}

func (co *ClientOptionsImpl) SetAudience(value string) ClientOptions {
	co.audience = value
	return co
}

func NewClientOptions() ClientOptions {
	return &ClientOptionsImpl{}
}

func WithProtocol(protocol string) ClientOption {
	return func(o ClientOptions) {
		o.SetProtocol(strings.ToLower(protocol))
	}
}

func WithURL(URL string) ClientOption {
	return func(o ClientOptions) {
		o.SetUrl(URL)
	}
}

func WithHost(host string) ClientOption {
	return func(o ClientOptions) {
		o.SetHost(host)
	}
}

func WithBasePath(basePath string) ClientOption {
	return func(o ClientOptions) {
		o.SetBasePath(basePath)
	}
}

func WithAudience(audience string) ClientOption {
	return func(o ClientOptions) {
		o.SetAudience(audience)
	}
}

func WithQueryParam(qParam string, value string) RequestOption {
	return withQueryParam(qParam, value)
}

func WithQueryParamList(qParam string, values ...string) RequestOption {
	return WithReqOptQueryParamList(qParam, values)
}

func WithExpand(paths ...string) RequestOption {
	return WithReqOptQueryParamList("expand", paths)
}

func WithHeader(name string, value string) RequestOption {
	header := http.Header{}
	header.Add(name, value)
	return withHeaders(header)
}

func GetExpandFields(fieldIds ...string) string {
	if len(fieldIds) > 0 {
		return fmt.Sprintf("fields[%s]", strings.Join(fieldIds, ","))
	}
	return "fields"
}

func WithUpdatedAfter(ts int) RequestOption {
	return withQueryParam("updated_after", strconv.Itoa(ts))
}
