package service

import (
	"go-fiber-clean-arch/entity"
	"go-fiber-clean-arch/model"
	"go-fiber-clean-arch/repository"
	"go-fiber-clean-arch/validation"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) Create(request model.RegisterRequest) (response model.RegisterResponse) {
	validation.ValidateCreateUserRequest(request)

	user := entity.User{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	service.UserRepository.Create(user)

	response = model.RegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return response
}

func (service *userServiceImpl) List() (responses []model.GetUserResponse) {
	users := service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return responses
}

func(service *userServiceImpl) FindById(id string)(response model.GetUserResponse){
	user := service.UserRepository.FindById(id)
	response = model.GetUserResponse{
		Id: user.Id,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	return response
}
