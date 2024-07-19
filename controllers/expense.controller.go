package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/services"
)
type ExpenseInterface interface {
	GetExpenses(c *fiber.Ctx) error
	GetExpense(c *fiber.Ctx) error
	CreateExpense(c *fiber.Ctx) error
	UpdateExpense(c *fiber.Ctx) error
	DeleteExpense(c *fiber.Ctx) error
}
type ExpenseController struct {
	*GenericController[models.Expense]
}


func NewExpenseController(service *services.ExpenseService) *ExpenseController {
	return &ExpenseController{
        GenericController: NewGenericController(service),
    }
}

func (e *ExpenseController) GetExpenses(c *fiber.Ctx) error {

	expenses, err := e.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(expenses)
}


func (e *ExpenseController) GetExpense(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }
	expense, err := e.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(expense)
}

func (r *ExpenseController) CreateExpense(c *fiber.Ctx) error {
	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	createdExpense, err := r.service.Create(*expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdExpense)
}

func (r *ExpenseController) UpdateExpense(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updatedExpense, err := r.service.Update(uint(id), *expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(updatedExpense)
}

func (r *ExpenseController) DeleteExpense(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	if err := r.service.Delete(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

	return c.SendStatus(fiber.StatusNoContent)
}
