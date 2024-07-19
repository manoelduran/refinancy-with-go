package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/manoelduran/refinancy-with-go/services"
)

type GenericController[T any] struct {
    service services.Service[T]
}

func NewGenericController[T any](service services.Service[T]) *GenericController[T] {
    return &GenericController[T]{service}
}

func (c *GenericController[T]) GetAll(ctx *fiber.Ctx) error {
    items, err := c.service.GetAll()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }
    return ctx.JSON(items)
}

func (c *GenericController[T]) GetByID(ctx *fiber.Ctx) error {
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID"})
    }
    item, err := c.service.GetByID(uint(id))
    if err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
    }
    return ctx.JSON(item)
}

func (c *GenericController[T]) Create(ctx *fiber.Ctx) error {
    item := new(T)
    if err := ctx.BodyParser(item); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
    }
    log.Printf("Item recebido: %+v\n", item)
    createdItem, err := c.service.Create(*item)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }
    return ctx.Status(fiber.StatusCreated).JSON(createdItem)
}

func (c *GenericController[T]) Update(ctx *fiber.Ctx) error {
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID"})
    }
    item := new(T)
    if err := ctx.BodyParser(item); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
    }
    updatedItem, err := c.service.Update(uint(id), *item)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }
    return ctx.JSON(updatedItem)
}

func (c *GenericController[T]) Delete(ctx *fiber.Ctx) error {
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid ID"})
    }
    if err := c.service.Delete(uint(id)); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }
    return ctx.SendStatus(fiber.StatusNoContent)
}
