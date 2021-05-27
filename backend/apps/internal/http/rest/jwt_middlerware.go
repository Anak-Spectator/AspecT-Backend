package rest

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	useridentityKey     = "user_identity"
	childrenidentityKey = "children_identity"
)

func parseUserAuthorizationHeader() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := strings.Split(c.Get("Authorization", ""), " ")
		if len(authHeader) != 2 {
			return c.Status(fiber.StatusBadRequest).JSON(newErrorResp(errors.New("invalid access token")))
		}

		if authHeader[0] != "Bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(newErrorResp(errors.New("invalid access token")))
		}

		c.Locals(useridentityKey, authHeader[1])
		return c.Next()
	}
}

func parseChildrenAuthorizationHeader() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := strings.Split(c.Get("Authorization", ""), " ")
		if len(authHeader) != 2 {
			return c.Status(fiber.StatusBadRequest).JSON(newErrorResp(errors.New("invalid access token")))
		}

		if authHeader[0] != "Bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(newErrorResp(errors.New("invalid access token")))
		}

		c.Locals(childrenidentityKey, authHeader[1])
		return c.Next()
	}
}
