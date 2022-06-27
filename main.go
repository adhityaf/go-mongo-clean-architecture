package main

import (
	"go-fiber-clean-arch/config"
	"go-fiber-clean-arch/controller"
	"go-fiber-clean-arch/exception"
	"go-fiber-clean-arch/repository"
	"go-fiber-clean-arch/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	app := fiber.New(*config.NewFiberConfig())
	app.Use(recover.New())

	userController.Route(app)

	err:= app.Listen(":3000")
	exception.PanicIfErr(err)
}
