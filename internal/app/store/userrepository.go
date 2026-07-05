package store

import "TestApi/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	r.store.db.QueryRow("INSERT INTO user(email, encrypted_password) VALUES ($1, $2) RETURNING id")

	return nil, nil

}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
