package category_handlers

import (
	"encoding/json"
	"net/http"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	category := getCategoryFromContext(w, r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
