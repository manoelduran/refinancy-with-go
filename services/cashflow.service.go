package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type CashflowService struct {
    *GenericService[models.Cashflow]
}

func NewCashflowService(repository *repositories.CashflowRepository) *CashflowService {
    return &CashflowService{
        GenericService: NewGenericService(repository),
    }
}

func (e *CashflowService) GetCashflow(id uint) (models.Cashflow, error) {

	return e.repository.GetByID(id)

}

func (e *CashflowService) CreateCashflow(cashflow models.Cashflow) (models.Cashflow, error) {

	return e.repository.Create(cashflow)

}