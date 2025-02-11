package category_group_handlers

import (
	"encoding/json"
	"net/http"
)

func RenameCategoryGroup(w http.ResponseWriter, r *http.Request) {
	var requestBody struct{ Name string }
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
}
