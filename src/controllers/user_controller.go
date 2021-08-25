package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"teste.com/todos/src/dtos"
	"teste.com/todos/src/services"
)

func RegisterUser(c *fiber.Ctx) error {
	userDto := new(dtos.UserDto)
	if err := c.BodyParser(userDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}

	err := services.RegisterUser(*userDto)
	if err != nil {
		c.JSON(err)
		return errors.New(err.Error())
	}
	c.SendStatus(fiber.StatusCreated)
	return nil
}

func Login(c *fiber.Ctx) error {
	userDto := new(dtos.UserDto)
	if err := c.BodyParser(userDto); err != nil {
		c.JSON(fiber.ErrBadRequest)
		return errors.New(err.Error())
	}
	token, err := services.Login(*userDto)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		return err
	}
	c.JSON(fiber.Map{"token": token})
	return nil
}
