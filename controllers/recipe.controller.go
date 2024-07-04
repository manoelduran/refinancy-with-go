package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/services"
)
type RecipeInterface interface {
	GetRecipes(c *fiber.Ctx) error
	GetRecipe(c *fiber.Ctx) error
	CreateRecipe(c *fiber.Ctx) error
	UpdateRecipe(c *fiber.Ctx) error
	DeleteRecipe(c *fiber.Ctx) error
}
type RecipeController struct {
	service *services.RecipeService
}


func NewRecipeController(service *services.RecipeService) *RecipeController {
	return &RecipeController{service}
}

func (r *RecipeController) GetRecipes(c *fiber.Ctx) error {

	recipes, err := r.service.GetRecipes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(recipes)
}


func (r *RecipeController) GetRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }
	recipe, err := r.service.GetRecipe(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(recipe)
}

func (r *RecipeController) CreateRecipe(c *fiber.Ctx) error {
	recipe := new(models.Recipe)
	if err := c.BodyParser(recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	createdRecipe, err := r.service.CreateRecipe(*recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdRecipe)
}

func (r *RecipeController) UpdateRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	recipe := new(models.Recipe)
	if err := c.BodyParser(recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updatedRecipe, err := r.service.UpdateRecipe(uint(id), *recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(updatedRecipe)
}

func (r *RecipeController) DeleteRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	if err := r.service.DeleteRecipe(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

	return c.SendStatus(fiber.StatusNoContent)
}
