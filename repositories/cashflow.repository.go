package repositories

import (
	"database/sql"

	"github.com/manoelduran/refinancy-with-go/models"
)
type CashflowRepository struct {
	*GenericRepository[models.Cashflow]
}
func NewCashflowRepository(db *sql.DB) *CashflowRepository {
	fields := []string{"Title", "User_Id", "Description", "Earnings","Costs","Total"}
	return &CashflowRepository{
		GenericRepository: NewGenericRepository[models.Cashflow](db, "cashflow", fields),
	}
}

func (e *CashflowRepository) GetCashflow(id uint) (models.Cashflow, error) {
	return e.GetByID(id)
}

func (e *CashflowRepository) CreateCashflow(cashflow models.Cashflow) (models.Cashflow, error) {
	return e.Create(cashflow)
}