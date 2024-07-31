package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/services"
)
type CashflowInterface interface {
	GetCashflow(c *fiber.Ctx) error
	CreateCashflow(c *fiber.Ctx) error
	DeleteCashflow(c *fiber.Ctx) error
}
type CashflowController struct {
	*GenericController[models.Cashflow]
}


func NewCashflowController(service *services.CashflowService) *CashflowController {
	return &CashflowController{
        GenericController: NewGenericController(service),
    }
}

// func (r *CashflowController) GetCashflows(c *fiber.Ctx) error {

// 	cashflows, err := r.service.GetAll()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	return c.JSON(cashflows)
// }


func (r *CashflowController) GetCashflow(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }
	cashflow, err := r.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(cashflow)
}

func (r *CashflowController) CreateCashflow(c *fiber.Ctx) error {
	cashflow := new(models.Cashflow)
	if err := c.BodyParser(cashflow); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	createdCashflow, err := r.service.Create(*cashflow)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdCashflow)
}

// func (r *CashflowController) UpdateCashflow(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
//     }

// 	Cashflow := new(models.Cashflow)
// 	if err := c.BodyParser(Cashflow); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	updatedCashflow, err := r.service.Update(uint(id), *Cashflow)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	return c.JSON(updatedCashflow)
// }

// func (r *CashflowController) DeleteCashflow(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
//     }

// 	if err := r.service.Delete(uint(id)); err != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//     }

// 	return c.SendStatus(fiber.StatusNoContent)
// }
