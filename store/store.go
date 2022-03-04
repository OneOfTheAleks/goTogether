package store

import (
	"context"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	config *Config
	db *pgxpool.Pool
	UserRepo *UserRepo
}

func New(config *Config) *Store { //Store
	return &Store{
		config: config,
	}
}

func (s *Store) Open(ctx context.Context) error {
	/*	 db, err := sqlx.Connect("pgx", s.config.DataBaseUrl)
	if err != nil {
		return err
	}*/

	cfg, err := pgxpool.ParseConfig(s.config.DataBaseUrl)
	if err != nil {
		return err
	}
	//ctx = context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return err
	}

	err = pool.Ping(ctx)
    if err != nil{
    	return err
	}
	s.db = pool
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User () *UserRepo {
	if s.UserRepo != nil{
		return s.UserRepo
	}
	s.UserRepo = &UserRepo{
		store: s,
	}
return s.UserRepo

}