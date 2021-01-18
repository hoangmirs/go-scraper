package forms

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

func (form *SessionForm) Authenticate() (*models.User, error) {
	valid := validation.Validation{}

	success, err := valid.Valid(form)
	if err != nil {
		logs.Error("Validation error:", err)
		return nil, err
	}

	if !success {
		for _, err := range valid.Errors {
			return nil, err
		}
	}

	user := models.User{
		Email: form.Email,
	}

	o := orm.NewOrm()
	err = o.Read(&user, "Email")

	if err != nil {
		return nil, err
	}

	err = helpers.ComparePassword([]byte(user.EncryptedPassword), []byte(form.Password))

	return &user, err
}
