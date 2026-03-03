package core

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Scopes []string `json:"scopes"`
	jwt.StandardClaims
}

func GetJwt(key Key, iss string, sub string, aud string, scopes []string, expiresIn int64) (string, error) {
	k, err := key.GetDecoded()
	if err != nil {
		return "", &ApiErrorImpl{
			Params: map[string]interface{}{
				"desc": "Error decoding key",
			},
			Err: err,
		}
	}

	issuedAt := time.Now().Unix()

	token := jwt.NewWithClaims(key.GetJwtSigningMethod(), claims{
		Scopes: scopes,
		StandardClaims: jwt.StandardClaims{
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			IssuedAt:  issuedAt,
			ExpiresAt: issuedAt + expiresIn,
		},
	})
	t, err := token.SignedString(k)
	if err != nil {
		return "", &ApiErrorImpl{
			Params: map[string]interface{}{
				"desc": "Error signing JWT",
			},
			Err: err,
		}
	}
	return t, nil
}
