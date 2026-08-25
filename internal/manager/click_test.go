package manager

import (
	"strings"
	"testing"
)

func TestHtmlWrapperKeepsSelfLinks(t *testing.T) {
	content := `<a href="">empty</a><a href="#top">anchor</a><a href="https://example.com">real</a>`

	result := htmlWrapper("event-1", &content)

	if !strings.Contains(*result, `href=""`) {
		t.Errorf("empty href was rewritten: %s", *result)
	}

	if !strings.Contains(*result, `href="#top"`) {
		t.Errorf("anchor href was rewritten: %s", *result)
	}

	if !strings.Contains(*result, `href="/messages/event-1/links/`) {
		t.Errorf("real href was not rewritten: %s", *result)
	}
}
