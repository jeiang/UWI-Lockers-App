package handlers

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

const (
	// used to fill time on the check
	sample_hash = "$2a$14$eNKVCPjXz48hg7E6qAbupO/Vv/0L5LVg0wE/.ycLKDrZ.4BOnPkyS"
	Expiration  = 2 * time.Hour
	Issuer      = "aidanp"
)

var PrivateKey ed25519.PrivateKey

func init() {
	seed := make([]byte, ed25519.SeedSize)
	_, err := rand.Read(seed)
	if err != nil {
		log.Fatal(err)
	}
	PrivateKey = ed25519.NewKeyFromSeed(seed)
}

type LoginHandler struct {
	DB *pgxpool.Pool
}

type Auth struct {
	Token string `json:"token"`
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("received a request", "ipaddr", r.RemoteAddr, "form values", r.Form)
	// TODO: sanitize username
	// TODO: check data type
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

	currentTime := time.Now()
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(Expiration)),
		Issuer:    Issuer,
		IssuedAt:  jwt.NewNumericDate(currentTime),
		NotBefore: jwt.NewNumericDate(currentTime),
	})
	tokenString, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Error("error occurred while trying to sign the jwt", "err", err, "user", username)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}
