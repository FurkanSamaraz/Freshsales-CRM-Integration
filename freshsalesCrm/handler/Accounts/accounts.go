package crm_handlers

import (
	crm_services "main/service/Accounts"

	"github.com/gofiber/fiber/v2"
)

type AccountsHandler struct {
	Service crm_services.Service_Accounts
}

func (h AccountsHandler) Accounts(c *fiber.Ctx) error {
	return nil
}

func (h AccountsHandler) AccountsCreate(c *fiber.Ctx) error {
	return nil
}
func (h AccountsHandler) AccountsById(c *fiber.Ctx) error {
	return nil
}
func (h AccountsHandler) AccountsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h AccountsHandler) AccountsDelete(c *fiber.Ctx) error {
	return nil
}
