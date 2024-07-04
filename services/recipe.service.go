package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type RecipeService struct {
    *GenericService[models.Recipe]
}

func NewRecipeService(repository *repositories.RecipeRepository) *RecipeService {
    return &RecipeService{
        GenericService: NewGenericService(repository),
    }
}

func (r *RecipeService) GetRecipes() ([]models.Recipe, error) {

	return r.repository.GetAll()

}

func (r *RecipeService) GetRecipe(id uint) (models.Recipe, error) {

	return r.repository.GetByID(id)

}

func (r *RecipeService) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {

	return r.repository.Create(recipe)

}


func (r *RecipeService) UpdateRecipe(id uint, recipe models.Recipe) (models.Recipe, error) {

	return r.repository.Update(id, recipe)

}

func (r *RecipeService) DeleteRecipe(id uint) error {

	return r.repository.Delete(id)

}