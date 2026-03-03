package core

import (
	"crypto"

	"github.com/dgrijalva/jwt-go"
)

//go:generate mockery --name=Key --filename=key.go --output=mocks
type Key interface {
	GetJwtSigningMethod() jwt.SigningMethod
	GetDecoded() (crypto.PrivateKey, error)
}
