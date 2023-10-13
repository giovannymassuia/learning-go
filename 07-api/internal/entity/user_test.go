package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@email.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "john.doe@email.com", "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
	assert.NotEqual(t, "123456", user.Password)
}
