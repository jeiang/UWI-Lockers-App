package handlers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	cache "github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

const (
	// used to fill time on the check
	sample_hash = "$2a$14$eNKVCPjXz48hg7E6qAbupO/Vv/0L5LVg0wE/.ycLKDrZ.4BOnPkyS"
)

type LoginHandler struct {
	DB         *pgxpool.Pool
	TokenCache *cache.Cache
}

type Auth struct {
	Token string `json:"token"`
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("received a request", "ipaddr", r.RemoteAddr, "form values", r.Form)
	// TODO: sanitize usernaem
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "missing arguments", http.StatusBadRequest)
		return
	}

	var passwordHash []byte
	hashFound := true
	err := lh.DB.QueryRow(
		context.Background(),
		"select passwordhash from users where name=$1",
		username,
	).Scan(&passwordHash)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			log.Warn("login attempt with unknown username", "username", username)
			passwordHash = []byte(sample_hash)
			hashFound = false
		} else {
			log.Error(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if !hashFound || err != nil {
		log.Warn("failed login attempt detected", "username", username)
		http.Error(w, "Access Denied", http.StatusForbidden)
		return
	}

	// Generate an auth token user CSPRNG
	bytes := make([]byte, 128)
	_, err = rand.Read(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := hex.EncodeToString(bytes)
	auth := Auth{Token: token}

	lh.TokenCache.Set(auth.Token, username, 2*time.Hour)

	js, err := json.Marshal(auth)
	if err != nil {
		log.Error("error occurred while trying to marshal the bearer token to json", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
