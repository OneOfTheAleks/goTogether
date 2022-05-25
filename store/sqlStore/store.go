package sqlStore

import (
	"GoTogether/store"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {

	db *pgxpool.Pool
	UserRepo *UserRepo
}

func New(db *pgxpool.Pool) *Store { //Store
	return &Store{
		db: db,
	}
}



func (s *Store) User () store.UserRepository {
	if s.UserRepo != nil{
		return s.UserRepo
	}
	s.UserRepo = &UserRepo{
		store: s,
	}
return s.UserRepo

}