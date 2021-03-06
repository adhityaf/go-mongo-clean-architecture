package repository

import "go-fiber-clean-arch/entity"

type UserRepository interface{
	Create(user entity.User)
	FindAll()(users []entity.User)
	FindById(id string)(user *entity.User, err error)
	FindByEmail(email string)(user *entity.User, err error)
}