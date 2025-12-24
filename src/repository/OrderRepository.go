package repository

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/model"
	"time"
)

type OrderRepository interface {
	Create(order *model.Order) (*model.Order, error)
	FindById(id string) (*model.Order, error)
	Update(order *model.Order) (*model.Order, error)
	Delete(id string) error
	FindAllByPageAndSize(page, size int, sorts []*dto.PageRequestOrder, startDate, endDate *time.Time, filters []*dto.PageRequestFilter) ([]*model.Order, error)
	Count(startDate, endDate *time.Time, filters []*dto.PageRequestFilter) (int64, error)
	GetNextDailyOrderNumber() (int, error)
	GetNextInvoiceNumber() (int64, error)
}
