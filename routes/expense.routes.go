package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/middlewares"
)

func ExpenseRoutes(app *fiber.App, controller *controllers.ExpenseController, validator *validator.Validate) {
    expense := app.Group("/expenses")
    expense.Get("/", controller.GetAll)
    expense.Get("/:id", controller.GetByID)
    expense.Post("/", middlewares.ValidateExpense(validator), controller.Create)
    expense.Put("/:id", middlewares.ValidateExpense(validator), controller.Update)
    expense.Delete("/:id", controller.Delete)
}

