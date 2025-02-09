package provider

import (
	category_group_service "gofinn/internal/category-group/service"
	category_service "gofinn/internal/category/service"
	user_service "gofinn/internal/user/service"
	"log"
)

type ServiceProvider struct {
	CategoryService      category_service.CategoryService
	CategoryGroupService category_group_service.CategoryGroupService
	UserService          user_service.UserService
}

var Provider *ServiceProvider

func Initialize() {
	log.Printf("Initializing service provider...")
	defer log.Printf("Service provider initialized!")

	Provider = &ServiceProvider{
		CategoryService:      category_service.NewSQLiteCategoryService(),
		CategoryGroupService: category_group_service.NewSQLiteCategoryGroupService(),
		UserService:          user_service.NewSQLiteUserService(),
	}
}
