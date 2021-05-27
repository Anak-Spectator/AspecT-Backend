package rest

import (
	"errors"
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
		[ ] Edit info
		[ ] Add new child
			[ ] add image to child

	[ ] Children
		[ ] Check if id exist
		[ ] Send Text within the children id
		[ ]

	[ ] Profanity
		[ ] Send to ML service using Goroutine

	[ ] Auth
		[ ] Sign up
		[ ] Sign in

	[ ] Reporting
		[ ] Reporting Each Day
		[ ] Reporting Each Week
		[ ] Reporting Each Month
*/

type Service struct {
	AccSvc *AccountMainService
}

type App struct {
	app *fiber.App

	apiV1 fiber.Router // Independent Router

	// Grouping API
	apiUser  fiber.Router
	apiAdmin fiber.Router
	// End Of Grouping API

	// Interface for each service
	service *Service
}

func NewApp(service *Service) *App {
	app := fiber.New(fiber.Config{})
	apiV1 := app.Group("/api/v1")
	return &App{
		app:      app,
		apiV1:    apiV1,
		apiUser:  apiV1.Group("/user"),  // for user purpose
		apiAdmin: apiV1.Group("/admin"), // for admin purpose "develope soon"
		service: &Service{
			AccSvc: service.AccSvc,
		},
	}
}

func (fiberApp *App) StartApp(key, port string) {

	// Define the app
	app := fiberApp.app

	// Header Api KEY
	app.Use(func(c *fiber.Ctx) error {
		apiKey := c.Get("api_key", "")
		if apiKey != key {
			return c.Status(fiber.StatusForbidden).JSON(newErrorResp(errors.New("forbidden unknow client")))
		}
		return c.Next()
	})
	// Recovery when get PANIC
	app.Use(recover.New(recover.ConfigDefault))

	// Setup cors
	app.Use(cors.New())

	// Handle logger
	app.Use(logger.New())

	// Start Each Service Routes
	fiberApp.apiOk()                    // Start Endpoint
	fiberApp.initPublicChildrenRoutes() // indpendent public children route
	fiberApp.InitAuthRoutes()           // for auth
	// Private API Start Here
	fiberApp.initUserChildrenRoutes() // for user children
	// Private API End

	// End Each Service Routes

	// Handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(newErrorResp(fmt.Errorf("your requested URL %s can't be found, please recheck your requested URL", c.Path())))
	})

	// Start app
	panic(app.Listen(port))
}
