package server

import (
	"context"
	"net/http"

	"go-chi-sqlite-jwt-starter/internal/auth"
	category_group_handlers "go-chi-sqlite-jwt-starter/internal/category-group/handlers"
	models "go-chi-sqlite-jwt-starter/internal/models"
	provider "go-chi-sqlite-jwt-starter/internal/provider"
	"go-chi-sqlite-jwt-starter/internal/utils"

	"github.com/go-chi/chi/v5"
)

func categoryGroupRouter() http.Handler {
	r := chi.NewRouter()
	auth.UseAuthMiddleware(r)

	r.Get("/list", category_group_handlers.ListCategoryGroups)
	r.Post("/create", category_group_handlers.CreateCategoryGroup)

	r.Route("/{categoryGroupID}", func(r chi.Router) {
		r.Use(CategoryGroupCtx)
		r.Get("/", category_group_handlers.GetCategoryGroup)
		r.Post("/rename", category_group_handlers.RenameCategoryGroup)
		r.Delete("/", category_group_handlers.DeleteCategoryGroup)
	})

	return r
}

func CategoryGroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryGroupID := chi.URLParam(r, "categoryGroupID")
		id, err := utils.StringToInt64(categoryGroupID)
		if err != nil {
			http.Error(w, "Invalid category group ID", http.StatusBadRequest)
			return
		}

		user := utils.GetUserFromContext(w, r.Context())
		catgoryGroup, err := provider.Provider.CategoryGroupService.GetCategoryGroupForUser(id, user.ID)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), models.ContextKeys.CategoryGroup, catgoryGroup)
		ctx = context.WithValue(ctx, models.ContextKeys.CategoryGroupID, categoryGroupID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
