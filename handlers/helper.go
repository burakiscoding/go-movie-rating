package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"iss": "movie-rating",
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWTSECRET")))
}

func VerifyToken(tokenString string) (types.TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(os.Getenv("JWTSECRET")), nil
	})

	if err != nil {
		return types.TokenPayload{}, fmt.Errorf("Invalid token")
	}

	if !token.Valid {
		return types.TokenPayload{}, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return types.TokenPayload{}, fmt.Errorf("Invalid token")
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return types.TokenPayload{}, fmt.Errorf("Invalid token")
	}

	return types.TokenPayload{Id: id}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
