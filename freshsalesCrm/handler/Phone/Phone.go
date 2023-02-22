package crm_handlers

import (
	crm_services "main/service/Phone"

	"github.com/gofiber/fiber/v2"
)

type PhoneHandler struct {
	Service crm_services.Service_Phone
}

func (h PhoneHandler) Phone(c *fiber.Ctx) error {
	return nil
}

func (h PhoneHandler) PhoneCreate(c *fiber.Ctx) error {
	return nil
}
func (h PhoneHandler) PhoneById(c *fiber.Ctx) error {
	return nil
}
func (h PhoneHandler) PhoneUpdate(c *fiber.Ctx) error {
	return nil
}
func (h PhoneHandler) PhoneDelete(c *fiber.Ctx) error {
	return nil
}
