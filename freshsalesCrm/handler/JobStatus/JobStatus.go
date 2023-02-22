package crm_handlers

import (
	crm_services "main/service/JobStatus"

	"github.com/gofiber/fiber/v2"
)

type JobStatusHandler struct {
	Service crm_services.Service_JobStatus
}

func (h JobStatusHandler) JobStatus(c *fiber.Ctx) error {
	return nil
}

func (h JobStatusHandler) JobStatusCreate(c *fiber.Ctx) error {
	return nil
}
func (h JobStatusHandler) JobStatusById(c *fiber.Ctx) error {
	return nil
}
func (h JobStatusHandler) JobStatusUpdate(c *fiber.Ctx) error {
	return nil
}
func (h JobStatusHandler) JobStatusDelete(c *fiber.Ctx) error {
	return nil
}
