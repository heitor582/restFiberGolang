package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"teste.com/todos/src/dtos"
	"teste.com/todos/src/services"
)

func GetAllTodos(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uuid.MustParse(claims["id"].(string))
	c.JSON(services.GetAllTodos(userId))
	return nil
}

func GetTodo(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uuid.MustParse(claims["id"].(string))
	todo, err := services.GetTodo(id, userId)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}

func NewTodo(c *fiber.Ctx) error {
	todoDto := new(dtos.TodoDto)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uuid.MustParse(claims["id"].(string))
	if err := c.BodyParser(todoDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}

	todo, err := services.NewTodo(*todoDto, userId)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.JSON(todo)
	return nil
}

func UpdateTodo(c *fiber.Ctx) error {
	todoDto := new(dtos.TodoDto)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uuid.MustParse(claims["id"].(string))
	if err := c.BodyParser(todoDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}
	todo, err := services.UpdateTodo(*todoDto, userId)
	if err != nil {
		c.JSON(err)
		return err
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
