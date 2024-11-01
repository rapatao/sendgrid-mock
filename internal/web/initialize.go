package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/sendgrid"
	"sendgrid-mock/internal/web/restrouters"
)

func (c *Controller) Initialize(container *injector.Container) error {
	var (
		healthRouter   restrouters.HealthRouter
		sendgridRouter sendgrid.Service
		cfg            config.Config
	)

	err := container.Get(&cfg)
	if err != nil {
		return err
	}

	err = container.Get(&healthRouter)
	if err != nil {
		return err
	}

	err = container.Get(&sendgridRouter)
	if err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)

	c.config = &cfg
	c.engine = gin.New()

	c.engine.Use(
		gin.Recovery(),
		c.configureLogger(),
		c.configureCORS(),
	)

	c.registerControllers(&healthRouter, &sendgridRouter)

	return nil
}
