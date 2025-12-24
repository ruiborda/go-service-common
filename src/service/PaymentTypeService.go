package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/payment_type"
)

type PaymentTypeService interface {
	CreatePaymentType(request *payment_type.CreatePaymentTypeRequest) *dto.Response[*payment_type.CreatePaymentTypeResponse]
	GetPaymentTypeById(paymentTypeId string) *dto.Response[*payment_type.GetPaymentTypeByIdResponse]
	UpdatePaymentType(request *payment_type.UpdatePaymentTypeRequest) *dto.Response[*payment_type.UpdatePaymentTypeResponse]
	DeletePaymentType(paymentTypeId string) *dto.Response[interface{}]
	GetPaymentTypesPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*payment_type.GetPaymentTypesPaginatedResponse]]
	GetAllPaymentTypes() *dto.Response[[]*payment_type.GetAllPaymentTypesResponse]
}
