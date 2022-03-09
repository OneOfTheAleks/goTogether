package model

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
)

type User struct {
	Id       int
	Email    string `valid:"email,required"`
	Password string
	OpenPas  string `valid:"stringlength(8|20),required"`

}

func (u* User) Validate() (bool, error){
	result,err:=govalidator.ValidateStruct(u)
  return result,err
}

func (u *User) BeforeCreate() error {

	if len(u.OpenPas) >0 {
		enc, err := encryptString(u.OpenPas)
		if err!= nil{
			return err
		}
		u.Password = enc
	}


	return nil
}

func encryptString (pas string) (string,error)  {

	//str := base64.StdEncoding.EncodeToString([]byte(pas))
	h := md5.Sum([]byte(pas))
	//
	//h =  string(h[:])
	str := base64.StdEncoding.EncodeToString(h[:])
	return  str,nil

}