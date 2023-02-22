package crm_handlers

import (
	crm_services "main/service/Users"

	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	Service crm_services.Service_Users
}

func (h UsersHandler) Users(c *fiber.Ctx) error {
	return nil
}
