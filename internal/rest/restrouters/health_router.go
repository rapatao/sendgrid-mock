package restrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/rapatao/go-injector"
	"net/http"
	"sendgrid-mock/internal/repository"
)

type HealthRouter struct {
	storage repository.Service
}

func (h *HealthRouter) Routes() []Route {
	return []Route{
		{
			Method: http.MethodGet,
			Path:   "/health",
			Handler: func(context *gin.Context) {
				context.String(http.StatusOK, "ok")
			},
		},
	}
}

func (h *HealthRouter) Initialize(container *injector.Container) error {
	var db repository.Service
	err := container.Get(&db)
	if err != nil {
		return err
	}

	h.storage = db

	return nil
}

var (
	_ injector.Injectable = (*HealthRouter)(nil)
	_ Router              = (*HealthRouter)(nil)
)
