package order

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/order"
)

func (this *OrderServiceImpl) GetOrdersPaginated(request *order.GetOrdersPaginatedRequest) *dto.Response[*dto.PageBody[*order.GetOrdersPaginatedResponse]] {
	page := *request.PageNumber
	size := *request.PageSize
	zeroBasedPage := page - 1
	if zeroBasedPage < 0 {
		zeroBasedPage = 0
	}

	var start, end *time.Time
	if request.StartDate != nil && *request.StartDate != "" {
		t, err := time.Parse("2006-01-02", *request.StartDate)
		if err == nil {
			start = &t
		} else {
			slog.Warn("Failed to parse StartDate", "value", *request.StartDate, "error", err)
		}
	}
	if request.EndDate != nil && *request.EndDate != "" {
		t, err := time.Parse("2006-01-02", *request.EndDate)
		if err == nil {
			end = &t
		} else {
			slog.Warn("Failed to parse EndDate", "value", *request.EndDate, "error", err)
		}
	}

	orderModels, err := this.orderRepository.FindAllByPageAndSize(zeroBasedPage, size, request.Sort, start, end, request.GetFilters())
	if err != nil {
		slog.Error("Failed to get paginated orders", "error", err)
		return dto.ResponseBuilder[*dto.PageBody[*order.GetOrdersPaginatedResponse]](http.StatusInternalServerError, nil).SetMessage("Error retrieving paginated orders")
	}

	totalItems, err := this.orderRepository.Count(start, end, request.GetFilters())
	if err != nil {
		slog.Error("Failed to count orders", "error", err)
		return dto.ResponseBuilder[*dto.PageBody[*order.GetOrdersPaginatedResponse]](http.StatusInternalServerError, nil).SetMessage("Error counting orders")
	}

	enrichedOrders, _ := this.enrichOrdersWithCustomers(orderModels)
	enrichedOrders, _ = this.enrichOrdersWithTables(enrichedOrders)
	enrichedOrders, _ = this.enrichOrdersWithPaymentTypes(enrichedOrders)
	enrichedOrders, _ = this.enrichOrdersWithUsers(enrichedOrders)
	responseDtos := this.orderMapper.ModelsToPaginatedResponse(enrichedOrders)

	pageBody := dto.PageBodyBuilder[*order.GetOrdersPaginatedResponse]().
		SetItems(responseDtos).
		SetCurrentPage(page).
		SetPageSize(size).
		SetTotalItems(int(totalItems))

	return dto.ResponseBuilder(http.StatusOK, pageBody).SetMessage("Paginated orders retrieved successfully")
}
