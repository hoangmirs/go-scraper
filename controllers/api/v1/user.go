package apiv1controllers

import (
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"
)

type User struct {
	baseController
}

func (c *User) Post() {
	c.ensureAuthenticatedClient()

	registrationForm := forms.RegistrationForm{}
	err := c.ParseForm(&registrationForm)
	if err != nil {
		c.renderGenericError(err)
	}

	_, err = registrationForm.CreateUser()
	if err != nil {
		c.renderError("Validation error", err.Error(), "validation_error", http.StatusUnprocessableEntity, nil)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusCreated)
}
