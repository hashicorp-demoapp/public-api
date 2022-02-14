package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/keratin/authn-go/authn"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the session token and packs the session into the context.
func Middleware(auth *authn.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
			w.Header().Set("Content-Type", "application/json")

			ctx := r.Context()
			authorization := r.Header.Get("Authorization")
			token := strings.TrimPrefix(authorization, "Bearer ")
			// decodedToken, err := auth.SubjectFrom(token)
			// if err != nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			ctx = context.WithValue(ctx, userCtxKey, token)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// GetAuthHeader returns the auth header.
// REQUIRES Middleware to have run.
func GetAuthHeader(ctx context.Context) string {
	return ctx.Value(userCtxKey).(string)
}
