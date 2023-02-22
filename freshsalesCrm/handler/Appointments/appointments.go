package crm_handlers

import (
	crm_services "main/service/Appointments"

	"github.com/gofiber/fiber/v2"
)

type AppointmentsHandler struct {
	Service crm_services.Service_Appointments
}

func (h AppointmentsHandler) Appointments(c *fiber.Ctx) error {
	return nil
}

func (h AppointmentsHandler) AppointmentsCreate(c *fiber.Ctx) error {
	return nil
}
func (h AppointmentsHandler) AppointmentsById(c *fiber.Ctx) error {
	return nil
}
func (h AppointmentsHandler) AppointmentsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h AppointmentsHandler) AppointmentsDelete(c *fiber.Ctx) error {
	return nil
}
