package middleware

import "net/http"

// TODO: configでoriginを切り分けられるようにする
func CORSForGraphql(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := "localhost:8080"
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST")
		}
		next.ServeHTTP(w, r)
	})
}
