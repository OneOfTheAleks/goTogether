package apiServer

import (
	"GoTogether/store/sqlStore"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

func Start(config *Config) error {

	db, err := newDb(config.DataBaseUrl)
	if err !=nil {
		return err
	}
	defer db.Close()

   store:= sqlStore.New(db)
   srv:= NewServer(store)

	return http.ListenAndServe(config.BindAdrr,srv)
}

func newDb(DataBaseUrl string)(*pgxpool.Pool,error){
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(DataBaseUrl)
	if err != nil {
		return nil,err
	}

  db, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return  nil,err
	}

	if err:= db.Ping(ctx); err != nil{
		return nil,err
	}
	return db,nil
}