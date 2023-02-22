package crm_handlers

import (
	crm_services "main/service/Documents"

	"github.com/gofiber/fiber/v2"
)

type DocumentsHandler struct {
	Service crm_services.Service_Documents
}

func (h DocumentsHandler) Documents(c *fiber.Ctx) error {
	return nil
}

func (h DocumentsHandler) DocumentsCreate(c *fiber.Ctx) error {
	return nil
}
func (h DocumentsHandler) DocumentsById(c *fiber.Ctx) error {
	return nil
}
func (h DocumentsHandler) DocumentsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h DocumentsHandler) DocumentsDelete(c *fiber.Ctx) error {
	return nil
}
