package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) handleDelete(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	err := s.repo.Delete(context.Request.Context(), eventID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	context.Status(http.StatusNoContent)
}

func (s *Service) handleDeleteAll(context *gin.Context) {
	if !s.config.BlockDeleteAll {
		context.AbortWithStatus(http.StatusForbidden)

		return
	}

	err := s.repo.DeleteAll(context.Request.Context())
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	context.Status(http.StatusNoContent)
}
