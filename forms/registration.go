package forms

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"
)

type RegistrationForm struct {
	Email                string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password             string `form:"password" valid:"Required; MinSize(6)"`
	PasswordConfirmation string `form:"password_confirmation" valid:"Required; MinSize(6)"`
}

func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	if registrationForm.Password != registrationForm.PasswordConfirmation {
		// Set error messages of Name by SetError and HasErrors will return true
		err := v.SetError("Password", "Password confirmation does not match")
		if err != nil {
			logs.Error("SetError error:", err)
		}
	}
}

func (registrationForm *RegistrationForm) CreateUser() error {
	valid := validation.Validation{}

	success, err := valid.Valid(registrationForm)
	if err != nil {
		logs.Error("Validation error:", err)
	}

	if !success {
		for _, err := range valid.Errors {
			return err
		}
	}

	user := models.User{}
	user.Email = registrationForm.Email

	hash, err := helpers.EncryptPassword([]byte(registrationForm.Password))
	if err != nil {
		logs.Error("Encryption error:", err)
	}

	user.EncryptedPassword = string(hash)

	o := orm.NewOrm()
	_, err = o.Insert(&user)

	return err
}
