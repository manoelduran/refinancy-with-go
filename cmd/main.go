package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/repositories"
	"github.com/manoelduran/refinancy-with-go/routes"
)

func main() {
    app := fiber.New()

    recipeRepository := repositories.NewRecipeRepository(db)
    recipeService := services(recipeRepository)
    recipeController := controllers.NewRecipeController(recipeService)

    routes.RecipeRoutes(app, recipeController)
    log.Fatal(app.Listen(":3000"))
}