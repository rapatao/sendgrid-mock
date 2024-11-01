package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rapatao/go-injector"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sendgrid-mock/internal/sendgrid"
	"sendgrid-mock/internal/web/restrouters"
)

type Controller struct {
	engine *gin.Engine
}

func (c *Controller) Initialize(container *injector.Container) error {
	gin.SetMode(gin.ReleaseMode)

	c.engine = gin.New()

	c.engine.Use(
		gin.Recovery(),
		c.configureLogger(),
		c.configureCORS(),
	)

	var (
		healthRouter   restrouters.HealthRouter
		sendgridRouter sendgrid.Service
	)

	err := container.Get(&healthRouter)
	if err != nil {
		return err
	}

	err = container.Get(&sendgridRouter)
	if err != nil {
		return err
	}

	c.registerControllers(&healthRouter, &sendgridRouter)

	return nil
}

func (c *Controller) Listen(address string) error {
	return c.engine.Run(address)
}

func (c *Controller) registerControllers(controllers ...restrouters.Router) {
	for _, controller := range controllers {
		for _, route := range controller.Routes() {
			log.Info().
				Str("web", "service").
				Msgf("registering %s %s", route.Method, route.Path)

			c.engine.Handle(route.Method, route.Path, route.Handler)
		}
	}

	//c.serveStaticContent(c.engine)
}

func (c *Controller) configureCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
		AllowWildcard:   true,
	})
}

func (c *Controller) configureLogger() gin.HandlerFunc {
	nonLoggablePaths := []string{"/", "/health"}
	//nonLoggablePaths = append(nonLoggablePaths, c.listStaticFiles(c.staticContentPath(), "/")...)

	return logger.SetLogger(logger.WithLogger(func(context *gin.Context, log zerolog.Logger) zerolog.Logger {
		writer := gin.DefaultWriter

		for _, uri := range nonLoggablePaths {
			if context.Request.RequestURI == uri {
				return log.Output(writer).With().Logger()
			}
		}

		return log.With().Logger()
	}))
}

var _ injector.Injectable = (*Controller)(nil)
