package auth_handlers

import (
	"encoding/json"
	"go-chi-sqlite-jwt-starter/internal/models"
	"go-chi-sqlite-jwt-starter/internal/provider"
	"go-chi-sqlite-jwt-starter/internal/utils"
	"go-chi-sqlite-jwt-starter/internal/validation"
	"net/http"
)

func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var registerCreds models.AuthParams
	if err := json.NewDecoder(r.Body).Decode(&registerCreds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validation.AccountDoesNotExist(w, registerCreds.Username)

	hashedPassword := utils.HashPassword(registerCreds.Password)
	_, err := provider.Provider.UserService.CreateUser(models.UserFields{
		Username:       registerCreds.Username,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
