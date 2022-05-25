package apiServer

import (
	"GoTogether/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_UserCreate(t *testing.T) {
	rec:= httptest.NewRecorder()
	req:=httptest.NewRequest(http.MethodPost,"/users",nil)
	s:= NewServer(teststore.New())
	s.ServeHTTP(rec,req)
	assert.Equal(t,rec.Code,http.StatusOK)
}
