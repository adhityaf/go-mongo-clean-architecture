package controller

import (
	"go-fiber-clean-arch/exception"
	"go-fiber-clean-arch/model"
	"go-fiber-clean-arch/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/v1/api/register", controller.Create)
	app.Get("/v1/api/users", controller.GetAll)
	app.Get("/v1/api/users/:id", controller.GetById)
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	var request model.RegisterRequest
	err := ctx.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfErr(err)
	// if err != nil {
	// 	return ctx.JSON(model.WebResponse{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "BAD REQUEST",
	// 	})
	// }

	response := controller.UserService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Register an account",
		Data:   response,
	})
}

func (controller *UserController) GetAll(ctx *fiber.Ctx) error {
	responses := controller.UserService.List()
	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "Success retrieve all data",
		Data:   responses,
	})
}

func (controller *UserController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := controller.UserService.FindById(id)
	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   response,
	})
}
