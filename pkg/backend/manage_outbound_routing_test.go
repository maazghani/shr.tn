package backend

import (
	"backend"
	"net/http"
	"testing"
)

func TestManageOutboundRouting_Redirect(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	req.Header.Set("Username", "tensorflow")
	req.Header.Set("Repository", "tensorflow")

	resp := http.Recorder{}

	_, err = backend.ManageOutboundRouting(&resp, req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp.Code != http.StatusFound {
		t.Errorf("Expected status code 302 Found, got %d", resp.Code)
	}

	expectedURL := "https://github.com/tensorflow/tensorflow"
	if resp.Header().Get("Location") != expectedURL {
		t.Errorf("Expected redirect to %s, got %s", expectedURL, resp.Header().Get("Location"))
	}
}

func TestManageOutboundRouting_MissingHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Missing username header
	req.Header.Set("Repository", "tensorflow")

	resp := http.Recorder{}

	_, err = backend.ManageOutboundRouting(&resp, req)
	if err == nil {
		t.Errorf("Expected error due to missing headers")
	}
}
