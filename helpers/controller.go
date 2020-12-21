package helpers

import (
	"github.com/beego/beego/v2/server/web"
)

// SetControllerAttributes : Set some attributes of controller
func SetControllerAttributes(c *web.Controller) {
	controllerName, actionName := c.GetControllerAndAction()

	c.Data["ControllerName"] = ToSnakeCase(controllerName)
	c.Data["ActionName"] = ToSnakeCase(actionName)
}
