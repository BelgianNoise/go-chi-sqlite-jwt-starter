package category_handlers

import (
	"encoding/json"
	"gofinn/internal/models"
	"gofinn/internal/provider"
	"gofinn/internal/utils"
	"gofinn/internal/validation"
	"net/http"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	category := r.Context().Value(models.ContextKeys.Category).(models.Category)
	user := utils.GetUserFromContext(w, r.Context())
	err := validation.HasAccessToCategoryGroup(w, category.CategoryGroupID, user.ID)
	if err != nil {
		return
	}

	err = provider.Provider.CategoryService.DeleteCategory(category.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{ id int64 }{
		id: category.ID,
	})
}
