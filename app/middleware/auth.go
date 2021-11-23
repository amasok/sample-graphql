package middleware

import (
	"net/http"

	"github.com/amasok/sample-graphql/app/domain/authToken"
)

// TODO: configでoriginを切り分けられるようにする
func AuthForGraphql(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}
		bearer := "Bearer "
		auth = auth[len(bearer):]

		token, err := authToken.GenerateAuthByToken(auth)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		ctx := authToken.SetToken(r.Context(), token)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
