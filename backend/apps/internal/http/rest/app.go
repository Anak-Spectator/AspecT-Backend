package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

/*
TODO

[ ] Define All Each Service
	[ ] Parrent
		[ ]

	[ ] Children
		[ ]

	[ ] Profanity
		[ ]

	[ ] Auth
		[ ]

	[ ] Reporting
		[ ]
*/

type Service struct {
	Test string
}

type App struct {
	app         *fiber.App
	apiParent   fiber.Router
	apiChildren fiber.Router
	service     *Service
}

func NewApp(service *Service) *App {
	app := fiber.New(fiber.Config{})
	return &App{
		app:         app,
		apiParent:   app.Group("/parrent"),
		apiChildren: app.Group("/children"),
		service: &Service{
			Test: service.Test,
		},
	}
}

func (fiberApp *App) StartApp(port string) {

	// Define the app
	app := fiberApp.app

	// Recovery when get PANIC
	app.Use(recover.New(recover.ConfigDefault))

	// Setup cors
	app.Use(cors.New())

	// Handle logger
	app.Use(logger.New())

	// Start Each Service Routes
	fiberApp.initChildrenRoutes()
	// End Each Service Routes

	// Handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(newErrorResp(fmt.Errorf("your requested URL %s can't be found, please recheck your requested URL", c.Path())))
	})

	// Start app
	panic(app.Listen(port))
}
