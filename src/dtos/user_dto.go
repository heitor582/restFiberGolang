package dtos

type UserDto struct {
	Username string `json:"username" validate:"required,string"`
	Password string `json:"password" validate:"required,string"`
}
