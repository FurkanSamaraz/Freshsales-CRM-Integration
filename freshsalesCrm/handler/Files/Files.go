package crm_handlers

import (
	crm_services "main/service/Files"

	"github.com/gofiber/fiber/v2"
)

type FilesHandler struct {
	Service crm_services.Service_Files
}

func (h FilesHandler) Files(c *fiber.Ctx) error {
	return nil
}

func (h FilesHandler) FilesCreate(c *fiber.Ctx) error {
	return nil
}
func (h FilesHandler) FilesById(c *fiber.Ctx) error {
	return nil
}
func (h FilesHandler) FilesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h FilesHandler) FilesDelete(c *fiber.Ctx) error {
	return nil
}
