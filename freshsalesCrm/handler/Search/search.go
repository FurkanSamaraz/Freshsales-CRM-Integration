package crm_handlers

import (
	crm_services "main/service/Search"

	"github.com/gofiber/fiber/v2"
)

type SearchHandler struct {
	Service crm_services.Service_Search
}

func (h SearchHandler) Search(c *fiber.Ctx) error {
	return nil
}
