package provider

import (
	category_service "gofinn/internal/category/service"
	user_service "gofinn/internal/user/service"
	"log"
)

type ServiceProvider struct {
	CategoryService category_service.CategoryService
	UserService     user_service.UserService
}

var Provider *ServiceProvider

func Initialize() {
	log.Printf("Initializing service provider...")
	defer log.Printf("Service provider initialized!")

	Provider = &ServiceProvider{
		CategoryService: category_service.NewSQLiteCategoryService(),
		UserService:     user_service.NewSQLiteUserService(),
	}
}
