package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/middlewares"
)

func UserRoutes(app *fiber.App, controller *controllers.UserController, validator *validator.Validate) {
    user := app.Group("/users")
    user.Get("/", controller.GetUsers)
    user.Get("/:id", controller.GetUser)
    user.Post("/", middlewares.ValidateUser(validator), controller.CreateUser)
    user.Put("/:id", middlewares.ValidateUser(validator), controller.UpdateUser)
    user.Delete("/:id", controller.DeleteUser)
}