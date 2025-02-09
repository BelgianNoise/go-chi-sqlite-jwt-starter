package category_group_handlers

import (
	"encoding/json"
	"gofinn/internal/models"
	"gofinn/internal/provider"
	"gofinn/internal/utils"
	"net/http"
)

func CreateCategoryGroup(w http.ResponseWriter, r *http.Request) {
	var requestBody struct{ Name string }
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := utils.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	groupFields := models.CategoryGroupFields{
		Name:    requestBody.Name,
		OwnerID: userID,
	}
	createdCategoryGroup, err := provider.Provider.CategoryGroupService.CreateCategoryGroup(groupFields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategoryGroup)
}
