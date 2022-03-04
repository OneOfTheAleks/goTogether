package store_test

import (
	"GoTogether/model"
	"GoTogether/store"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s,teardown := store.TestStore(t,dataBaseUrl)
	defer teardown("users")
     var ctx context.Context
     ctx= context.TODO()
    // ctx, _ =context.WithTimeout(ctx,3*time.Second)

	u,err := s.User().Create(
		&model.User{
		Email: "test@mail.ru",
	     },
	     ctx,
	)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s,teardown := store.TestStore(t,dataBaseUrl)
	defer teardown("users")
	var ctx context.Context
	ctx= context.TODO()
	email:= "test@mail.ru"
	u,err := s.User().FindByEmail(email,ctx)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}