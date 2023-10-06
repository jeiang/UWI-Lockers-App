package middleware

import (
	"net/http"
	"strings"

	cache "github.com/patrickmn/go-cache"
)

func AuthMW(tokenCache *cache.Cache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		token := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))
		if _, ok := tokenCache.Get(token); ok {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	})
}
