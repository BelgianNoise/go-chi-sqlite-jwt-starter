package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(adminOnly)

	r.Get("/test-token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are successfully authenticated as an admin!"))
	})

	return r
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if true {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
