package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}

func TestUserWhenNameIsRequired(t *testing.T) {
	p, err := NewUser("", "j@j.com", "123456")
	assert.Nil(t, p)
	assert.Equal(t, len(err), 1)
	fmt.Println(err[0].Error())
	assert.Equal(t, err[0].Error(), ErrNameIsMandatory.Error())
}

func TestUserWhenEmailIsRequired(t *testing.T) {
	p, err := NewUser("John Doe", "", "123456")
	assert.Nil(t, p)
	assert.Equal(t, len(err), 1)
	fmt.Println(err[0].Error())
	assert.Equal(t, err[0].Error(), ErrEmailIsMandatory.Error())
}
