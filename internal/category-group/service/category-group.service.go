package category_group_service

import (
	"go-chi-sqlite-jwt-starter/internal/models"
)

type CategoryGroupService interface {
	ListCategoryGroupsForUser(ownerID int64) ([]models.CategoryGroup, error)
	CreateCategoryGroup(categoryGroup models.CategoryGroupFields) (models.CategoryGroup, error)
	GetCategoryGroup(categoryGroupID int64) (models.CategoryGroup, error)
	GetCategoryGroupForUser(categoryGroupID int64, ownerID int64) (models.CategoryGroup, error)
	UpdateCategoryGroupName(id int64, name string) (models.CategoryGroup, error)
	DeleteCategoryGroup(id int64) error
}
