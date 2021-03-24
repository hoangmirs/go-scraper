package apiv1controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

type baseController struct {
	web.Controller
}

func (c *baseController) renderJSON(data interface{}) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

	err := jsonapi.MarshalPayload(c.Ctx.ResponseWriter, data)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func (c *baseController) renderGenericError(gErr error) error {
	return c.renderError("Generic error", gErr.Error(), "generic_error", http.StatusBadRequest, nil)
}

func (c *baseController) renderError(title string, detail string, code string, status int, meta *map[string]interface{}) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.SetStatus(status)

	err := jsonapi.MarshalErrors(c.Ctx.ResponseWriter, []*jsonapi.ErrorObject{{
		Title:  title,
		Detail: detail,
		Code:   code,
		Meta:   meta,
	}})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
