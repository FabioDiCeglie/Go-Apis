package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fabio/graphql/graph/model"
	"github.com/fabio/graphql/pkg/jwt"
	"github.com/fabio/graphql/pkg/utils"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

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
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// create user and check if user exists in db
		user := model.User{Name: username}
		id, err := utils.GetUserIdByUsername(username)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user.ID = strconv.Itoa(id)
		// put it in context
		ctx := context.WithValue(r.Context(), userCtxKey, &user)
		r = r.WithContext(ctx)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r)
	})
}

// Finds the user from the context. REQUIRES Middleware to have run.
func FindUser(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
