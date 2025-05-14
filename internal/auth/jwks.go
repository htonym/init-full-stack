package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

type JWKSCache struct {
	mutex sync.RWMutex
	value *JsonWebKeySet
	url   string
}

type JsonWebKeySet struct {
	Keys []JsonWebKey `json:"keys"`
}

type JsonWebKey struct {
	Kid string `json:"kid"`
	E   string `json:"e"`
	N   string `json:"n"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
}

func NewJWKSCache(jwksUrl string) *JWKSCache {
	return &JWKSCache{
		url: jwksUrl,
	}
}

func (jc *JWKSCache) RefreshJWKS() error {
	jc.mutex.Lock()
	defer jc.mutex.Unlock()

	resp, err := http.Get(jc.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&jc.value); err != nil {
		return err
	}

	log.Printf("Loaded %d new JSON web token public keys", len(jc.value.Keys))

	return nil
}

func (jc *JWKSCache) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		publicKey, err := jc.buildRSAPublicKey(token.Header["kid"].(string))
		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Build an rsa.PublicKey from JsonWebKey
func (jc *JWKSCache) buildRSAPublicKey(kid string) (*rsa.PublicKey, error) {

	var jwk *JsonWebKey
	for _, key := range jc.value.Keys {
		if key.Kid == kid {
			jwk = &key
		}
	}

	nBytes, err := decodeBase64URL(jwk.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := decodeBase64URL(jwk.E)
	if err != nil {
		return nil, err
	}
	n := new(big.Int).SetBytes(nBytes)
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}
	return &rsa.PublicKey{N: n, E: e}, nil
}

func decodeBase64URL(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
