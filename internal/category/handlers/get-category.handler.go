package category_handlers

import (
	"encoding/json"
	"net/http"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	category, ok := getCategoryFromContext(w, r)
	if !ok {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
