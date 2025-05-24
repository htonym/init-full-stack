package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thofftech/init-full-stack/internal/models"
)

func ExtractUser(rawIDToken, rawAccessToken string, jwksCache *JWKSCache) (*models.User, error) {
	idToken, err := jwksCache.ParseToken(rawIDToken)
	if err != nil {
		return nil, err
	}

	_, err = jwksCache.ParseToken(rawAccessToken)
	if err != nil {
		return nil, err
	}

	user, err := userFromIDToken(idToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func userFromIDToken(token *jwt.Token) (*models.User, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	user := &models.User{
		FirstName: getStringClaim(claims, "given_name"),
		LastName:  getStringClaim(claims, "family_name"),
		Nickname:  getStringClaim(claims, "nickname"),
		ID:        getStringClaim(claims, "sub"),
		Email:     getStringClaim(claims, "email"),
		Username:  getStringClaim(claims, "cognito:username"),
	}

	return user, nil
}

func getStringClaim(claims jwt.MapClaims, key string) string {
	if value, ok := claims[key]; ok {
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}
	return ""
}
