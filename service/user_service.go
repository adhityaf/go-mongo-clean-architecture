package service

import "go-fiber-clean-arch/model"

type UserService interface{
	Create(request model.RegisterRequest)(response model.RegisterResponse)
	List()(responses []model.GetUserResponse)
	FindById(id string)(response model.GetUserResponse)
}