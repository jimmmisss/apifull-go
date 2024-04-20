package entity

import (
	"errors"

	"github.com.br/jimmmisss/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNameIsMandatory     = errors.New("name is mandatory")
	ErrEmailIsMandatory    = errors.New("email is mandatory")
	ErrPasswordIsMandatory = errors.New("password is mandatory")
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := GeneratePasswordBCrypt(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	return user, nil
}

func GeneratePasswordBCrypt(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Validate() []error {
	var allErrors []error

	if u.Name == "" {
		allErrors = append(allErrors, ErrNameIsMandatory)
	}
	if u.Email == "" {
		allErrors = append(allErrors, ErrEmailIsMandatory)
	}
	if u.Password == "" {
		allErrors = append(allErrors, ErrPasswordIsMandatory)
	}

	if len(allErrors) > 0 {
		return allErrors
	}

	return nil
}
