package dtos

type TodoDto struct {
	Completed bool   `json:"completed" validate:"required,bool"`
	Message   string `json:"message" validate:"required,string"`
}
