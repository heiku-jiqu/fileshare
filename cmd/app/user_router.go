package main

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	json.NewDecoder(r.Body).Decode(&login)
	if login.Password != "securepassword" || login.Username != "secureusername" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     "username",
		Value:    "secureusername",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func NewUserRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", LoginHandler)
	return mux
}
