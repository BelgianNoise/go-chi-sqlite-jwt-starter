package category_handlers

import (
	"gofinn/internal/models"
	"net/http"
)

func getCategoryFromContext(w http.ResponseWriter, r *http.Request) (models.Category, bool) {
	category, ok := r.Context().Value(models.ContextKeys.Category).(models.Category)
	if !ok {
		http.Error(w, "Category not found", http.StatusNotFound)
	}
	return category, ok
}
