package rest

import "github.com/gofiber/fiber/v2"

func (fiberApp *App) initUserProfanityRoutes() {
	v1 := fiberApp.apiUser.Group("/profanity")
	v1.Post("/send", parseChildrenAuthorizationHeader(), func(c *fiber.Ctx) error {
		return nil
	}) // children send a text for checking
}

func (fiberApp *App) initUserChildrenProfanityRoutes() {
	v1 := fiberApp.apiUser.Group("/profanity")

	v1.Use(parseUserAuthorizationHeader())
	v1.Get("/:children_id", func(c *fiber.Ctx) error {
		return nil
	}) // get full profanity list using children id

	v1.Post("/:children_id/:profanity_id", func(c *fiber.Ctx) error {
		return nil
	}) // Report profanity to admin
}
