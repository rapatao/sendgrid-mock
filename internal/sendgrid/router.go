package sendgrid

import (
	_ "embed"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"net/http"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"
	"sendgrid-mock/internal/web/restrouters"
)

var (
	//go:embed json/schema.json
	schema string

	definition = gojsonschema.NewStringLoader(schema)
)

type Service struct {
	config *config.Config
	repo   *repository.Service
	event  *eventsender.Service
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

	id, err := s.persist(context.Request.Context(), bytes)
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

func validate(body []byte) error {
	current := gojsonschema.NewBytesLoader(body)

	result, err := gojsonschema.Validate(definition, current)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return errors.New("invalid JSON")
	}

	return nil
}

var (
	_ restrouters.Router = (*Service)(nil)
)
