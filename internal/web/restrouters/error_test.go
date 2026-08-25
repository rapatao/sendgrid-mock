package restrouters

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAbortWithErrorOmitsEmptyEntries(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	AbortWithError(context, http.StatusNotFound, "", "message not found", "")

	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", recorder.Code)
	}

	var body struct {
		Errors []map[string]string `json:"errors"`
	}

	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v", err)
	}

	if len(body.Errors) != 1 {
		t.Fatalf("expected one error entry, got %d", len(body.Errors))
	}

	entry := body.Errors[0]
	if entry["message"] != "message not found" {
		t.Errorf("unexpected message: %q", entry["message"])
	}

	if _, set := entry["field"]; set {
		t.Error("empty field should be omitted")
	}

	if _, set := entry["help"]; set {
		t.Error("empty help should be omitted")
	}
}
