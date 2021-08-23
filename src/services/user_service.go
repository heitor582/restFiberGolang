package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"teste.com/todos/src/configuration"
	"teste.com/todos/src/dtos"
	"teste.com/todos/src/entities"
	"teste.com/todos/src/functions"
)

func Login(userDto dtos.UserDto) (string, error) {
	db := configuration.DBConn
	var user entities.UserModel
	db.Find(&user, "username = ?", userDto.Username)
	if user.ID == uuid.Nil {
		return "", errors.New("user not found")
	}
	if !functions.CheckPasswordHash(userDto.Password, user.Password) {
		return "", errors.New("Passwords dont match")
	}
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
func RegisterUser(userDto dtos.UserDto) error {
	db := configuration.DBConn
	hashedPassword, err := functions.HashPassword(userDto.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	var user entities.UserModel = entities.UserModel{
		Username:  userDto.Username,
		Password:  hashedPassword,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = db.Create(&user).Error; err != nil {
		return errors.New(err.Error())
	}
	return nil
}
