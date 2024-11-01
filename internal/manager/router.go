package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sendgrid-mock/internal/web/restrouters"
	"strconv"
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

var _ restrouters.Router = (*Service)(nil)
