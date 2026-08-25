package manager

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sendgrid-mock/internal/model"
	"sendgrid-mock/internal/web/restrouters"

	"github.com/gin-gonic/gin"
)

func (s *Service) handleDownloadAttachment(context *gin.Context) {
	eventID := context.Param("event_id")
	filename := context.Param("filename")

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

	var target *model.Attachment

	for _, att := range message.Attachments {
		if att.Filename == filename {
			target = &att

			break
		}
	}

	if target == nil {
		restrouters.AbortWithError(context, http.StatusNotFound,
			"filename", "attachment not found",
			fmt.Sprintf("message %s has no attachment named %q", eventID, filename))

		return
	}

	data, err := base64.StdEncoding.DecodeString(target.Content)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusInternalServerError,
			"filename", "attachment content is not valid base64", err.Error())

		return
	}

	context.Header("Content-Disposition", "attachment; filename="+target.Filename)
	context.Data(http.StatusOK, target.Type, data)
}
