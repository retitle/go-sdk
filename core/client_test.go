package core_test

import (
	"crypto"
	"errors"
	"testing"

	glide "github.com/retitle/go-sdk/v3"

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
	c := glide.GetClient(TEST_CLIENT_KEY, stubDummyKey).(*glide.ClientImpl)

	wantAud := "api.glide.com"
	if gotAud := c.GetOptions().GetAudience(); gotAud != wantAud {
		t.Errorf("GetClient().GetOptions().audience = %q, want %q", gotAud, wantAud)
	}
	wantBasePath := ""
	if gotBasePath := c.GetOptions().GetBasePath(); gotBasePath != wantBasePath {
		t.Errorf("GetClient().GetOptions().basePath = %q, want %q", gotBasePath, wantBasePath)
	}
	wantProtocol := "https"
	if gotProtocol := c.GetOptions().GetProtocol(); gotProtocol != wantProtocol {
		t.Errorf("GetClient().GetOptions().protocol = %q, want %q", gotProtocol, wantProtocol)
	}
	wantHost := "api.glide.com"
	if gotHost := c.GetOptions().GetHost(); gotHost != wantHost {
		t.Errorf("GetClient().GetOptions().host = %q, want %q", gotHost, wantHost)
	}
	wantURL := "api.glide.com"
	if gotURL := c.GetOptions().GetHost(); gotURL != wantURL {
		t.Errorf("GetClient().GetOptions().url = %q, want %q", gotURL, wantURL)
	}
}
