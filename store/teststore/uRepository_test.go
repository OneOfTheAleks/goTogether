package teststore_test

import (
	"GoTogether/model"
	"GoTogether/store"
	"GoTogether/store/teststore"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u,ctx))
	assert.NotNil(t, u.Id)
}
/*
func TestUserRepository_Find(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1,ctx)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}*/

func TestUserRepository_FindByEmail(t *testing.T) {
	ctx := context.Background()
	s := teststore.New()
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email,ctx)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	assert.NoError(t,err)
	assert.NotNil(t,u1)

/*	s.User().Create(u1,ctx)
	u2, err := s.User().FindByEmail(u1.Email,ctx)
	assert.NoError(t, err)
	assert.NotNil(t, u2)*/
}
