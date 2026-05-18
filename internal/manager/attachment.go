package manager

import (
	"encoding/base64"
	"net/http"
	"sendgrid-mock/internal/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) handleDownloadAttachment(context *gin.Context) {
	eventID := context.Param("event_id")
	filename := context.Param("filename")

	message, err := s.repo.Get(context.Request.Context(), eventID)
	if err != nil || message == nil {
		context.AbortWithStatus(http.StatusNotFound)
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
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	data, err := base64.StdEncoding.DecodeString(target.Content)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.Header("Content-Disposition", "attachment; filename="+target.Filename)
	context.Data(http.StatusOK, target.Type, data)
}
