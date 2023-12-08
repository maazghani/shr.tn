package backend

import (
	"backend"
	"reflect"
	"testing"
)

func TestExecuteFuzzySearch_ExactMatch(t *testing.T) {
	words := []string{"tensorflow", "gpt-2", "fastai"}
	expected := []backend.FuzzySearchResult{
		{Username: "tensorflow", Repository: "", Score: 1.0},
	}
	actual, err := backend.ExecuteFuzzySearch("tensorflow", words, 1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestExecuteFuzzySearch_PartialMatch(t *testing.T) {
	words := []string{"tensorflow", "gpt-2", "fastai"}
	expected := []backend.FuzzySearchResult{
		{Username: "tensorflow", Repository: "", Score: 0.888},
	}
	actual, err := backend.ExecuteFuzzySearch("tensoflow", words, 1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestExecuteFuzzySearch_NoMatch(t *testing.T) {
	words := []string{"tensorflow", "gpt-2", "fastai"}
	expected := []backend.FuzzySearchResult{}
	actual, err := backend.ExecuteFuzzySearch("nonexistentword", words, 1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
