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
	c := GetClient(TEST_CLIENT_KEY, stubDummyKey).(*client)

	wantAud := "api.glide.com"
	if gotAud := c.options.audience; gotAud != wantAud {
		t.Errorf("GetClient().options.audience = %q, want %q", gotAud, wantAud)
	}
	wantBasePath := ""
	if gotBasePath := c.options.basePath; gotBasePath != wantBasePath {
		t.Errorf("GetClient().options.basePath = %q, want %q", gotBasePath, wantBasePath)
	}
	wantProtocol := "https"
	if gotProtocol := c.options.protocol; gotProtocol != wantProtocol {
		t.Errorf("GetClient().options.protocol = %q, want %q", gotProtocol, wantProtocol)
	}
	wantHost := "api.glide.com"
	if gotHost := c.options.host; gotHost != wantHost {
		t.Errorf("GetClient().options.host = %q, want %q", gotHost, wantHost)
	}
	wantURL := "api.glide.com"
	if gotURL := c.options.host; gotURL != wantURL {
		t.Errorf("GetClient().options.url = %q, want %q", gotURL, wantURL)
	}
}
