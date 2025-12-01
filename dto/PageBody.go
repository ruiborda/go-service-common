package dto

type PageBody[T any] struct {
	Items       []T `json:"items"`
	TotalItems  int `json:"totalItems"`
	TotalPages  int `json:"totalPages"`
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"` // it's request page size
}

func PageBodyBuilder[T any]() *PageBody[T] {
	return &PageBody[T]{
		Items:       []T{},
		CurrentPage: 0,
		TotalItems:  0,
		TotalPages:  0,
	}
}

func (p *PageBody[T]) SetItems(items []T) *PageBody[T] {
	p.Items = items
	return p
}

func (p *PageBody[T]) SetPageSize(pageSize int) *PageBody[T] {
	p.PageSize = pageSize
	return p
}

func (p *PageBody[T]) SetCurrentPage(currentPage int) *PageBody[T] {
	if currentPage < 1 {
		currentPage = 1
	}
	p.CurrentPage = currentPage
	return p
}

func (p *PageBody[T]) SetTotalItems(totalItems int) *PageBody[T] {
	if totalItems < 0 {
		totalItems = 1
	}
	p.TotalItems = totalItems
	if p.PageSize > 0 {
		p.TotalPages = (totalItems + p.PageSize - 1) / p.PageSize
	} else {
		p.TotalPages = 0
	}
	return p
}
