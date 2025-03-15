package category_group_handlers

import (
	"encoding/json"
	"go-chi-sqlite-jwt-starter/internal/models"
	"go-chi-sqlite-jwt-starter/internal/provider"
	"go-chi-sqlite-jwt-starter/internal/utils"
	"net/http"
)

func CreateCategoryGroup(w http.ResponseWriter, r *http.Request) {
	var requestBody struct{ Name string }
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID := utils.GetUserIDFromContext(w, r.Context())
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
