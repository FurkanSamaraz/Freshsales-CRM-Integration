package crm_handlers

import (
	crm_services "main/service/SalesActivities"

	"github.com/gofiber/fiber/v2"
)

type SalesActivitiesHandler struct {
	Service crm_services.Service_SalesActivities
}

func (h SalesActivitiesHandler) SalesActivities(c *fiber.Ctx) error {
	return nil
}

func (h SalesActivitiesHandler) SalesActivitiesCreate(c *fiber.Ctx) error {
	return nil
}
func (h SalesActivitiesHandler) SalesActivitiesById(c *fiber.Ctx) error {
	return nil
}
func (h SalesActivitiesHandler) SalesActivitiesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h SalesActivitiesHandler) SalesActivitiesDelete(c *fiber.Ctx) error {
	return nil
}
