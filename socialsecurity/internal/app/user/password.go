package user

import "golang.org/x/crypto/bcrypt"

type Password string

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p Password) String() string {
	return string(p)
}
