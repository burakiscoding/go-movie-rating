package middleware

import (
	"log"
	"net/http"
	"time"
)

const AuthUserID = "middleware.auth.userID"

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
		// ctx := context.WithValue(r.Context(), AuthUserID, "user-id")
		// req := r.WithContext(ctx)
		// userID, ok := r.Context().Value(AuthUserID).(string)
		// if !ok {

		// }
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
