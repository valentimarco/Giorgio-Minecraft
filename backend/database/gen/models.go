// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package gen

import (
	ulid "github.com/oklog/ulid/v2"
)

type User struct {
	ID       ulid.ULID
	Username string
	Password string
}
