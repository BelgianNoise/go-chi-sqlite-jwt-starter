package category_service

import (
	"gofinn/internal/models"
)

type CategoryGroupService interface {
	ListCategoryGroupsForUser(ownerID int64) ([]models.CategoryGroup, error)
	CreateCategoryGroup(categoryGroup models.CategoryGroupFields) (models.CategoryGroup, error)
	GetCategoryGroup(categoryGroupID int64) (models.CategoryGroup, error)
	UpdateCategoryGroup(categoryGroup models.CategoryGroup) (models.CategoryGroup, error)
	DeleteCategoryGroup(categoryGroupID int64) error
}
