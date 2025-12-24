package controller

import (
	"github.com/gin-gonic/gin"
	common_dto "github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/go-swagger-generator/v2/src/openapi"
	"github.com/ruiborda/go-swagger-generator/v2/src/openapi_spec/mime"
	"github.com/ruiborda/go-swagger-generator/v2/src/swagger"
	"github.com/ruiborda/pos-api/src/dto/category"
	"github.com/ruiborda/pos-api/src/service"
	"github.com/ruiborda/pos-api/src/service/impl"
	"net/http"
)

var _ = swagger.Swagger().Tag("category", func(tag openapi.Tag) {
	tag.Description("Operaciones sobre categorías de productos")
})

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: impl.NewCategoryServiceImpl(),
	}
}

var _ = swagger.Swagger().Path("/api/v1/categories").
	Post(func(op openapi.Operation) {
		op.Summary("Crear una nueva categoría").
			OperationID("createCategory").
			Tag("category").
			RequestBody(func(rb openapi.RequestBody) {
				rb.Description("Datos de la categoría a crear").Required(true).
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&category.CreateCategoryRequest{})
					})
			}).
			Response(http.StatusCreated, func(r openapi.Response) {
				r.Description("Categoría creada exitosamente").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[*category.CreateCategoryResponse]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var request category.CreateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common_dto.ResponseBuilder[any](http.StatusBadRequest, nil).SetMessage("Invalid request format"))
		return
	}
	response := cc.categoryService.CreateCategory(&request)
	c.JSON(response.Status, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories/{id}").
	Get(func(op openapi.Operation) {
		op.Summary("Obtener una categoría por su ID").
			OperationID("getCategoryById").
			Tag("category").
			PathParameter("id", func(p openapi.Parameter) {
				p.Description("ID de la categoría").Required(true).Schema(func(s openapi.Schema) { s.Type("string").Format("uuid") })
			}).
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Categoría encontrada").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[*category.GetCategoryByIdResponse]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	response := cc.categoryService.GetCategoryById(id)
	c.JSON(response.Status, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories/all").
	Get(func(op openapi.Operation) {
		op.Summary("Obtener todas las categorías").
			OperationID("getAllCategories").
			Tag("category").
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Lista de todas las categorías").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[[]*category.GetAllCategoriesResponse]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	response := cc.categoryService.GetAllCategories()
	c.JSON(response.Status, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories").
	Put(func(op openapi.Operation) {
		op.Summary("Actualizar una categoría existente").
			OperationID("updateCategory").
			Tag("category").
			RequestBody(func(rb openapi.RequestBody) {
				rb.Description("Datos de la categoría a actualizar").Required(true).
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&category.UpdateCategoryRequest{})
					})
			}).
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Categoría actualizada exitosamente").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[category.UpdateCategoryResponse]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var request category.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common_dto.ResponseBuilder[any](http.StatusBadRequest, nil).SetMessage("Invalid request format"))
		return
	}
	response := cc.categoryService.UpdateCategory(&request)
	c.JSON(response.Status, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories/{id}").
	Delete(func(op openapi.Operation) {
		op.Summary("Eliminar una categoría por su ID").
			OperationID("deleteCategory").
			Tag("category").
			PathParameter("id", func(p openapi.Parameter) {
				p.Description("ID de la categoría a eliminar").Required(true).Schema(func(s openapi.Schema) { s.Type("string").Format("uuid") })
			}).
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Categoría eliminada").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[any]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	response := cc.categoryService.DeleteCategory(id)
	c.JSON(response.Status, response)
}

var _ = swagger.Swagger().Path("/api/v1/categories/pages").
	Get(func(op openapi.Operation) {
		op.Summary("Obtener categorías con paginación").
			OperationID("getCategoriesPaginated").
			Tag("category").
			QueryParameter("pageNumber", func(p openapi.Parameter) {
				p.Description("Número de página").Required(false).
					Schema(func(s openapi.Schema) { s.Type("integer").Default(1) })
			}).
			QueryParameter("pageSize", func(p openapi.Parameter) {
				p.Description("Número de elementos por página").Required(false).
					Schema(func(s openapi.Schema) { s.Type("integer").Default(10) })
			}).
			QueryParameter("search", func(p openapi.Parameter) {
				p.Description("Término de búsqueda").Required(false).
					Schema(func(s openapi.Schema) { s.Type("string") })
			}).
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Lista paginada de categorías").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[*common_dto.PageBody[*category.GetCategoriesPaginatedResponse]]{})
					})
			})
	}).
	Doc()

func (cc *CategoryController) GetCategoriesPaginated(c *gin.Context) {
	var pageRequest = &common_dto.PageRequest[any]{}
	if err := c.ShouldBindQuery(pageRequest); err != nil {
		response := common_dto.ResponseBuilder[*common_dto.PageBody[*category.GetCategoriesPaginatedResponse]](http.StatusBadRequest, nil).
			SetMessage("Invalid query parameters")
		c.JSON(response.Status, response)
		return
	}
	pageRequest = common_dto.DefaultPageRequest[any](pageRequest)
	response := cc.categoryService.GetCategoriesPaginated(pageRequest)
	c.JSON(response.Status, response)
}
