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

var _ restrouters.Router = (*Service)(nil)
