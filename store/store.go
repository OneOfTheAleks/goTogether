package store

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	config *Config
	db *sqlx.DB
}

func New (config *Config) *Store{ //Store
	return &Store{
		config: config,
	}
}

func (s *Store) Open () error {
	db, err := sqlx.Open("postgres", s.config.DataBaseUrl)
	if err != nil {
		return err
	}
	if err:= db.Ping();err!= nil{
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close () {
	s.db.Close()
}