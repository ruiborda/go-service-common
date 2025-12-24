package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/product"
)

type ProductService interface {
	CreateProduct(request *product.CreateProductRequest) *dto.Response[*product.CreateProductResponse]
	GetProductById(productId string) *dto.Response[*product.GetProductByIdResponse]
	GetAllProducts() *dto.Response[[]*product.GetAllProductsResponse]
	UpdateProduct(request *product.UpdateProductRequest) *dto.Response[*product.UpdateProductResponse]
	DeleteProduct(productId string) *dto.Response[interface{}]
	GetProductsPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*product.GetProductsPaginatedResponse]]
}
