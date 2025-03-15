package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tanmaygupta069/order-service-go/config"
	"google.golang.org/grpc/metadata"
)

type AuthPackage interface {
	GetTokenFromMetadata(md metadata.MD) (string, error)
	ExtractUserIDFromToken(tokenString string) (string, error)
}

var cfg, _ = config.GetConfig()

type AuthPackageImp struct{
}


func NewAuthPackage()AuthPackage{
	return &AuthPackageImp{
	}
}

func (r *AuthPackageImp) GetTokenFromMetadata(md metadata.MD) (string, error) {
	token := md.Get("Authorization")
	if len(token) == 0 {
		return "", fmt.Errorf("no token found")
	}
	return token[0], nil
}

func (r *AuthPackageImp) ExtractUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return "", fmt.Errorf("email not found in token")
		}
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}