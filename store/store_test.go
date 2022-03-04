package store_test

import (
	"os"
	"testing"
)

var (
	dataBaseUrl string
)

func TestMain(m *testing.M) {
   dataBaseUrl = os.Getenv("dataBaseUrl")
   if dataBaseUrl == ""	{
   	dataBaseUrl = "user=web password=123 host=localhost port=5432 dbname=postgres sslmode=disable pool_max_conns=10"
	}
	os.Exit(m.Run())
}
