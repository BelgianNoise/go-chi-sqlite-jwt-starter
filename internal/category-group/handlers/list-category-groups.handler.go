package category_group_handlers

import "net/http"

func ListCategoryGroups(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of category groups"))
}
