package sqlStore

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"testing"
)

func TestDb(t * testing.T, DbUrl string) ( *pgxpool.Pool,func(... string)) {
	ctx := context.Background()
	t.Helper()
	cfg, err := pgxpool.ParseConfig(DbUrl)
	if err != nil {
		t.Fatal(err)
	}
	//ctx = context.Background()
	db, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err:=db.Ping(ctx);err!=nil{
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables)>0{
			db.Exec(ctx,"TRUNCATE %s CASCADE",strings.Join(tables,", "))
		}
		db.Close()
	}

}