package category_handlers

import "net/http"

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Category"))
}
