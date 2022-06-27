package config

import (
	"go-fiber-clean-arch/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() *fiber.Config{
	return &fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
