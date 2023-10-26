package middleware

import (
	"net/http"
	"strings"

	"aidanpinard.co/lockers-app/handlers"
	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt/v5"
	cache "github.com/patrickmn/go-cache"
)

func AuthMW(tokenCache *cache.Cache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		tokenString := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return handlers.PrivateKey.Public(), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodEdDSA.Alg()}))

		if err != nil || !token.Valid {
			log.Error("attempted login with invalid jwt", "err", err, "valid", token.Valid)
			http.Error(w, "Access Denied", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
