package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type ExpenseService struct {
    *GenericService[models.Expense]
}

func NewExpenseService(repository *repositories.ExpenseRepository) *ExpenseService {
    return &ExpenseService{
        GenericService: NewGenericService(repository),
    }
}

func (e *ExpenseService) GetExpenses() ([]models.Expense, error) {

	return e.repository.GetAll()

}

func (e *ExpenseService) GetExpense(id uint) (models.Expense, error) {

	return e.repository.GetByID(id)

}

func (e *ExpenseService) CreateExpense(expense models.Expense) (models.Expense, error) {

	return e.repository.Create(expense)

}


func (e *ExpenseService) UpdateExpense(id uint, expense models.Expense) (models.Expense, error) {

	return e.repository.Update(id, expense)

}

func (e *ExpenseService) DeleteExpense(id uint) error {

	return e.repository.Delete(id)

}