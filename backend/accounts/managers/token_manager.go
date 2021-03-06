package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/core"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	expirationTime    = time.Minute * 90
	tokenAccessSecret = "TOKEN_ACCESS_SECRET"
)

func init() {
	err := core.Injector.Provide(newTokenManager)
	core.CheckInjection(err, "TokenManager")
}

type tokenManager struct {
}

func newTokenManager() ports.TokenManager {
	return &tokenManager{}
}

func (m *tokenManager) Generate(acc *accounts.Account) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true

	claims["account_id"] = acc.ID
	claims["name"] = acc.Name
	claims["email"] = acc.Email
	claims["exp"] = time.Now().Add(expirationTime).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenAccessSecret := os.Getenv(tokenAccessSecret)
	sign := []byte(tokenAccessSecret)

	token, err := at.SignedString(sign)
	if err != nil {
		return "", errors.New(TokenError)
	}
	return token, nil
}
