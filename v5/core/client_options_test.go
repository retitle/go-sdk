package core_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/retitle/go-sdk/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestWithFunctions(t *testing.T) {
	var (
		clientOptionFunc core.ClientOption
		reqOptionFunc    core.RequestOption
		basePath, audience, host, protocol,
		url, keyId, michaelJordanId, babeRuthId,
		xSomeHeader, xSomeHeaderValue, messiId string
		updated_after        int
		clientOptions        core.ClientOptions
		requestClientOptions *core.RequestOptionsImpl
	)
	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should set up new protocol",
			arrange: func() {
				protocol = "https"
				clientOptions = core.NewClientOptions()
				clientOptionFunc = core.WithProtocol(protocol)
			},
			act: func() {
				clientOptionFunc(clientOptions)
			},
			assert: func() {
				clProtocol := clientOptions.GetProtocol()
				assert.Equal(t, protocol, clProtocol)
			},
		},
		{
			name: "Should set up new url",
			arrange: func() {
				clientOptions = core.NewClientOptions()
				requestClientOptions = &core.RequestOptionsImpl{}
				url = "some.random.url/path"
				clientOptionFunc = core.WithURL(url)
			},
			act: func() {
				clientOptionFunc(clientOptions)
			},
			assert: func() {
				clUrl := clientOptions.GetUrl()

				assert.Equal(t, url, clUrl)
			},
		},
		{
			name: "Should set up new host",
			arrange: func() {
				host = "some.other.url"
				clientOptions = core.NewClientOptions()
				clientOptionFunc = core.WithHost(host)
			},
			act: func() {
				clientOptionFunc(clientOptions)
			},
			assert: func() {
				clHost := clientOptions.GetHost()

				assert.Equal(t, host, clHost)
			},
		},
		{
			name: "Should set up new base path",
			arrange: func() {
				basePath = "/some/path"
				clientOptions = core.NewClientOptions()
				clientOptionFunc = core.WithBasePath(basePath)
			},
			act: func() {
				clientOptionFunc(clientOptions)
			},
			assert: func() {
				clBasePath := clientOptions.GetBasePath()

				assert.Equal(t, basePath, clBasePath)
			},
		},
		{
			name: "Should set up new audience",
			arrange: func() {
				audience = "ALL"
				clientOptions = core.NewClientOptions()
				clientOptionFunc = core.WithAudience(audience)
			},
			act: func() {
				clientOptionFunc(clientOptions)
			},
			assert: func() {
				clAudience := clientOptions.GetAudience()

				assert.Equal(t, audience, clAudience)
			},
		},
		{
			name: "Should set up query params",
			arrange: func() {
				keyId = "theBests"
				michaelJordanId = "23"
				babeRuthId = "3"
				requestClientOptions = &core.RequestOptionsImpl{}
				reqOptionFunc = core.WithQueryParamList(keyId, michaelJordanId, babeRuthId)
			},
			act: func() {
				reqOptionFunc(requestClientOptions)
			},
			assert: func() {
				reqClParams := requestClientOptions.GetQParams()
				val, ok := reqClParams[keyId]
				assert.True(t, ok)
				assert.Equal(t, val, strings.Join([]string{michaelJordanId, babeRuthId}, ","))
			},
		},
		{
			name: "Should set up expand params",
			arrange: func() {
				messiId = "10"
				reqOptionFunc = core.WithExpand(messiId)
			},
			act: func() {
				reqOptionFunc(requestClientOptions)
			},
			assert: func() {
				reqClParams := requestClientOptions.GetQParams()
				val, ok := reqClParams["expand"]
				assert.True(t, ok)
				assert.Equal(t, val, strings.Join([]string{messiId}, ","))
			},
		},
		{
			name: "Should set up expand params",
			arrange: func() {
				messiId = "10"
				reqOptionFunc = core.WithExpand(messiId)
			},
			act: func() {
				reqOptionFunc(requestClientOptions)
			},
			assert: func() {
				reqClParams := requestClientOptions.GetQParams()
				val, ok := reqClParams["expand"]
				assert.True(t, ok)
				assert.Equal(t, val, strings.Join([]string{messiId}, ","))
			},
		},
		{
			name: "Should set up updated_after",
			arrange: func() {
				updated_after = 1000000000
				reqOptionFunc = core.WithUpdatedAfter(updated_after)
			},
			act: func() {
				reqOptionFunc(requestClientOptions)
			},
			assert: func() {
				reqClParams := requestClientOptions.GetQParams()
				val, ok := reqClParams["updated_after"]
				assert.True(t, ok)
				assert.Equal(t, val, strconv.Itoa(updated_after))
			},
		},
		{
			name: "Should add headers",
			arrange: func() {
				xSomeHeader = "x-some-header"
				xSomeHeaderValue = "x-header"
				reqOptionFunc = core.WithHeader(xSomeHeader, xSomeHeaderValue)
			},
			act: func() {
				reqOptionFunc(requestClientOptions)
			},
			assert: func() {
				headers := requestClientOptions.GetHeaders()
				h := headers.Get(xSomeHeader)
				assert.Equal(t, h, xSomeHeaderValue)
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange()

			tt.act()

			tt.assert()
		})
	}
}
