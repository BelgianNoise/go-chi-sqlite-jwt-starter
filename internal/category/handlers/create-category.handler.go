package category_handlers

import (
	"encoding/json"
	"gofinn/internal/models"
	"gofinn/internal/provider"
	"gofinn/internal/validation"
	"net/http"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryFields models.CategoryFields
	if err := json.NewDecoder(r.Body).Decode(&categoryFields); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validation.HasAccessToCategoryGroup(categoryFields.CategoryGroupID)

	createdCategory, err := provider.Provider.CategoryService.CreateCategory(categoryFields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)
}
