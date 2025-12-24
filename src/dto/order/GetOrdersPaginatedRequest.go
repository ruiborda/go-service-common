package order

import "github.com/ruiborda/go-service-common/dto"

type GetOrdersPaginatedRequest struct {
	dto.PageRequest
	StartDate *string `json:"startDate" form:"startDate"`
	EndDate   *string `json:"endDate" form:"endDate"`
}
