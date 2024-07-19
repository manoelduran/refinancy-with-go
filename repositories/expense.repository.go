package repositories

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"gorm.io/gorm"
)
type ExpenseRepository struct {
	GenericRepository[models.Expense]
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {

	return &ExpenseRepository{
		GenericRepository: GenericRepository[models.Expense]{db: db},
	}

}

func (e *ExpenseRepository) GetExpenses() ([]models.Expense, error) {

	var expenses []models.Expense
	result := e.db.Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil

}

func (e *ExpenseRepository) GetExpense(id uint) (models.Expense, error) {

	var expense models.Expense
	result := e.db.First(&expense, id)
	if result.Error != nil {
		return models.Expense{}, result.Error
	}
	return expense, nil

}

func (e *ExpenseRepository) CreateExpense(expense models.Expense) (models.Expense, error) {

	result := e.db.Create(&expense)
	if result.Error != nil {
		return models.Expense{}, result.Error
	}
	return expense, nil

}

func (e *ExpenseRepository) UpdateExpense(id uint, expense models.Expense) (models.Expense, error) {

	result := e.db.Model(&models.Expense{}).Where("id = ?", id).Updates(expense)
	if result.Error != nil {
		return models.Expense{}, result.Error
	}
	return expense, nil

}
func (e *ExpenseRepository) DeleteExpense(id uint) error {

	result := e.db.Delete(&models.Expense{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}