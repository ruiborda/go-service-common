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
		return &PageRequest{
			PageNumber: types.Pointer(1),
			PageSize:   types.Pointer(10),
			Search:     nil,
			Sort:       nil,
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

func (pr *PageRequest) GetPageNumber() int {
	if pr.PageNumber == nil {
		return 1
	}
	return *pr.PageNumber
}

func (pr *PageRequest) GetPageSize() int {
	if pr.PageSize == nil {
		return 10
	}
	return *pr.PageSize
}

func (pr *PageRequest) GetSearch() *string {
	if pr.Search == nil {
		return nil
	}
	return pr.Search
}
