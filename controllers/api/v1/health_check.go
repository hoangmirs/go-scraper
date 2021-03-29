package apiv1controllers

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"

	"github.com/beego/beego/v2/core/logs"
)

type HealthCheck struct {
	baseController
}

func (c *HealthCheck) Get() {
	healthCheckSerializer := v1serializers.HealthCheck{
		HealthCheck: true,
	}

	err := c.renderJSON(healthCheckSerializer.Data())
	if err != nil {
		err = c.renderGenericError(err)
		if err != nil {
			logs.Error("Error: %v", err.Error())
		}
	}
}
