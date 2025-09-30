package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"golang.org/x/crypto/bcrypt"
)

const hardcodedUser = "user"
const hardcodedPassword = "hardcodedpw"

// Key used to retrieve current User in the session from a context object
const UserContextKey = "UserContextKey"

type Login struct {
	hardcodedUser     string
	hardcodedPassword string
	sessionManager    *scs.SessionManager
}

func NewLogin(sess *scs.SessionManager) Login {
	return Login{
		hardcodedUser:     hardcodedUser,
		hardcodedPassword: hardcodedPassword,
		sessionManager:    sess,
	}
}

func (app Login) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	user := r.PostFormValue("user")
	password := r.PostFormValue("password")

	if user != app.hardcodedUser {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(app.hardcodedPassword), 14)
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
		return
	}
	app.sessionManager.Put(r.Context(), UserContextKey, user)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CreateUserPostHandler(w http.ResponseWriter, r *http.Request) {
	user := r.PostFormValue("user")
	password := r.PostFormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO: Store user & hash
	log.Print(user, hash)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

}
