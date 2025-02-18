package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
)

type AuthClaims struct {
	UserID ulid.ULID `json:"userId"`
	jwt.RegisteredClaims
}
