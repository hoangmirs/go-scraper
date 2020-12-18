package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Base

	Email             string `orm:"unique"`
	EncryptedPassword string
}

func init() {
	orm.RegisterModel(new(User))
}
