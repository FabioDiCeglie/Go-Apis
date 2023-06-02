package auth

import (
	"context"
	"net/http"

	"github.com/fabio/graphql/pkg/jwt"
)

type contextKey struct {
	name string
}

const tokenCtxKey = "token"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		// Allow unauthenticated users
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Validate JWT token
		tokenStr := header
		_, err := jwt.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// Put token in the request context
		ctx := context.WithValue(r.Context(), tokenCtxKey, tokenStr)
		r = r.WithContext(ctx)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r)
	})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
// func ForContext(ctx context.Context) *model.User {
// 	raw, _ := ctx.Value(tokenCtxKey).(*model.User)
// 	return raw
// }
