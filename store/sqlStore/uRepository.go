package sqlStore

import (
	"GoTogether/model"
	"GoTogether/store"
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserRepo struct {
	store *Store
}


func (r *UserRepo) Create (u *model.User,ctx context.Context) error{

	res,err := u.Validate()
	if err != nil {
		return  err
	}

	if !res {
		return errors.New("Not valid user")
	}


	if err:=u.BeforeCreate();err!=nil{
		return  err
	}


    ctx, _= context.WithTimeout(ctx, 3 * time.Second)  //context.Background()
    sqlstring := "INSERT INTO users (email, password) VALUES ($1,$2) RETURNING id "
	return r.store.db.QueryRow(ctx,sqlstring,u.Email,u.Password).Scan(&u.Id)
}

func (r *UserRepo) FindByEmail(email string,ctx context.Context)(*model.User,error){
	u :=&model.User{
		OpenPas: "00000000000",
	}
	sqlString:= "SELECT id, email, password FROM users WHERE email = $1"
	if err:= r.store.db.QueryRow(ctx,sqlString,email).Scan(
		&u.Id,
		&u.Email,
		&u.Password);err != nil{
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	//&u.OpenPas = &u.Password
	return u, nil
}