package sqlStore_test

import (
	"GoTogether/model"
	"GoTogether/store"
	"GoTogether/store/sqlStore"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db,teardown := sqlStore.TestDb(t, dataBaseUrl)
	defer teardown("users")

    ctx:= context.TODO()
	s:= sqlStore.New(db)
	u:= model.TestUser(t)

	assert.NoError(t,s.User().Create(u,ctx))
	assert.NotNil(t,u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db,teardown := sqlStore.TestDb(t, dataBaseUrl)
	defer teardown("users")
	var ctx context.Context
	ctx= context.TODO()
	s:= sqlStore.New(db)
	email:= "test@mail.ru"
	u,err := s.User().FindByEmail(email,ctx)
	assert.NoError(t,err)
	assert.NotNil(t,u)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}