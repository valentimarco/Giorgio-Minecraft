package util

import (
	"backend/model"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateHash(password string) (hash string, err error) {
	pwbyte := []byte(password)
	if len(pwbyte) > 72 {
		err = errors.New("password too long")
		return
	}
	hashbyte, err := bcrypt.GenerateFromPassword(pwbyte, bcrypt.DefaultCost)
	hash = string(hashbyte)
	return
}

func GenerateToken(id ulid.ULID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, model.AuthClaims{
		UserID: id,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
