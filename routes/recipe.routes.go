package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/middlewares"
)

func RecipeRoutes(app *fiber.App, controller *controllers.RecipeController, validator *validator.Validate) {
    recipe := app.Group("/recipes")
    recipe.Get("/", controller.GetRecipes)
    recipe.Get("/:id", controller.GetRecipe)
    recipe.Post("/", middlewares.ValidateRecipe(validator), controller.CreateRecipe)
    recipe.Put("/:id", middlewares.ValidateRecipe(validator), controller.UpdateRecipe)
    recipe.Delete("/:id", controller.DeleteRecipe)
}

