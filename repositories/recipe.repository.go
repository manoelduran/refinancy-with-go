package repositories

import (
	"database/sql"

	"github.com/manoelduran/refinancy-with-go/models"
)

type RecipeRepository struct {
	*GenericRepository[models.Recipe]
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	fields := []string{"Id", "Title", "FromBy", "Description", "Value", "ReceivedAt", "ReceivedBy", "CreatedAt", "UpdatedAt"}
	return &RecipeRepository{
		GenericRepository: NewGenericRepository[models.Recipe](db, "recipes", fields),
	}
}

func (r *RecipeRepository) GetRecipes() ([]models.Recipe, error) {
	return r.GetAll()
}

func (r *RecipeRepository) GetRecipe(id uint) (models.Recipe, error) {
	return r.GetByID(id)
}

func (r *RecipeRepository) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
	return r.Create(recipe)
}

func (r *RecipeRepository) UpdateRecipe(id uint, recipe models.Recipe) (models.Recipe, error) {
	return r.Update(id, recipe)
}

func (r *RecipeRepository) DeleteRecipe(id uint) error {
	return r.Delete(id)
}
