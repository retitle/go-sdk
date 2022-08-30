package core

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	ApiErrorHttpRequestMethod = "Error while initializing http request"
	ApiErrorHttpBodyResponse  = "Error reading body response"
	ApiErrorHttpRequest       = "Error when sending http request to Glide API"
	ApiErrorHttpReadResponse  = "Error while reading response body"
)

type ApiError interface {
	error
	Unwrap() error
	IsMissingScopes() bool
	GetMissingScopes() []string
	HasToRequestScopes() bool
	RequestScopesUrl() string
	GetDescription() string
	GetStatusCode() int
}

type ApiErrorImpl struct {
	Description     string
	StatusCode      int
	ResponseHeaders http.Header
	Params          map[string]interface{}
	Err             error
}

func (e *ApiErrorImpl) GetStatusCode() int {
	return e.StatusCode
}

func (e *ApiErrorImpl) Unwrap() error {
	return e.Err
}

func (e *ApiErrorImpl) Error() string {
	params := ""
	for k, v := range e.Params {
		params += fmt.Sprintf("\n\t\t%s: %+v", k, v)
	}
	return strings.Join([]string{
		"ApiError",
		fmt.Sprintf("Description: %s", e.Description),
		fmt.Sprintf("Status Code: %d", e.StatusCode),
		fmt.Sprintf("Params: %s", params),
	}, "\n\t")
}

func (e *ApiErrorImpl) IsMissingScopes() bool {
	_, found := e.Params["missing_scopes"]
	return found
}

func (e *ApiErrorImpl) GetDescription() string {
	return e.Description
}

func (e *ApiErrorImpl) GetMissingScopes() []string {
	if e.IsMissingScopes() {
		res := []string{}
		for _, missingScope := range e.Params["missing_scopes"].([]interface{}) {
			res = append(res, missingScope.(string))
		}
		return res
	}
	return []string{}
}

func (e *ApiErrorImpl) HasToRequestScopes() bool {
	_, found := e.Params["request_scopes_url"]
	return found
}

func (e *ApiErrorImpl) RequestScopesUrl() string {
	if e.HasToRequestScopes() {
		return e.Params["request_scopes_url"].(string)
	}

	return ""
}

func GetApiError(e error) *ApiErrorImpl {
	if apiError, ok := e.(*ApiErrorImpl); ok {
		return apiError
	} else {
		return &ApiErrorImpl{
			Description: "Unknown Error",
			StatusCode:  http.StatusBadRequest,
			Err:         e,
		}
	}
}

func NewApiErrorWithArgs(
	description string,
	statusCode int,
	responseHeaders http.Header,
	params map[string]interface{},
	baseError error,
) *ApiErrorImpl {
	return &ApiErrorImpl{
		Description:     description,
		StatusCode:      statusCode,
		ResponseHeaders: responseHeaders,
		Params:          params,
		Err:             baseError,
	}
}

func NewHttpMethodApiError(e error) *ApiErrorImpl {
	return &ApiErrorImpl{
		Description: ApiErrorHttpRequestMethod,
		StatusCode:  http.StatusInternalServerError,
		Err:         e,
	}
}

func NewHttpRequestApiError(e error) *ApiErrorImpl {
	return &ApiErrorImpl{
		Description: ApiErrorHttpRequest,
		Err:         e,
	}
}
