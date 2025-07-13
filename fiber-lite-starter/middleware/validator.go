package middleware

import (
	"fiber-lite-starter/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ValidatorMiddleware(v *validator.Validator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("validator", func(i interface{}) error {
			return v.Validate(i)
		})
		return c.Next()
	}
}
