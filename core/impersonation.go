package core

import (
	"time"
)

//go:generate mockery --name=Impersonation --filename=impersonation.go --output=mocks
type Impersonation interface {
	GetSub() string
	GetScopes() []string
	GetAccessToken() string
	GetAccessTokenExpires() time.Time
	SetSub(string) Impersonation
	SetScopes([]string) Impersonation
	SetAccessToken(string) Impersonation
	SetAccessTokenExpires(time.Time) Impersonation
}

type ImpersonationImpl struct {
	sub                string
	scopes             []string
	accessToken        string
	accessTokenExpires time.Time
}

func (imp *ImpersonationImpl) GetSub() string {
	return imp.sub
}

func (imp *ImpersonationImpl) GetScopes() []string {
	return imp.scopes
}

func (imp *ImpersonationImpl) GetAccessToken() string {
	return imp.accessToken
}

func (imp *ImpersonationImpl) GetAccessTokenExpires() time.Time {
	return imp.accessTokenExpires
}

func (imp *ImpersonationImpl) SetSub(value string) Impersonation {
	imp.sub = value
	return imp
}

func (imp *ImpersonationImpl) SetScopes(value []string) Impersonation {
	imp.scopes = value
	return imp
}

func (imp *ImpersonationImpl) SetAccessToken(value string) Impersonation {
	imp.accessToken = value
	return imp
}

func (imp *ImpersonationImpl) SetAccessTokenExpires(value time.Time) Impersonation {
	imp.accessTokenExpires = value
	return imp
}

func NewImpersonation() Impersonation {
	return &ImpersonationImpl{}
}

func NewImpersonationWithParams(
	sub string,
	scopes []string,
	accessToken string,
	accessTokenExpires time.Time) Impersonation {
	return &ImpersonationImpl{
		sub:                sub,
		scopes:             scopes,
		accessTokenExpires: accessTokenExpires,
		accessToken:        accessToken,
	}
}
