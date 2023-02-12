package middleware

import (
	"chaos/backend/database/model"
	"chaos/backend/service"
	"chaos/backend/utility"
	"context"
	"net/http"
)

var userCtxKey = &userContext{"user"}

type userContext struct {
	name string
}

// using jwt token for authentication
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			// retrieve token from bearer token
			token := utility.GetTokenStringFromBearerToken(header)

			if token == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			// validate jwt token
			tokenStr := token
			user := service.ValidateToken(tokenStr)
			if user == nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// if user is not activated, then it will be unauthorized
			if !user.IsActivated {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
