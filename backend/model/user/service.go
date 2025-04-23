package user

import (
	"backend/database"
	"backend/database/gen"
	"context"

	"github.com/oklog/ulid/v2"
)

type Service interface {
	CreateUser(username, password string) (gen.User, error)
	FetchUsers() (*[]gen.User, error)
	FetchUser(username string) (*gen.User, error)
	UpdateUser(user *gen.User) (*gen.User, error)
	RemoveUser(ID ulid.ULID) error
	ExistUser() (bool, error)
}

type service struct {
	db *database.DB
}

func NewService(db *database.DB) Service {
	return &service{
		db,
	}
}

func (u *service) CreateUser(username, password string) (gen.User, error) {
	return u.db.CreateUser(context.TODO(), gen.CreateUserParams{
		Username: username,
		Password: password,
	})
}

func (u *service) FetchUsers() (*[]gen.User, error) {
	return nil, nil
}

func (u *service) FetchUser(username string) (*gen.User, error) {
	return nil, nil
}

func (u *service) UpdateUser(book *gen.User) (*gen.User, error) {
	return nil, nil
}

func (u *service) RemoveUser(ID ulid.ULID) error {
	return nil
}

func (u *service) ExistUser() (bool, error) {
	return u.db.ExistUser(context.TODO())
}
