package manager

import (
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
			Path:    "/messages/:event_id/links/*link",
			Handler: s.handleClick,
		},
		{
			Method:  http.MethodGet,
			Path:    "/messages/:event_id/attachments/:filename",
			Handler: s.handleDownloadAttachment,
		},
		{
			Method:  http.MethodGet,
			Path:    "/messages/:event_id",
			Handler: s.handleGet,
		},
	}
}

var _ restrouters.Router = (*Service)(nil)
