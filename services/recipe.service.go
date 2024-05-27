package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

// ReceitaService is a struct that contains the repository

type RecipeService struct {

	repository repositories.RecipeRepository

}

// NewRecipeService is a function that returns a new instance of RecipeService

func NewRecipeService(repository repositories.RecipeRepository) *RecipeService {

	return &RecipeService{repository}

}

// GetRecipes is a function that returns all recipes

func (r *RecipeService) GetRecipes() ([]models.Recipe, error) {
	
	return r.repository.GetRecipes()

}	

// GetRecipe is a function that returns a recipe by id

func (r *RecipeService) GetRecipe(id uint) (models.Recipe, error) {
	
	return r.repository.GetRecipe(id)

}

// CreateRecipe is a function that creates a new recipe

func (r *RecipeService) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {

	return r.repository.CreateRecipe(recipe)

}

// UpdateRecipe is a function that updates a recipe

func (r *RecipeService) UpdateRecipe(id uint, recipe models.Recipe) (models.Recipe, error) {

	return r.repository.UpdateRecipe(id, recipe)

}

// DeleteRecipe is a function that deletes a recipe

func (r *RecipeService) DeleteRecipe(id uint) error {

	return r.repository.DeleteRecipe(id)

}