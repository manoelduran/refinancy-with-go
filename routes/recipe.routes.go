package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
)

func RecipeRoutes(app *fiber.App, controller *controllers.RecipeController) {
	recipe := app.Group("/recipes")
	recipe.Get("/", controller.GetRecipes)
	recipe.Get("/:id", controller.GetRecipe)
	recipe.Post("/", controller.CreateRecipe)
	recipe.Put("/:id", controller.UpdateRecipe)
	recipe.Delete("/:id", controller.DeleteRecipe)
}