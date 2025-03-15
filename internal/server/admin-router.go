package server

import (
	"go-chi-sqlite-jwt-starter/internal/auth"
	"go-chi-sqlite-jwt-starter/internal/models"
	"go-chi-sqlite-jwt-starter/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func adminRouter() http.Handler {
	r := chi.NewRouter()
	auth.UseAuthMiddleware(r)
	r.Use(adminOnly)

	r.Get("/test-token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are successfully authenticated as an admin!"))
	})

	return r
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := utils.GetUserFromContext(w, r.Context())
		if user.Role != models.Admin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
