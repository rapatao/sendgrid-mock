package manager

import (
	"bytes"
	"fmt"
	"net/http"
	"sendgrid-mock/internal/web/restrouters"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func (s *Service) handleClick(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"event_id", "missing event id", "use /messages/{event_id}/links/{encoded_link}")

		return
	}

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

	encodedLink := context.Param("link")
	encodedLink = strings.TrimPrefix(encodedLink, "/")

	if encodedLink == "" {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"link", "missing link", "the link segment must be a base64 encoded URL")

		return
	}

	link, err := decode(encodedLink)
	if err != nil {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"link", "link is not valid base64", err.Error())

		return
	}

	if link == "" {
		restrouters.AbortWithError(context, http.StatusBadRequest,
			"link", "link decodes to an empty URL", "redirecting to an empty URL would loop back to this endpoint")

		return
	}

	s.event.TriggerClick(context.Request.Context(),
		message,
		context.GetHeader("User-Agent"),
		context.ClientIP(),
		link,
	)

	context.Redirect(http.StatusTemporaryRedirect, link)
}

func htmlWrapper(eventID string, content *string) *string {
	if content == nil {
		return nil
	}

	parse, err := html.Parse(strings.NewReader(*content))
	if err != nil {
		return content
	}

	replaceLink(eventID, parse)

	buf := new(bytes.Buffer)
	err = html.Render(buf, parse)
	if err != nil {
		return content
	}

	result := buf.String()

	return &result
}

func isSelfLink(href string) bool {
	href = strings.TrimSpace(href)

	return href == "" || strings.HasPrefix(href, "#")
}

func replaceLink(eventID string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for ix, attribute := range n.Attr {
			// self-referencing hrefs would redirect back to the tracking URL itself
			if attribute.Key == "href" && !isSelfLink(attribute.Val) {
				n.Attr[ix].Val = fmt.Sprintf("/messages/%s/links/%s", eventID, encode(attribute.Val))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceLink(eventID, c)
	}
}
