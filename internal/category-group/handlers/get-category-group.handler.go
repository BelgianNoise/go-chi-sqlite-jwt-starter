package category_group_handlers

import (
	"encoding/json"
	"net/http"
)

func GetCategoryGroup(w http.ResponseWriter, r *http.Request) {
	categoryGroup, ok := getCategoryGroupFromContext(w, r)
	if !ok {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoryGroup)
}
