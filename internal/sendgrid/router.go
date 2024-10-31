package sendgrid

import (
	"github.com/gin-gonic/gin"
	"github.com/rapatao/go-injector"
	"io"
	"net/http"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/repository"
	"sendgrid-mock/internal/rest/restrouters"
)

type Service struct {
	config  *config.Config
	storage *repository.Service
}

func (s *Service) Routes() []restrouters.Route {
	return []restrouters.Route{
		{
			Method:  http.MethodPost,
			Path:    "/v3/mail/send",
			Handler: s.HandleSend,
		},
	}
}

func (s *Service) HandleSend(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if "Bearer "+s.config.ApiKey != token {
		context.JSON(http.StatusUnauthorized,
			gin.H{
				"errors": []gin.H{
					{
						"message": "failed authentication",
						"field":   "authorization",
						"help":    "check used api-key for authentication",
					},
				},
			})

		return
	}

	bytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"errors": []gin.H{
				{
					"field":   "body",
					"message": "unable to parse body",
					"help":    err.Error(),
				},
			}},
		)
	}

	err = validate(bytes)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"errors": []gin.H{
				{
					"field":   "body",
					"message": "invalid request body",
					"help":    err.Error(),
				},
			}},
		)

		return
	}

	id, err := s.persist(context.Request.Context(), s.storage, bytes)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"errors": []gin.H{
				{
					"message": "internal failure persisting message",
					"help":    err.Error(),
				},
			}},
		)

		return
	}

	context.Header("X-Message-Id", id)
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusAccepted, gin.H{})
}

var (
	_ injector.Injectable = (*Service)(nil)
	_ restrouters.Router  = (*Service)(nil)
)
