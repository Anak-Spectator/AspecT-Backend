package rest

import "github.com/gofiber/fiber/v2"

func ChildrenGreat() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(newResp("great", nil))
	}

}
