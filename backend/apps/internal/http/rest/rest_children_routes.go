package rest

import "github.com/gofiber/fiber/v2"

func (fiberApp *App) initChildrenRoutes() {
	v1 := fiberApp.apiChildren.Group("/v1/")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(newResp("great", nil))
	})
}
