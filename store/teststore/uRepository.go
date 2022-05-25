package teststore

import (
	"GoTogether/model"
	"GoTogether/store"
	"context"
	"errors"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User,ctx context.Context) error {
	if _,err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

func (r *UserRepository) Find(id int,ctx context.Context) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil,store.ErrRecordNotFound
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string,ctx context.Context) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("email not found")
}