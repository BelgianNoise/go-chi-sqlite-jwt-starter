package category_group_handlers

import "net/http"

func UpdateCategoryGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update CategoryGroup"))
}
