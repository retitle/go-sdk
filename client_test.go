package glide

import (
	"crypto"
	"errors"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

const TEST_CLIENT_KEY = "test-client-key"

type dummyKey struct {
}

func (k dummyKey) GetJwtSigningMethod() jwt.SigningMethod {
	return nil
}

func (k dummyKey) GetDecoded() (crypto.PrivateKey, error) {
	return nil, errors.New("Dummy")
}

func TestGetClientOptions(t *testing.T) {
	stubDummyKey := dummyKey{}
	c := GetClient(TEST_CLIENT_KEY, stubDummyKey, nil)

	wantAud := "api.glide.com"
	if gotAud := c.options.Audience; gotAud != wantAud {
		t.Errorf("GetClient().options.Audience = %q, want %q", gotAud, wantAud)
	}
	wantBasePath := "/"
	if gotBasePath := c.options.BasePath; gotBasePath != wantBasePath {
		t.Errorf("GetClient().options.BasePath = %q, want %q", gotBasePath, wantBasePath)
	}
	wantProtocol := "https"
	if gotProtocol := c.options.Protocol; gotProtocol != wantProtocol {
		t.Errorf("GetClient().options.Protocol = %q, want %q", gotProtocol, wantProtocol)
	}
	wantServer := "api.glide.com"
	if gotServer := c.options.Server; gotServer != wantServer {
		t.Errorf("GetClient().options.Server = %q, want %q", gotServer, wantServer)
	}
}
