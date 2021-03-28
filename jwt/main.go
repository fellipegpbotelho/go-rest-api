package jwt

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type JWTManager struct {
	secretKey string
	issuer    string
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "4150953A511F9F6D74E17D5AC22E8BE54E9FAB56"
	}
	return secret
}

func BuildJWTManager() *JWTManager {
	return &JWTManager{
		secretKey: getSecretKey(),
		issuer:    "John Doe",
	}
}

func (jwtManager *JWTManager) GenerateToken(email string) string {
	isUser := true
	claims := &AuthCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    jwtManager.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generatedToken, err := token.SignedString([]byte(jwtManager.secretKey))
	if err != nil {
		panic(err)
	}
	return generatedToken
}

func (jwtManager *JWTManager) ValidateToken(tokenToValidate string) (*jwt.Token, error) {
	return jwt.Parse(tokenToValidate, func(token *jwt.Token) (interface{}, error) {
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(jwtManager.secretKey), nil
	})
}
