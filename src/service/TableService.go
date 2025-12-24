package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/table"
)

type TableService interface {
	CreateTable(request *table.CreateTableRequest) *dto.Response[*table.CreateTableResponse]
	GetTableById(tableId string) *dto.Response[*table.GetTableByIdResponse]
	UpdateTable(request *table.UpdateTableRequest) *dto.Response[*table.UpdateTableResponse]
	DeleteTable(tableId string) *dto.Response[interface{}]
	GetTablesPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*table.GetTablesPaginatedResponse]]
	GetAllTables() *dto.Response[[]*table.GetAllTablesResponse]
}
