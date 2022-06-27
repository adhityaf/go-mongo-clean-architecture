package validation

import (
	"go-fiber-clean-arch/exception"
	"go-fiber-clean-arch/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateCreateUserRequest(request model.RegisterRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required, validation.Length(8, 0)),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
