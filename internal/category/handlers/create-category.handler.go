package category_handlers

import (
	"encoding/json"
	"gofinn/internal/models"
	"gofinn/internal/provider"
	"gofinn/internal/utils"
	"gofinn/internal/validation"
	"net/http"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryFields models.CategoryFields
	if err := json.NewDecoder(r.Body).Decode(&categoryFields); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := utils.GetUserFromContext(w, r.Context())
	err := validation.HasAccessToCategoryGroup(w, categoryFields.CategoryGroupID, user.ID)
	if err != nil {
		return
	}

	createdCategory, err := provider.Provider.CategoryService.CreateCategory(categoryFields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)
}
