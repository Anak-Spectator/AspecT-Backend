package rest

import (
	"aspect_apps/internal/services/account/accountapp"

	"github.com/gofiber/fiber/v2"
)

func userLogin(svc *accountapp.AccountApplicationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		txt := svc.Test("bos")
		return c.Status(fiber.StatusOK).JSON(newResp(txt, txt))
	}
}

func userRegister(svc *accountapp.AccountApplicationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		txt := svc.Test("bos")
		return c.Status(fiber.StatusOK).JSON(newResp(txt, txt))
	}
}
