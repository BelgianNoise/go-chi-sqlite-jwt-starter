package category_service

import (
	"gofinn/internal/models"
)

type CategoryService interface {
	ListCategoriesForUser(ownerID int64) ([]models.Category, error)
	CreateCategory(category models.CategoryFields) (models.Category, error)
	GetCategory(categoryID int64) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(categoryID int64) error
}
