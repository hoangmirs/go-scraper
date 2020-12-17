package models

import "github.com/astaxie/beego/orm"

type User struct {
	Base

	Email             string `orm:"unique"`
	EncryptedPassword string
}

func init() {
	orm.RegisterModel(new(User))
}
