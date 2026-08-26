package manager

import (
	"net/http"
	"sendgrid-mock/internal/web/restrouters"

	"github.com/gin-gonic/gin"
)

func (s *Service) handleDelete(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"event_id", "missing event id", "use DELETE /messages/{event_id}")

		return
	}

	err := s.repo.Delete(context.Request.Context(), eventID)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"event_id", "unable to delete message", err.Error())

		return
	}

	context.Status(http.StatusNoContent)
}

func (s *Service) handleDeleteAll(context *gin.Context) {
	if s.config.BlockDeleteAll {
		restrouters.AbortWithError(context, http.StatusForbidden,
			"", "deleting all messages is disabled", "set BLOCK_DELETE_ALL to false to enable this endpoint")

		return
	}

	err := s.repo.DeleteAll(context.Request.Context())
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"", "unable to delete messages", err.Error())

		return
	}

	context.Status(http.StatusNoContent)
}
