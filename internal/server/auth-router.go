package server

import (
	"net/http"

	auth_handlers "go-chi-sqlite-jwt-starter/internal/auth/handlers"

	"github.com/go-chi/chi/v5"
)

func authRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", auth_handlers.RegisterAccount)
	r.Post("/login", auth_handlers.Login)

	return r
}
