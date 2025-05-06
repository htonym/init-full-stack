package auth

import (
	"fmt"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

type JWKSCache struct {
	mutex sync.RWMutex
	// value *JsonWebKeySet
	url string
}

func (jc *JWKSCache) ParseToken(tokenString string) (*jwt.Token, error) {

	return nil, nil
}

func NewJWKSCache(domain string) *JWKSCache {
	return &JWKSCache{
		url: fmt.Sprintf("http://%s/.well-known/jwks.json", domain),
	}
}
