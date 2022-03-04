package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t * testing.T, DbUrl string) (*Store,func(... string)) {
	ctx := context.Background()
	t.Helper()
	config :=NewConfig()
	config.DataBaseUrl = DbUrl
	s:= New(config)
	if err := s.Open(ctx); err != nil{
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) >0 {
			strSQL := fmt.Sprintf("TRUNCATE %s CASCADE",strings.Join(tables,". "))
			if _,err := s.db.Exec(ctx,strSQL); err != nil{
				t.Fatal(err)
			}
		}
		s.Close()
	}

}
