package server

import (
	"gofinn/internal/auth"
	"gofinn/internal/models"
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
		user := r.Context().Value(models.ContextKeys.User).(models.User)
		if user.Role != models.Admin {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
