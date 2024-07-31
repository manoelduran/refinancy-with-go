package repositories

import (
	"database/sql"

	"github.com/manoelduran/refinancy-with-go/models"
)
type ExpenseRepository struct {
	*GenericRepository[models.Expense]
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	fields := []string{"Title", "From_By", "Description", "Value", "Received_By"}
	return &ExpenseRepository{
		GenericRepository: NewGenericRepository[models.Expense](db, "expenses", fields),
	}
}

func (e *ExpenseRepository) GetExpenses() ([]models.Expense, error) {
	return e.GetAll()
}

func (e *ExpenseRepository) GetExpense(id uint) (models.Expense, error) {
	return e.GetByID(id)
}

func (e *ExpenseRepository) CreateExpense(expense models.Expense) (models.Expense, error) {
	return e.Create(expense)
}

func (e *ExpenseRepository) UpdateExpense(id uint, expense models.Expense) (models.Expense, error) {
	return e.Update(id, expense)
}
func (e *ExpenseRepository) DeleteExpense(id uint) error {
	return e.Delete(id)
}