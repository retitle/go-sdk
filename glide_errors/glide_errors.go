package glide_errors

import (
	"fmt"
	"net/http"
	"strings"
)

type ApiError struct {
	Description     string
	StatusCode      int
	ResponseHeaders http.Header
	Params          map[string]interface{}
}

func (e *ApiError) Error() string {
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

func (e *ApiError) IsMissingScopes() bool {
	_, found := e.Params["missing_scopes"]
	return found
}

func (e *ApiError) GetMissingScopes() []string {
	if e.IsMissingScopes() {
		res := []string{}
		for _, missingScope := range e.Params["missing_scopes"].([]interface{}) {
			res = append(res, missingScope.(string))
		}
		return res
	}
	return []string{}
}

func (e *ApiError) HasToRequestScopes() bool {
	_, found := e.Params["request_scopes_url"]
	return found
}

func (e *ApiError) RequestScopesUrl() string {
	if e.HasToRequestScopes() {
		return e.Params["request_scopes_url"].(string)
	}

	return ""
}

func GetApiError(e error) *ApiError {
	if apiError, ok := e.(*ApiError); ok {
		return apiError
	} else {
		return &ApiError{
			Description: "Unknown Error",
			Params: map[string]interface{}{
				"err": e,
			},
		}
	}
}
