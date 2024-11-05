package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sendgrid-mock/internal/web/restrouters"
)

func (s *Service) Routes() []restrouters.Route {
	return []restrouters.Route{
		{
			Method:  http.MethodGet,
			Path:    "/messages",
			Handler: s.handleSearch,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/messages/:event_id",
			Handler: s.handleDelete,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/messages",
			Handler: s.handleDeleteAll,
		},
		{
			Method:  http.MethodGet,
			Path:    "/messages/:event_id",
			Handler: s.handleGet,
		},
	}
}

func (s *Service) handleSearch(context *gin.Context) {
	to := strOrNil(context, "to")
	from := strOrNil(context, "subject")
	page := intOrDefault(context, "page", 0)
	rows := intOrDefault(context, "rows", 10)

	search, err := s.repo.Search(context.Request.Context(), to, from, page, rows)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	context.JSON(http.StatusOK, search)
}

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
	err := s.repo.DeleteAll(context.Request.Context())
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	context.Status(http.StatusNoContent)
}

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
		content = message.Content.Html
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

	userAgent := context.GetHeader("User-Agent")
	ip := context.ClientIP()

	s.event.TriggerOpen(context.Request.Context(), message, userAgent, ip)

	context.Header("Content-Type", mime)
	context.String(http.StatusOK, *content)
}

var _ restrouters.Router = (*Service)(nil)
