package fabricators

import (
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/onsi/ginkgo"
)

func FabricateUser(email string, password string) *models.User {
	o := orm.NewOrm()

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		ginkgo.Fail("Hashing password error:" + err.Error())
	}
	user := &models.User{Email: email}
	user.EncryptedPassword = hashedPassword

	_, err = o.Insert(user)
	if err != nil {
		ginkgo.Fail("Fabricate user error:" + err.Error())
	}

	return user
}
