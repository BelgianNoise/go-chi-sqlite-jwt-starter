package utils

import (
	"context"
	"fmt"
	"gofinn/internal/models"
)

func GetUserFromContext(ctx context.Context) (models.User, error) {
	user, ok := ctx.Value(models.ContextKeys.User).(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("user not found in context")
	}
	return user, nil
}

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
