package auth_handlers

import (
	"encoding/json"
	"gofinn/internal/models"
	"gofinn/internal/provider"
	"gofinn/internal/utils"
	"gofinn/internal/validation"
	"net/http"
)

func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var registerCreds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&registerCreds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validation.AccountDoesNotExist(w, registerCreds.Username)

	hashedPassword := utils.HashPassword(registerCreds.Password)
	_, err := provider.Provider.UserService.CreateUser(models.UserFields{
		Username:       registerCreds.Username,
		HashedPassword: hashedPassword,
		Currency:       "EUR",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
