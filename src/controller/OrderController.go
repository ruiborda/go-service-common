package controller

import (
	"net/http"

	"github.com/ruiborda/go-jwt/src/domain/entity"
	order2 "github.com/ruiborda/pos-api/src/service/order"

	"github.com/gin-gonic/gin"
	common_dto "github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/go-swagger-generator/v2/src/openapi"
	"github.com/ruiborda/go-swagger-generator/v2/src/openapi_spec/mime"
	"github.com/ruiborda/go-swagger-generator/v2/src/swagger"
	"github.com/ruiborda/pos-api/src/dto/order"
	"github.com/ruiborda/pos-api/src/service"
)

var _ = swagger.Swagger().Tag("order", func(tag openapi.Tag) {
	tag.Description("Operaciones sobre ordenes")
})

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderService: order2.NewOrderServiceImpl(),
	}
}

var _ = swagger.Swagger().Path("/api/v1/orders/pages").
	Get(func(op openapi.Operation) {
		op.Summary("Obtener ordenes con paginación").
			OperationID("getOrdersPaginated").
			Tag("order").
			QueryParameter("pageNumber", func(p openapi.Parameter) {
				p.Description("Número de página").Required(false).
					Schema(func(s openapi.Schema) { s.Type("integer").Default(1) })
			}).
			QueryParameter("pageSize", func(p openapi.Parameter) {
				p.Description("Número de elementos por página").Required(false).
					Schema(func(s openapi.Schema) { s.Type("integer").Default(10) })
			}).
			QueryParameter("startDate", func(p openapi.Parameter) {
				p.Description("Fecha de inicio").Required(false).Schema(func(s openapi.Schema) { s.Type("string").Format("date") })
			}).
			QueryParameter("endDate", func(p openapi.Parameter) {
				p.Description("Fecha de fin").Required(false).Schema(func(s openapi.Schema) { s.Type("string").Format("date") })
			}).
			Response(http.StatusOK, func(r openapi.Response) {
				r.Description("Lista paginada de ordenes").
					Content(mime.ApplicationJSON, func(mt openapi.MediaType) {
						mt.SchemaFromDTO(&common_dto.Response[*common_dto.PageBody[*order.GetOrdersPaginatedResponse]]{})
					})
			})
	}).
	Doc()

func (oc *OrderController) GetOrdersPaginated(c *gin.Context) {
	var request = &order.GetOrdersPaginatedRequest{}
	if err := c.ShouldBind(request); err != nil {
		response := common_dto.ResponseBuilder[*common_dto.PageBody[*order.GetOrdersPaginatedResponse]](http.StatusBadRequest, nil).
			SetMessage("Invalid request parameters")
		c.JSON(response.Status, response)
		return
	}

	request.PageRequest = *common_dto.DefaultPageRequest(&request.PageRequest)
	response := oc.orderService.GetOrdersPaginated(request)
	c.JSON(response.Status, response)
}

// ... Rest of the controllers (CreateOrder, GetOrderById, etc) remains the same as previously defined
