package user

import (
	"backend/database/gen"

	"github.com/oklog/ulid/v2"
)

type UserDTO struct{
	Username string `json:"username"`
	Password string `json:"password"`
}


func NewUser(id ulid.ULID, username, password string) *gen.User{
	return &gen.User{
			ID: id,
			Username: username,
			Password: password,
		}
}