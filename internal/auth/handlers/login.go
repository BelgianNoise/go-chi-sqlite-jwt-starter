package auth_handlers

import (
	"encoding/json"
	"go-chi-sqlite-jwt-starter/internal/auth"
	"go-chi-sqlite-jwt-starter/internal/models"
	"go-chi-sqlite-jwt-starter/internal/provider"
	"go-chi-sqlite-jwt-starter/internal/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var authParams models.AuthParams
	if err := json.NewDecoder(r.Body).Decode(&authParams); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := provider.Provider.UserService.GetUserByUsername(authParams.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !utils.CompareHashedToNonHashedPassword(user.HashedPassword, authParams.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
