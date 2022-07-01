package service

import "go-fiber-clean-arch/model"

type UserService interface{
	Login(request model.LoginRequest)(response model.LoginResponse)
	Create(request model.RegisterRequest)(response model.RegisterResponse)
	List()(responses []model.GetUserResponse)
}