package apiv1controllers

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
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
		c.renderGenericError(err)
	}
}
