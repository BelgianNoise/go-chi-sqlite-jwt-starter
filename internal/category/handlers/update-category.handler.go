package category_handlers

import "net/http"

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Category"))
}
