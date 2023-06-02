package auth

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/fabio/graphql/database"
	"github.com/fabio/graphql/graph/model"
	"github.com/fabio/graphql/pkg/jwt"
	"github.com/jinzhu/gorm"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := model.User{Name: username}
			id, err := GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

// GetUserIdByUsername checks if a user exists in the database by the given username
func GetUserIdByUsername(username string) (int, error) {
	db := database.Db
	var user model.User
	if err := db.Where("name = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("No user found with username: ", username)
			return 0, nil
		}
		return 0, err
	}

	num, err := strconv.Atoi(user.ID)
	if err != nil {
		return 0, err
	}

	return num, nil
}
