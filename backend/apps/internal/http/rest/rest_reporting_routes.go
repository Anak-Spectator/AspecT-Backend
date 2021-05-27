package rest

import "github.com/gofiber/fiber/v2"

/*
? ===== TODO ===== ?
[ ] Setup handling using query paramaters
[ ] Get reporting data for that day request
[ ] Get reporitng for a week
[ ] Get reporitng for a month
? === END TODO === ?
*/

func (fiberApp *App) initUserReportingRoutes() {
	v1 := fiberApp.apiUser.Group("reports")

	v1.Use(parseUserAuthorizationHeader())
	v1.Get("/", func(c *fiber.Ctx) error {
		return nil
	}) // Get all info * Read TODO *
}
