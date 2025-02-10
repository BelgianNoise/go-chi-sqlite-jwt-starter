package utils

import (
	"context"
	"gofinn/internal/models"
	"net/http"
)

func GetUserFromContext(w http.ResponseWriter, ctx context.Context) models.User {
	user, ok := ctx.Value(models.ContextKeys.User).(models.User)
	if !ok {
		http.Error(w, "user was not found in context, you want to try to login again", http.StatusForbidden)
		panic("[ERR-5864] user was not found in context")
	}
	return user
}

func GetUserIDFromContext(w http.ResponseWriter, ctx context.Context) int64 {
	user := GetUserFromContext(w, ctx)
	return user.ID
}
