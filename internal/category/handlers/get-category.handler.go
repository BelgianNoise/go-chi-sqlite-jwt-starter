package category_handlers

import "net/http"

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Category"))
}
