package service

import (
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/dto/user"
)

type UserService interface {
	CreateUser(request *user.CreateUserRequest) *dto.Response[*user.CreateUserResponse]
	GetUserById(id string) *dto.Response[*user.GetUserByIdResponse]
	GetUserByEmail(email string) *dto.Response[*user.GetUserByIdResponse]
	GetAllUsers() *dto.Response[[]*user.GetUserByIdResponse]
	UpdateUserById(request *user.UpdateUserRequest) *dto.Response[*user.UpdateUserResponse]
	DeleteUserById(id string) *dto.Response[interface{}]
	FindAllUsersPaginated(pageRequest *dto.PageRequest) *dto.Response[*dto.PageBody[*user.GetUserByIdResponse]]
}
