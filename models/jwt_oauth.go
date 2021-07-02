package models

type JwtOauth struct {
	GrantType string `json:"grant_type"`
	Assertion string `json:"assertion"`
}

type JwtOauthAccessToken struct {
	AccessToken      string   `json:"access_token"`
	ExpiresIn        int      `json:"expires_in"`
	MissingScopes    []string `json:"missing_scopes"`
	RequestScopesUrl string   `json:"request_scopes_url"`
}
