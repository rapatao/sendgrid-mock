package manager

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func (s *Service) handleClick(context *gin.Context) {
	eventID := context.Param("event_id")
	if eventID == "" {
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	message, err := s.repo.Get(context.Request.Context(), eventID)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	if message == nil {
		context.AbortWithStatus(http.StatusNotFound)

		return
	}

	encodedLink := context.Param("link")
	if encodedLink == "" {
		context.AbortWithStatus(http.StatusBadRequest)

		return
	}

	link, err := decode(encodedLink)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
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

func replaceLink(eventID string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for ix, attribute := range n.Attr {
			if attribute.Key == "href" {
				n.Attr[ix].Val = fmt.Sprintf("/messages/%s/%s", eventID, encode(attribute.Val))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceLink(eventID, c)
	}
}
