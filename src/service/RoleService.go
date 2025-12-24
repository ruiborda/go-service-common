package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/role"
)

type RoleService interface {
	CreateRole(request *role.CreateRoleRequest) *dto.Response[*role.CreateRoleResponse]
	GetRoleById(id string) *dto.Response[*role.GetRoleByIdResponse]
	UpdateRoleById(request *role.UpdateRoleRequest) *dto.Response[*role.UpdateRoleResponse]
	DeleteRoleById(id string) *dto.Response[interface{}]
	FindAllRolesPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*role.GetRoleByIdResponse]]
}
