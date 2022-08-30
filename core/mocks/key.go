// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	crypto "crypto"

	jwt "github.com/dgrijalva/jwt-go"
	mock "github.com/stretchr/testify/mock"
)

// Key is an autogenerated mock type for the Key type
type Key struct {
	mock.Mock
}

// GetDecoded provides a mock function with given fields:
func (_m *Key) GetDecoded() (crypto.PrivateKey, error) {
	ret := _m.Called()

	var r0 crypto.PrivateKey
	if rf, ok := ret.Get(0).(func() crypto.PrivateKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJwtSigningMethod provides a mock function with given fields:
func (_m *Key) GetJwtSigningMethod() jwt.SigningMethod {
	ret := _m.Called()

	var r0 jwt.SigningMethod
	if rf, ok := ret.Get(0).(func() jwt.SigningMethod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.SigningMethod)
		}
	}

	return r0
}