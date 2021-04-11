package forms

import (
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email                string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password             string `form:"password" valid:"Required; MinSize(6)"`
	PasswordConfirmation string `form:"password_confirmation" valid:"Required; MinSize(6)"`
}

func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	if registrationForm.Password != registrationForm.PasswordConfirmation {
		_ = v.SetError("Password", "Password does not match password confirmation")
	}

	user := models.User{
		Email: registrationForm.Email,
	}

	o := orm.NewOrm()
	_ = o.Read(&user, "Email")

	if user.Id != 0 {
		_ = v.SetError("Email", "Email already exists")
	}
}

func (registrationForm *RegistrationForm) CreateUser() (*models.User, error) {
	valid := validation.Validation{}

	success, err := valid.Valid(registrationForm)
	if err != nil {
		logs.Error("Validate error:", err)
	}

	if !success {
		for _, err := range valid.Errors {
			return nil, err
		}
	}

	user := &models.User{
		Email: registrationForm.Email,
	}

	hashedPassword, err := helpers.HashPassword(registrationForm.Password)
	if err != nil {
		logs.Error("Encryption error:", err)
		return nil, err
	}

	user.EncryptedPassword = string(hashedPassword)

	o := orm.NewOrm()
	_, err = o.Insert(user)

	return user, err
}
