package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/golang-jwt/jwt/v5"
)

const AuthUserId = "middleware.auth.userId"

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		substrings := strings.Split(authHeader, " ")

		if len(substrings) != 2 {
			WriteUnauthorized(w)
			return
		}
		if substrings[0] != "Bearer" {
			WriteUnauthorized(w)
			return
		}

		token := substrings[1]

		payload, err := VerifyToken(token)
		if err != nil {
			WriteUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUserId, payload.Id)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

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
