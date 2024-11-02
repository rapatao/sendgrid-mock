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

var _ restrouters.Router = (*Service)(nil)
