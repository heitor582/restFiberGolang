package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"teste.com/todos/src/dtos"
	"teste.com/todos/src/services"
)

func GetAllTodos(c *fiber.Ctx) error {
	c.JSON(services.GetAllTodos())
	return nil
}

func GetTodo(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	todo, err := services.GetTodo(id)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}

func NewTodo(c *fiber.Ctx) error {
	todoDto := new(dtos.TodoDto)
	if err := c.BodyParser(todoDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}

	todo, err := services.NewTodo(*todoDto)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}

func DeleteTodo(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	todo, err := services.DeleteTodo(id)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}
