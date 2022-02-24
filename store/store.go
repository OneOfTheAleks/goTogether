package store

import (
	"context"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	config *Config
	//db *sqlx.DB
	db *pgxpool.Pool
}

func New(config *Config) *Store { //Store
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	/*	 db, err := sqlx.Connect("pgx", s.config.DataBaseUrl)
	if err != nil {
		return err
	}*/

	cfg, err := pgxpool.ParseConfig(s.config.DataBaseUrl)
	if err != nil {
		return err
	}
	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return err
	}

	err = pool.Ping(ctx)

	s.db = pool
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
