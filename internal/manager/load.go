package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) handleGet(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	format := strOrNil(context, "format")
	if format == nil {
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	message, err := s.repo.Get(context.Request.Context(), eventID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	if message == nil {
		context.AbortWithStatus(http.StatusNotFound)

		return
	}

	var (
		content *string
		mime    string
	)
	switch *format {
	case "html":
		content = htmlWrapper(message.EventID, message.Content.Html)
		mime = "text/html"
	case "text":
		content = message.Content.Text
		mime = "text/plain"
	default:
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if content == nil {
		context.AbortWithStatus(http.StatusNotFound)

		return
	}

	s.event.TriggerOpen(context.Request.Context(), message, context.GetHeader("User-Agent"), context.ClientIP())

	context.Header("Content-Type", mime)
	context.String(http.StatusOK, *content)
}
