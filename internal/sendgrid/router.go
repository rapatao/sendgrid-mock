package sendgrid

import (
	_ "embed"
	"errors"
	"io"
	"net/http"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"
	"sendgrid-mock/internal/web/restrouters"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
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
		restrouters.AbortWithError(context, http.StatusUnauthorized,
			"authorization", "failed authentication", "check used api-key for authentication")

		return
	}

	bytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"body", "unable to read body", err.Error())

		return
	}

	err = validate(bytes)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"body", "invalid request body", err.Error())

		return
	}

	id, err := s.persist(context.Request.Context(), bytes)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"", "internal failure persisting message", err.Error())

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
		reasons := make([]string, 0, len(result.Errors()))
		for _, desc := range result.Errors() {
			reasons = append(reasons, desc.String())
		}

		return errors.New(strings.Join(reasons, "; "))
	}

	return nil
}

var (
	_ restrouters.Router = (*Service)(nil)
)
