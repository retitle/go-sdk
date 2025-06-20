package core_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/retitle/go-sdk/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	var (
		apiError          *core.ApiErrorImpl
		err               error
		isMissingScope    bool
		hasToRequestScope bool
		requestScope      string
		expectedErr       string
		missingScopes     []string
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Error, Should the correct Error",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Err = fmt.Errorf("Error")
			},
			act: func() {
				err = apiError.Unwrap()
			},
			assert: func() {
				assert.Equal(t, err, apiError.Err)
			},
		},
		{
			name: "IsMissingScope, Should return False, scope is missing",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{}

			},
			act: func() {
				isMissingScope = apiError.IsMissingScopes()
			},
			assert: func() {
				assert.False(t, isMissingScope)
			},
		},
		{
			name: "IsMissingScope, Should return True, scope is not missing",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{"missing_scopes": ""}
			},
			act: func() {
				isMissingScope = apiError.IsMissingScopes()
			},
			assert: func() {
				assert.True(t, isMissingScope)
			},
		},
		{
			name: "GetMissingScopes, Should return the scopes that are missing",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{"missing_scopes": []interface{}{"TRX", "LOGIN"}}
			},
			act: func() {
				missingScopes = apiError.GetMissingScopes()
			},
			assert: func() {
				assert.Len(t, missingScopes, 2)
				assert.Equal(t, "TRX", missingScopes[0])
				assert.Equal(t, "LOGIN", missingScopes[1])
			},
		},
		{
			name: "GetMissingScopes, Should return empty missing scopes",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{}
			},
			act: func() {
				missingScopes = apiError.GetMissingScopes()
			},
			assert: func() {
				assert.Len(t, missingScopes, 0)
			},
		},
		{
			name: "HasToRequestScopes, Should return False, request_scopes_url is missing",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{}

			},
			act: func() {
				hasToRequestScope = apiError.HasToRequestScopes()
			},
			assert: func() {
				assert.False(t, hasToRequestScope)
			},
		},

		{
			name: "HasToRequestScopes, Should return True, request_scopes_url is present",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{"request_scopes_url": ""}

			},
			act: func() {
				hasToRequestScope = apiError.HasToRequestScopes()
			},
			assert: func() {
				assert.True(t, hasToRequestScope)
			},
		},
		{
			name: "RequestScopesUrl, Should return not empty request_scopes_url",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{"request_scopes_url": "SOME_URL"}

			},
			act: func() {
				requestScope = apiError.RequestScopesUrl()
			},
			assert: func() {
				assert.Equal(t, "SOME_URL", requestScope)
			},
		},
		{
			name: "RequestScopesUrl, Should return empty request_scopes_url",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Params = map[string]interface{}{}

			},
			act: func() {
				requestScope = apiError.RequestScopesUrl()
			},
			assert: func() {
				assert.Empty(t, requestScope)
			},
		},
		{
			name: "RequestScopesUrl, Should return custom Error string",
			arrange: func() {
				apiError = &core.ApiErrorImpl{}
				apiError.Description = "Some description"
				apiError.StatusCode = http.StatusOK
				apiError.Params = map[string]interface{}{"param": "1"}

			},
			act: func() {
				expectedErr = apiError.Error()
			},
			assert: func() {
				p := fmt.Sprintf("\n\t\t%s: %+v", "param", "1")
				expectedStrs := strings.Join([]string{
					"ApiError",
					fmt.Sprintf("Description: %s", apiError.Description),
					fmt.Sprintf("Status Code: %d", apiError.StatusCode),
					fmt.Sprintf("Params: %s", p),
				}, "\n\t")
				assert.Equal(t, expectedStrs, expectedErr)
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
