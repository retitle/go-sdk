package core_test

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/core/mocks"
	"github.com/stretchr/testify/assert"
)

type FakePublicKey struct{}

func (fpk *FakePublicKey) Equal(x crypto.PublicKey) bool {
	return true
}

type FakePrivateKey struct{}

func (fpk *FakePrivateKey) Public() crypto.PublicKey {
	return FakePublicKey{}
}
func (fpk *FakePrivateKey) Equal(x crypto.PrivateKey) bool {
	return true
}

func TestGetJwt(t *testing.T) {
	var (
		key       mocks.Key
		jwtToken  string
		err       error
		iss       string
		sub       string
		aud       string
		scopes    []string
		expiresIn int64
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should get a valid jwt",
			arrange: func() {
				key = mocks.Key{}
				cryptoPrivateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
				key.On("GetDecoded").Return(cryptoPrivateKey, nil)
				key.On("GetJwtSigningMethod").Return(jwt.SigningMethodRS256, nil)
				iss = "fakeiss"
				sub = "fakesub"
				aud = "fakeaud"
				scopes = []string{"FAKE_SCOPE"}
				expiresIn = int64((time.Minute * time.Duration(60)).Seconds())
			},
			act: func() {
				jwtToken, err = core.GetJwt(&key,
					iss,
					sub,
					aud,
					scopes,
					expiresIn)
			},
			assert: func() {
				assert.Nil(t, err)
				assert.NotEmpty(t, jwtToken)
			},
		},
		{
			name: "Should get err generating a jwt, getDecoded returning error",
			arrange: func() {
				key = mocks.Key{}
				key.On("GetDecoded").Return(nil, fmt.Errorf("Error"))
				key.On("GetJwtSigningMethod").Return(jwt.SigningMethodRS256, nil)
				iss = "fakeiss"
				sub = "fakesub"
				aud = "fakeaud"
				scopes = []string{"FAKE_SCOPE"}
				expiresIn = int64((time.Minute * time.Duration(60)).Seconds())
			},
			act: func() {
				jwtToken, err = core.GetJwt(&key,
					iss,
					sub,
					aud,
					scopes,
					expiresIn)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				assert.Empty(t, jwtToken)
			},
		},
		{
			name: "Should get err generating a jwt, empty private key",
			arrange: func() {
				key = mocks.Key{}
				emptyCryptoPrivateKey := FakePrivateKey{}
				key.On("GetDecoded").Return(emptyCryptoPrivateKey, nil)
				key.On("GetJwtSigningMethod").Return(jwt.SigningMethodRS256, nil)
				iss = "fakeiss"
				sub = "fakesub"
				aud = "fakeaud"
				scopes = []string{"FAKE_SCOPE"}
				expiresIn = int64((time.Minute * time.Duration(60)).Seconds())
			},
			act: func() {
				jwtToken, err = core.GetJwt(&key,
					iss,
					sub,
					aud,
					scopes,
					expiresIn)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				assert.Empty(t, jwtToken)
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
