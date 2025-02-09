package validation

import (
	"gofinn/internal/provider"
	"net/http"
)

func AccountDoesNotExist(w http.ResponseWriter, username string) {
	user, _ := provider.Provider.UserService.GetUserByUsername(username)
	if user.ID != 0 {
		http.Error(w, "Account already exists", http.StatusConflict)
		return
	}
}
