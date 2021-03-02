package fabricators

import (
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateUser(email string, password string) (*models.User, error) {
	o := orm.NewOrm()

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		logs.Error("Hashing password error:", err)
		return nil, err
	}
	user := &models.User{Email: email}
	user.EncryptedPassword = hashedPassword

	_, err = o.Insert(user)
	if err != nil {
		logs.Error("Creating user error: ", err)
		return nil, err
	}

	return user, nil
}
