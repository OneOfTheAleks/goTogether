package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email    : "test@mail.ru",
		OpenPas  : "1234567890",
	}
}
