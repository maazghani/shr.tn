package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/maazghani/shr.tn/pkg/backend"
)

func TestRequestrouter_FuzzySearch(t *testing.T) {
  // Mock words list
  words := []string{"tensorflow", "gpt-2", "fastai"}

  // Create a dummy request with JSON body
  searchTerm := "tensorflow"
  requestBody, err := json.Marshal(searchTerm)
  if err != nil {
    t.Fatalf("Error marshalling request body: %v", err)
  }
  req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))

  // Create a recorder to capture the response
  resp := httptest.NewRecorder()

  // Set up mock functions for backend operations
  spyOnExecuteFuzzySearch := func(searchTerm string, words []string, n int) ([]backend.FuzzySearchResult, error) {
    return []backend.FuzzySearchResult{
      {Username: "tensorflow", Repository: "", Score: 1.0},
    }, nil
  }
  spyOnManageOutboundRouting := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    // Verify headers for successful routing
    if w.Header().Get("Username") != "tensorflow" || w.Header().Get("Repository") != "" {
      t.Errorf("Unexpected headers: Username=%s, Repository=%s", w.Header().Get("Username"), w.Header().Get("Repository"))
    }
    return nil, nil
  }
  backend.ExecuteFuzzySearch = spyOnExecuteFuzzySearch
  backend.ManageOutboundRouting = spyOnManageOutboundRouting

  // Handle the request using the main function
  main.main()

  // Verify response status code
  if resp.Code != http.StatusFound {
    t.Errorf("Expected status code 302 Found, got %d", resp.Code)
  }

  // Verify headers for redirect
  expectedURL := "https://github.com/tensorflow/tensorflow"
  if resp.Header().Get("Location") != expectedURL {
    t.Errorf("Expected redirect to %s, got %s", expectedURL, resp.Header().Get("Location"))
  }

  // Verify mocks were called with expected arguments
  if spyOnExecuteFuzzySearch.CallCount() != 1 || spyOnManageOutboundRouting.CallCount() != 1 {
    t.Errorf("Expected mocks to be called once each, ExecuteFuzzySearch: %d, ManageOutboundRouting: %d", spyOnExecuteFuzzySearch.CallCount(), spyOnManageOutboundRouting.CallCount())
  }
  args := spyOnExecuteFuzzySearch.Calls()[0].Arguments
  if args.Get(0).(string) != searchTerm || args.Get(1).([]string) != words || args.Get(2).(int) != 10 {
    t.Errorf("Unexpected arguments passed to ExecuteFuzzySearch: searchTerm=%s, words=%v, n=%d", args.Get(0).(string), args.Get(1).([]string), args.Get(2).(int))
  }
}

func TestRequestrouter_ErrorHandling(t *testing.T) {
  // Test cases for different error scenarios
  tests := []struct {
    name        string
    requestBody string
    wantErr    string
  }{
    {
      name:        "Empty request body",
      requestBody: "",
      wantErr:    "Error decoding request body",
    },
    {
      name:        "Invalid JSON",
      requestBody: "{invalid_json",
      wantErr:    "Error decoding request body",
    },
    {
      name:        "ExecuteFuzzySearch error",
      requestBody: `"valid_search_term"`,
      wantErr:    "Error executing fuzzy search",
    },
    {
      name:        "ManageOutboundRouting error",
      requestBody: `"valid_search_term"`,
      wantErr:    "Error managing outbound routing",
    },
  }

  for _, tc := range tests {
    t.Run(tc.name, func(t *testing.T) {
      req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(tc.requestBody)))
      resp := httptest.NewRecorder()

      // Set up
