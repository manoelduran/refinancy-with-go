package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type RecipeService struct {

	repository *repositories.RecipeRepository

}

func NewRecipeService(repository *repositories.RecipeRepository) *RecipeService {

	return &RecipeService{repository}

}

func (r *RecipeService) GetRecipes() ([]models.Recipe, error) {

	return r.repository.GetRecipes()

}

func (r *RecipeService) GetRecipe(id uint) (models.Recipe, error) {

	return r.repository.GetRecipe(id)

}

func (r *RecipeService) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {

	return r.repository.CreateRecipe(recipe)

}


func (r *RecipeService) UpdateRecipe(id uint, recipe models.Recipe) (models.Recipe, error) {

	return r.repository.UpdateRecipe(id, recipe)

}

func (r *RecipeService) DeleteRecipe(id uint) error {

	return r.repository.DeleteRecipe(id)

}