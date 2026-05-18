package manager

import (
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo   *repository.Service
	event  *eventsender.Service
	config *config.Config
}

func strOrNil(context *gin.Context, param string) *string {
	value, set := context.GetQuery(param)
	if !set || value == "" {
		return nil
	}

	return &value
}

func intOrDefault(context *gin.Context, param string, def int) int {
	str, set := context.GetQuery(param)
	if !set || str == "" {
		return def
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return value
}
