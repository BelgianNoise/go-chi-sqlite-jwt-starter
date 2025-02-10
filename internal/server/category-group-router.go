package server

import (
	"context"
	"net/http"

	"gofinn/internal/auth"
	category_group_handlers "gofinn/internal/category-group/handlers"
	models "gofinn/internal/models"
	provider "gofinn/internal/provider"
	"gofinn/internal/utils"
	"gofinn/internal/validation"

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
		r.Put("/", category_group_handlers.UpdateCategoryGroup)
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

		catgoryGroup, err := provider.Provider.CategoryGroupService.GetCategoryGroup(id)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		user := r.Context().Value(models.ContextKeys.User).(models.User)
		err = validation.HasAccessToCategoryGroup(w, catgoryGroup.ID, user.ID)
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), models.ContextKeys.CategoryGroup, catgoryGroup)
		ctx = context.WithValue(ctx, models.ContextKeys.CategoryGroupID, categoryGroupID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
