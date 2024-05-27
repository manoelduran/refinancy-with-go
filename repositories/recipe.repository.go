package repositories

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"gorm.io/gorm"
)

// RecipeRepository is a struct that contains the database

type RecipeRepository struct {

	db *gorm.DB

}

// NewRecipeRepository is a function that returns a new instance of RecipeRepository

func NewRecipeRepository(db *gorm.DB) *RecipeRepository {

	return &RecipeRepository{db}

}

// GetRecipes is a function that returns all recipes

func (r *RecipeRepository) GetRecipes() ([]models.Recipe, error) {

	var recipes []models.Recipe
	result := r.db.Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}
	return recipes, nil

}

// GetRecipe is a function that returns a recipe by id

func (r *RecipeRepository) GetRecipe(id uint) (models.Recipe, error) {
	
	var recipe models.Recipe
	result := r.db.First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil

}

// CreateRecipe is a function that creates a new recipe

func (r *RecipeRepository) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {

	result := r.db.Create(&recipe)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil

}

// UpdateRecipe is a function that updates a recipe

func (r *RecipeRepository) UpdateRecipe(id uint, recipe models.Recipe) (models.Recipe, error) {

	result := r.db.Model(&models.Recipe{}).Where("id = ?", id).Updates(recipe)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil

}

// DeleteRecipe is a function that deletes a recipe

func (r *RecipeRepository) DeleteRecipe(id uint) error {

	result := r.db.Delete(&models.Recipe{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}