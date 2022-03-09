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

    ctx:= context.TODO()

	u,err := s.User().Create(
		model.TestUser(t),
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