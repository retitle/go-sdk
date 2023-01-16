package core

import (
	"crypto"

	"github.com/golang-jwt/jwt/v4"
)

//go:generate mockery --name=Key --filename=key.go --output=mocks
type Key interface {
	GetJwtSigningMethod() jwt.SigningMethod
	GetDecoded() (crypto.PrivateKey, error)
}
