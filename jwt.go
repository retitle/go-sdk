package glide

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/security"
)

type claims struct {
	Scopes []string `json:"scopes"`
	jwt.StandardClaims
}

func getJwt(key security.Key, iss string, sub string, aud string, scopes []string, expiresIn int64) (string, *glide_errors.ApiError) {
	k, err := key.GetDecoded()
	if err != nil {
		return "", &glide_errors.ApiError{
			Params: map[string]interface{}{
				"desc": "Error decoding key",
				"err":  err,
			},
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
		return "", &glide_errors.ApiError{
			Params: map[string]interface{}{
				"desc": "Error signing JWT",
				"err":  err,
			},
		}
	}
	return t, nil
}
