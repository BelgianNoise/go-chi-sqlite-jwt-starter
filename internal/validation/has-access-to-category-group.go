package validation

import (
	"fmt"
	"gofinn/internal/provider"
	"net/http"
)

func HasAccessToCategoryGroup(
	w http.ResponseWriter,
	groupID int64,
	userID int64,
) error {
	categoryGroup, err := provider.Provider.CategoryGroupService.GetCategoryGroup(groupID)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return err
	}
	if categoryGroup.OwnerID != userID {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
		return fmt.Errorf("User does not have access to category group")
	}
	return nil
}
