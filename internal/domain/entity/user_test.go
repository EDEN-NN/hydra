package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("zalison", "gadelhoscaralho", "gadders@gmail.com", "Vitorio")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUser_IsValid(t *testing.T) {
	user, err := CreateUser("zalison", "gadelhoscaralho", "gadders@gmail.com", "Vitorio")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	_, err = user.ChangeName("za")

	assert.NotNil(t, err)
}
