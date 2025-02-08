package category_service

import (
	"gofinn/internal/models"
)

type CategoryService interface {
	ListCategories() ([]models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	GetCategory(categoryID int) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(categoryID int) error
}
