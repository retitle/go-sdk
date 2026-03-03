package core_test

import (
	"testing"
	"time"

	"github.com/retitle/go-sdk/v5/core"
	"github.com/stretchr/testify/assert"
)

const One = 1
const Two = 2

func TestImpersonation(t *testing.T) {
	var (
		impersonation              core.Impersonation
		sub                        string
		scopes                     []string
		accessToken                string
		accessTokenExpires         time.Time
		expectedSub                string
		expectedScopes             []string
		expectedAccessToken        string
		expectedAccessTokenExpires time.Time
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should Get Sub",
			arrange: func() {
				expectedSub = "some sub"
				impersonation = core.NewImpersonationWithParams(expectedSub, []string{}, "", time.Now())
			},
			act: func() {
				sub = impersonation.GetSub()
			},
			assert: func() {
				assert.Equal(t, expectedSub, sub)
			},
		},
		{
			name: "Should Get scopes",
			arrange: func() {
				expectedScopes = []string{"Scope1", "Scope2"}
				impersonation = core.NewImpersonationWithParams("", expectedScopes, "", time.Now())
			},
			act: func() {
				scopes = impersonation.GetScopes()
			},
			assert: func() {
				assert.Len(t, scopes, Two)
				assert.Equal(t, expectedScopes, scopes)
			},
		},
		{
			name: "Should Get AccessToken",
			arrange: func() {
				expectedAccessToken = "Some valid token"
				impersonation = core.NewImpersonationWithParams("", []string{}, expectedAccessToken, time.Now())
			},
			act: func() {
				accessToken = impersonation.GetAccessToken()
			},
			assert: func() {
				assert.Equal(t, expectedAccessToken, accessToken)
			},
		},
		{
			name: "Should Get AccessTokenExpires",
			arrange: func() {
				expectedAccessTokenExpires = time.Now()
				impersonation = core.NewImpersonationWithParams("", []string{}, "", expectedAccessTokenExpires)
			},
			act: func() {
				accessTokenExpires = impersonation.GetAccessTokenExpires()
			},
			assert: func() {
				assert.Equal(t, expectedAccessTokenExpires, accessTokenExpires)
			},
		},
		{
			name: "Should Set Sub",
			arrange: func() {
				expectedSub = "some sub2"
				impersonation = core.NewImpersonation()
			},
			act: func() {
				impersonation.SetSub(expectedSub)
			},
			assert: func() {
				sub = impersonation.GetSub()
				assert.Equal(t, expectedSub, sub)
			},
		},
		{
			name: "Should Set scopes",
			arrange: func() {
				expectedScopes = []string{"Scope1"}
				impersonation = core.NewImpersonation()
			},
			act: func() {
				impersonation.SetScopes(expectedScopes)
			},
			assert: func() {
				scopes = impersonation.GetScopes()
				assert.Len(t, scopes, One)
				assert.Equal(t, expectedScopes, scopes)
			},
		},
		{
			name: "Should Set AccessToken",
			arrange: func() {
				expectedAccessToken = "Some valid token2"
				impersonation = core.NewImpersonation()
			},
			act: func() {
				impersonation.SetAccessToken(expectedAccessToken)
			},
			assert: func() {
				accessToken = impersonation.GetAccessToken()
				assert.Equal(t, expectedAccessToken, accessToken)
			},
		},
		{
			name: "Should Get AccessTokenExpires",
			arrange: func() {
				expectedAccessTokenExpires = time.Now()
				impersonation = core.NewImpersonation()
			},
			act: func() {
				impersonation.SetAccessTokenExpires(expectedAccessTokenExpires)
			},
			assert: func() {
				accessTokenExpires = impersonation.GetAccessTokenExpires()
				assert.Equal(t, expectedAccessTokenExpires, accessTokenExpires)
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
