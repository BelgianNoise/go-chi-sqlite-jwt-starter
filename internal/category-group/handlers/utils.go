package category_group_handlers

import (
	"go-chi-sqlite-jwt-starter/internal/models"
	"net/http"
)

func getCategoryGroupFromContext(w http.ResponseWriter, r *http.Request) (models.CategoryGroup, bool) {
	categoryGroup, ok := r.Context().Value(models.ContextKeys.CategoryGroup).(models.CategoryGroup)
	if !ok {
		http.Error(w, "Category group not found", http.StatusNotFound)
	}
	return categoryGroup, ok
}
