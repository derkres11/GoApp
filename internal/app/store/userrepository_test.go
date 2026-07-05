package store_test

import (
	"TestApi/internal/app/model"
	"TestApi/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserReposistory_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserReposistory_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, DatabaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}
