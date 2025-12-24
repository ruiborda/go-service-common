package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/category"
)

type CategoryService interface {
	CreateCategory(request *category.CreateCategoryRequest) *dto.Response[*category.CreateCategoryResponse]
	GetCategoryById(categoryId string) *dto.Response[*category.GetCategoryByIdResponse]
	UpdateCategory(request *category.UpdateCategoryRequest) *dto.Response[*category.UpdateCategoryResponse]
	DeleteCategory(categoryId string) *dto.Response[interface{}]
	GetCategoriesPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*category.GetCategoriesPaginatedResponse]]
	GetAllCategories() *dto.Response[[]*category.GetAllCategoriesResponse]
}
