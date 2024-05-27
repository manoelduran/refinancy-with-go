package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/database"
	"github.com/manoelduran/refinancy-with-go/repositories"
	"github.com/manoelduran/refinancy-with-go/routes"
	"github.com/manoelduran/refinancy-with-go/services"
)

func main() {
    app := fiber.New()
    database.InitDatabase()
    defer database.CloseDatabase()
    recipeRepository := repositories.NewRecipeRepository(database.DB)
    recipeService := services.NewRecipeService(*recipeRepository)
    recipeController := controllers.NewRecipeController(*recipeService)

    routes.RecipeRoutes(app, recipeController)
    log.Fatal(app.Listen(":3000"))
}