package dto

import "github.com/ruiborda/go-service-common/types"

type PageRequest struct {
	PageNumber *int                 `json:"pageNumber" form:"pageNumber"`
	PageSize   *int                 `json:"pageSize" form:"pageSize"`
	Search     *string              `json:"search" form:"search"`
	Sort       *[]*PageRequestOrder `json:"sort" form:"sort"`
}

func DefaultPageRequest(request *PageRequest) *PageRequest {
	if request == nil {
		request = &PageRequest{
			PageNumber: types.Pointer(1),
			PageSize:   types.Pointer(10),
			Search:     nil,
		}
	}
	if request.PageNumber == nil {
		request.PageNumber = types.Pointer(1)
	}
	if request.PageSize == nil {
		request.PageSize = types.Pointer(10)
	}

	return request
}
