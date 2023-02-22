package crm_handlers

import (
	crm_services "main/service/Leads"

	"github.com/gofiber/fiber/v2"
)

type LeadsHandler struct {
	Service crm_services.Service_Leads
}

func (h LeadsHandler) Leads(c *fiber.Ctx) error {
	return nil
}

func (h LeadsHandler) LeadsCreate(c *fiber.Ctx) error {
	return nil
}
func (h LeadsHandler) LeadsById(c *fiber.Ctx) error {
	return nil
}
func (h LeadsHandler) LeadsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h LeadsHandler) LeadsDelete(c *fiber.Ctx) error {
	return nil
}
