package manager

import (
	"github.com/gin-gonic/gin"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"
	"strconv"
)

type Service struct {
	repo  *repository.Service
	event *eventsender.Service
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
