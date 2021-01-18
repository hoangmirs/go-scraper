package controllers

type MainController struct {
	baseController
}

func (c *MainController) NestPrepare() {
	if c.CurrentUser == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}
}

func (c *MainController) Get() {
	c.Layout = "layouts/application.html"
	c.TplName = "index.html"
}
