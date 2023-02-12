package middleware

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/utility"
	"context"
	"net/http"
)

var keyCtx = &apikeyContext{"key"}

type apikeyContext struct {
	apikey string
}

// using apikey as an authentication method
func APIMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			var key model.APIKey
			key.Token = header
			if err := key.GetKeyByToken(database.Connection); err != nil {
				if utility.RecordNotFound(err) {
					// if key doesn't exist, it is unauthorized
					http.Error(w, "Invalid key", http.StatusUnauthorized)
					return
				} else {
					// otherwise, something went wrong for connecting to the database
					http.Error(w, "Unexpected error", http.StatusInternalServerError)
					return
				}
			}

			// check if key is enabled
			if !key.IsEnabled {
				http.Error(w, "APIkey disabled", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), keyCtx, &key)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForApiContext(ctx context.Context) *model.APIKey {
	raw, _ := ctx.Value(keyCtx).(*model.APIKey)
	return raw
}
