package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "used@example.org",
		Password: "password",
	}
}
