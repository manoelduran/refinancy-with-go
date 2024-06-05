package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IError struct {
    Field string `json:"field"`
    Tag   string `json:"tag"`
    Value string `json:"value"`
}

func ValidateMiddleware(v *validator.Validate, s interface{}) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var errors []*IError
        if err := c.BodyParser(s); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
        }

        err := v.Struct(s)
        if err != nil {
            for _, err := range err.(validator.ValidationErrors) {
                var el IError
                el.Field = err.Field()
                el.Tag = err.Tag()
                el.Value = err.Param()
                errors = append(errors, &el)
            }
            return c.Status(fiber.StatusBadRequest).JSON(errors)
        }
        return c.Next()
    }
}
