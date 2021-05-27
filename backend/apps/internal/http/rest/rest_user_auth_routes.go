package rest

import "github.com/gofiber/fiber/v2"

func (fiberApp *App) InitAuthRoutes() {
	v1 := fiberApp.apiUser.Group("/auth")
	v1.Post("/login", userLogin(fiberApp.service.AccSvc.accAppSvc))
	v1.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"hellow": "register"})
	})
}
