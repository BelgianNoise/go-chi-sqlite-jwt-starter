package validation

import (
	"fmt"
	"go-chi-sqlite-jwt-starter/internal/provider"
	"net/http"
)

func HasAccessToCategoryGroup(
	w http.ResponseWriter,
	groupID int64,
	userID int64,
) error {
	categoryGroup, err := provider.Provider.CategoryGroupService.GetCategoryGroupForUser(groupID, userID)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return err
	}
	if categoryGroup.OwnerID != userID {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return fmt.Errorf("user does not have access to category group")
	}
	return nil
}
