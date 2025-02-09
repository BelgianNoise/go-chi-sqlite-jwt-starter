package server

import (
	"context"
	"net/http"

	"gofinn/internal/auth"
	category_handlers "gofinn/internal/category/handlers"
	models "gofinn/internal/models"
	provider "gofinn/internal/provider"
	"gofinn/internal/utils"

	"github.com/go-chi/chi/v5"
)

func categoryRouter() http.Handler {
	r := chi.NewRouter()
	auth.UseAuthMiddleware(r)

	r.Get("/list", category_handlers.ListCategories)
	r.Post("/create", category_handlers.CreateCategory)

	r.Route("/{categoryID}", func(r chi.Router) {
		r.Use(CategoryCtx)
		r.Get("/", category_handlers.GetCategory)
		r.Put("/", category_handlers.UpdateCategory)
		r.Delete("/", category_handlers.DeleteCategory)
	})

	return r
}

func CategoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryID := chi.URLParam(r, "categoryID")
		id, err := utils.StringToInt64(categoryID)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		catgory, err := provider.Provider.CategoryService.GetCategory(id)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), models.ContextKeys.Category, catgory)
		ctx = context.WithValue(ctx, models.ContextKeys.CategoryID, categoryID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
