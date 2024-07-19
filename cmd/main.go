package main

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/manoelduran/refinancy-with-go/controllers"
	"github.com/manoelduran/refinancy-with-go/database"
	"github.com/manoelduran/refinancy-with-go/repositories"
	"github.com/manoelduran/refinancy-with-go/routes"
	"github.com/manoelduran/refinancy-with-go/services"
)

func main() {
    app := fiber.New(fiber.Config{
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError

            var e *fiber.Error
            if errors.As(err, &e) {
                code = e.Code
            }

            ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

            return ctx.Status(code).SendString(err.Error())
        },
    })
    validator := validator.New()
    app.Use(recover.New())

    database.InitDatabase()
    defer database.CloseDatabase()

    recipeRepository := repositories.NewRecipeRepository(database.DB)
    recipeService := services.NewRecipeService(recipeRepository)
    recipeController := controllers.NewRecipeController(recipeService)
    routes.RecipeRoutes(app, recipeController, validator)

	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
    routes.UserRoutes(app, userController, validator)
    log.Fatal(app.Listen(":3000"))
}
