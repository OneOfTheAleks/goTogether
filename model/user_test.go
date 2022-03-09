package model_test

import (
	"GoTogether/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t,u.BeforeCreate())
	assert.NotEmpty(t,u.Password)
}

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)
	res,err := u.Validate()  //result
	assert.NoError(t,err)
	result :=true
	assert.EqualValues(t,result,res)
}