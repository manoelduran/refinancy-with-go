package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/models"
)

func ValidateExpense(v *validator.Validate) fiber.Handler {
    return ValidateMiddleware(v, &models.Expense{})
}