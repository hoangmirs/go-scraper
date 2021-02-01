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

// CreateUser insert a new User into database and returns last inserted Id on success.
func CreateUser(user *User) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(user)
}
