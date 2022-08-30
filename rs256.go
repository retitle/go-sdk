package glide

import (
	"crypto"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

type rsa256Key struct {
	privateKeyFilePath string
	privateKeyBytes    []byte
}

func GetRsa256Key(privateKeyFilePath string) rsa256Key {
	return rsa256Key{
		privateKeyFilePath: privateKeyFilePath,
	}
}

func GetRsa256KeyFromPEMBytes(privateKeyBytes []byte) rsa256Key {
	return rsa256Key{
		privateKeyBytes: privateKeyBytes,
	}
}

func GetRsa256KeyFromPEMString(privateKeyString string) rsa256Key {
	return GetRsa256KeyFromPEMBytes([]byte(privateKeyString))
}

func (k rsa256Key) GetJwtSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodRS256
}

func (k rsa256Key) GetDecoded() (crypto.PrivateKey, error) {
	var data []byte
	if k.privateKeyFilePath != "" {
		var err error
		data, err = ioutil.ReadFile(k.privateKeyFilePath)
		if err != nil {
			return nil, err
		}
	} else {
		data = k.privateKeyBytes
	}

	return jwt.ParseRSAPrivateKeyFromPEM(data)
}
