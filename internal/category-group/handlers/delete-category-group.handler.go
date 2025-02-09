package category_group_handlers

import "net/http"

func DeleteCategoryGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete CategoryGroup"))
}
