package category_handlers

import "net/http"

func ListCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of categories"))
}
