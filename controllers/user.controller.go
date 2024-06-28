package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/services"
)
type UserInterface interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
type UserController struct {
	service *services.UserService
}


func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

func (u *UserController) GetUsers(c *fiber.Ctx) error {

	users, err := u.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(users)
}


func (u *UserController) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }
	user, err := u.service.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(user)
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	createdUser, err := u.service.CreateUser(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updatedUser, err := u.service.UpdateUser(uint(id), *user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(updatedUser)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
    }

	if err := u.service.DeleteUser(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

	return c.SendStatus(fiber.StatusNoContent)
}
