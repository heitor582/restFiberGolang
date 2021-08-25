package dtos

type TodoDto struct {
	Completed bool   `json:"completed"`
	Message   string `json:"message"`
}
