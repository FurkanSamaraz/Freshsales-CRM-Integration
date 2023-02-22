package crm_handlers

import (
	crm_integration "main/package"
	crm_services "main/service/Activities"

	"github.com/gofiber/fiber/v2"
)

type ActivitiesHandler struct {
	Service crm_services.Service_Activities
}

type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func (h ActivitiesHandler) Activities(c *fiber.Ctx) error {
	return nil
}

func (h ActivitiesHandler) ActivitiesCreate(c *fiber.Ctx) error {
	return nil
}

func (h ActivitiesHandler) ActivitiesById(c *fiber.Ctx) error {
	return nil
}

func (h ActivitiesHandler) ActivitiesUpdate(c *fiber.Ctx) error {
	return nil
}

func (h ActivitiesHandler) ActivitiesDelete(c *fiber.Ctx) error {
	return nil
}
