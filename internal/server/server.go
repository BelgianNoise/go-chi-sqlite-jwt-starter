package server

import (
	"go-chi-sqlite-jwt-starter/internal/auth"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize() *chi.Mux {
	log.Println("Initializing server...")
	defer log.Println("Server initialized")

	r := chi.NewRouter()
	useGlobalMiddleware(r)
	auth.InitializeTokenVerifier()

	r.Mount("/category", categoryRouter())
	r.Mount("/category-group", categoryGroupRouter())
	r.Mount("/admin", adminRouter())
	r.Mount("/auth", authRouter())

	return r
}

func useGlobalMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health"))
}
