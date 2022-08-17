package core

import (
	"crypto"

	"github.com/dgrijalva/jwt-go"
)

type Key interface {
	GetJwtSigningMethod() jwt.SigningMethod
	GetDecoded() (crypto.PrivateKey, error)
}
