package manager

import (
	"fmt"
	"net/http"
	"sendgrid-mock/internal/web/restrouters"

	"github.com/gin-gonic/gin"
)

func (s *Service) handleGet(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"event_id", "missing event id", "use /messages/{event_id}?format=html|text|raw")

		return
	}

	format := strOrNil(context, "format")
	if format == nil {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"format", "missing format query parameter", "supported values: html, text, raw")

		return
	}

	message, err := s.repo.Get(context.Request.Context(), eventID)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"event_id", "unable to load message", err.Error())

		return
	}

	if message == nil {
		restrouters.AbortWithError(context, http.StatusNotFound,
			"event_id", "message not found", fmt.Sprintf("no message stored with event id %q", eventID))

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
	case "raw":
		content = message.Content.Html
		mime = "text/plain"
	default:
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"format", fmt.Sprintf("unsupported format %q", *format), "supported values: html, text, raw")

		return
	}

	if content == nil {
		restrouters.AbortWithError(context, http.StatusNotFound,
			"format", fmt.Sprintf("message has no %s content", *format),
			fmt.Sprintf("message %s was sent without a %s part", eventID, mime))

		return
	}

	s.event.TriggerOpen(context.Request.Context(), message, context.GetHeader("User-Agent"), context.ClientIP())

	context.Header("Content-Type", mime)
	context.String(http.StatusOK, *content)
}
