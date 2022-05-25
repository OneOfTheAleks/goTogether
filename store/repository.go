package store

import (
  "GoTogether/model"
  "context"
)

type UserRepository interface {
  Create(user *model.User,ctx context.Context) error
  FindByEmail(string,context.Context) (*model.User,error)
}
