package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rapatao/go-injector"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"path"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/web/restrouters"
)

type Controller struct {
	engine *gin.Engine
	config *config.Config
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

	c.serveStaticContent(c.engine)
}

func (c *Controller) serveStaticContent(engine *gin.Engine) {
	engine.NoRoute(func(context *gin.Context) {
		file := path.Join(c.config.WebStaticFiles, context.Request.URL.Path)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			// if the file does not exist, return a custom 404 page
			// c.JSON(404, gin.H{"error": "Not Found"})
			context.Redirect(http.StatusTemporaryRedirect, "/")
		} else {
			// if the file exists, serve it as static file
			context.File(file)
		}
	})
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
	nonLoggablePaths := []string{"/health"}
	nonLoggablePaths = append(nonLoggablePaths, c.listStaticFiles(c.config.WebStaticFiles, "/")...)

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

func (c *Controller) listStaticFiles(staticDir string, base string) []string {
	var files []string

	dir, err := os.ReadDir(staticDir)
	if err != nil {
		log.Error().Err(err).Msg("failed to load list of static files")
	}

	for _, file := range dir {
		if file.IsDir() {
			files = append(files, c.listStaticFiles(path.Join(staticDir, file.Name()), base+file.Name())...)
		} else {
			files = append(files, base+file.Name())
		}
	}

	return files
}

var _ injector.Injectable = (*Controller)(nil)
