package sendgrid

import (
	"strings"
	"testing"
)

func TestValidateReportsSchemaViolations(t *testing.T) {
	err := validate([]byte(`{"personalizations":[]}`))
	if err == nil {
		t.Fatal("expected validation to fail")
	}

	if strings.Contains(err.Error(), "invalid JSON") || err.Error() == "" {
		t.Errorf("error is not descriptive: %s", err.Error())
	}
}
