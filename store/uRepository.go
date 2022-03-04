package store

import (
	"GoTogether/model"
	"context"
	"time"
)

type UserRepo struct {
	store *Store
}


func (r * UserRepo) Create (u *model.User,ctx context.Context)(*model.User, error){
    ctx, _= context.WithTimeout(ctx, 3 * time.Second)  //context.Background()
    sqlstring := "INSERT INTO users (email, password) VALUES ($1,$2) RETURNING id "
	if err := r.store.db.QueryRow(ctx,sqlstring,u.Email,u.Password).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u,nil
}

func (r *UserRepo) FindByEmail(email string,ctx context.Context)(*model.User,error){
	u :=&model.User{}
	sqlString:= "SELECT id, email, password FROM users WHERE email = $1"
	if err:= r.store.db.QueryRow(ctx,sqlString,email).Scan(&u.Id,&u.Email,&u.Password);err != nil{
		return nil, err
	}
	return u, nil
}