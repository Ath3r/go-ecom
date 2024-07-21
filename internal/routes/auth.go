package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)



func AuthRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/signup", handleSignup)
	r.Get("/login", handleLogin)
	r.Get("/refresh", handleRefresh)
	r.Get("/logout", handleLogout)
	return r
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup"))
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("refresh"))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}
