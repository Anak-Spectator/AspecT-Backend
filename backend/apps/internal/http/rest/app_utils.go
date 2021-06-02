package rest

import "github.com/gofiber/fiber/v2"

type errorResp struct {
	Error string `json:"error"`
}

func newErrorResp(err error) *errorResp {
	return &errorResp{
		Error: err.Error(),
	}
}

type resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func newResp(message string, data interface{}) *resp {
	return &resp{
		Message: message,
		Data:    data,
	}
}

func (app *App) apiOk() {
	app.apiV1.Get("/start", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(newResp("everything's is okay bro!!", fiber.Map{
			"version": "1.0.0",
		}))
	})
}
