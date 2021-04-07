package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Base

	Email             string `orm:"unique"`
	EncryptedPassword string

	Keywords []*Keyword `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// CreateUser insert a new User into database and returns last inserted Id on success.
func CreateUser(user *User) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(user)
}

// GetUserByID returns the User by passing userID
func GetUserByID(userID uint) (*User, error) {
	user := &User{Base: Base{Id: userID}}

	o := orm.NewOrm()
	err := o.Read(user)
	if err != nil {
		user = nil
	}

	return user, err
}
