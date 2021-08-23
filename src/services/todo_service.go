package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"teste.com/todos/src/configuration"
	"teste.com/todos/src/dtos"
	"teste.com/todos/src/entities"
)

func GetAllTodos() []entities.TodoModel {
	db := configuration.DBConn
	var todos []entities.TodoModel
	db.Find(&todos)
	return todos
}

func GetTodo(id uuid.UUID) (entities.TodoModel, error) {
	db := configuration.DBConn
	var todo entities.TodoModel
	db.Find(&todo, id)
	if todo.ID == uuid.Nil {
		return todo, errors.New("Todo not found")
	}
	return todo, nil
}

func NewTodo(todoDto dtos.TodoDto) (entities.TodoModel, error) {
	db := configuration.DBConn
	var todo entities.TodoModel = entities.TodoModel{
		Message:   todoDto.Message,
		Completed: todoDto.Completed,
		UserID:    todoDto.UserID,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := db.Create(&todo).Error
	if err != nil {
		return entities.TodoModel{}, errors.New("Todo not found")
	}
	return todo, nil
}

func DeleteTodo(id uuid.UUID) (entities.TodoModel, error) {
	db := configuration.DBConn
	var todo entities.TodoModel
	db.First(&todo, id)
	if todo.ID == uuid.Nil {
		return todo, errors.New("Todo not found")
	}
	db.Delete(&todo)
	return todo, nil
}
