package provider

import (
	category_service "gofinn/internal/category/service"
	"log"
)

type ServiceProvider struct {
	CategoryService category_service.CategoryService
}

var Provider *ServiceProvider

func Initialize() {
	log.Printf("Initializing service provider...")
	defer log.Printf("Service provider initialized!")

	Provider = &ServiceProvider{
		CategoryService: category_service.NewSQLiteCategoryService(),
	}
}
