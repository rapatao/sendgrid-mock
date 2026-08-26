package manager

import (
	"net/http"
	"sendgrid-mock/internal/web/restrouters"

	"github.com/gin-gonic/gin"
)

func (s *Service) handleSearch(context *gin.Context) {
	to := strOrNil(context, "to")
	from := strOrNil(context, "subject")
	page := intOrDefault(context, "page", 0)
	rows := intOrDefault(context, "rows", 10)

	search, err := s.repo.Search(context.Request.Context(), to, from, page, rows)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"", "unable to search messages", err.Error())

		return
	}

	context.JSON(http.StatusOK, search)
}
