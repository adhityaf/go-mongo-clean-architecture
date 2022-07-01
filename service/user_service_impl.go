package service

import (
	"go-fiber-clean-arch/entity"
	"go-fiber-clean-arch/helpers"
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

func (service *userServiceImpl) Login(request model.LoginRequest) (response model.LoginResponse) {
	validation.ValidateLoginUserRequest(request)

	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		// If email does not exist
		// response.Token = ""
		return response
	}

	isOk := helpers.ComparePass([]byte(user.Password), []byte(request.Password))
	if !isOk {
		// If db password and request password not match
		return response
	}

	response = model.LoginResponse{
		Token: helpers.GenerateToken(user.Id, user.Email, user.Username),
	}

	return response
}

func (service *userServiceImpl) Create(request model.RegisterRequest) (response model.RegisterResponse) {
	validation.ValidateCreateUserRequest(request)

	user := entity.User{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: helpers.HashPass(request.Password),
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
		})
	}
	return responses
}
