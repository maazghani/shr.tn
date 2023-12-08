package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// FuzzySearchResult represents a matched repository with its information
type FuzzySearchResult struct {
	Username   string  `json:"username"`
	Repository string  `json:"repository"`
	Score      float64 `json:"score"`
}

func ExecuteFuzzySearch(request *http.Request) (interface{}, error) {
	// Read the words file
	words, err := ioutil.ReadFile("backend/words.txt")
	if err != nil {
		return nil, fmt.Errorf("error reading words file: %w", err)
	}

	// Extract the search term from the request body
	var searchTerm string
	err = json.NewDecoder(request.Body).Decode(&searchTerm)
	if err != nil {
		return nil, fmt.Errorf("error decoding request body: %w", err)
	}

	// Convert words and search term to lowercase for case-insensitive search
	searchTerm = strings.ToLower(searchTerm)
	wordList := strings.Split(string(words), "\n")
	for i, word := range wordList {
		wordList[i] = strings.ToLower(word)
	}

	// Find the top 10 closest matches based on Levenshtein distance
	matches := FuzzyMatch(searchTerm, wordList, 10)

	// Convert matches to FuzzySearchResult format
	var results []FuzzySearchResult
	for _, match := range matches {
		results = append(results, FuzzySearchResult{
			Username:   match.Word,
			Repository: "", // You might need to modify this based on your data structure
			Score:      match.Score,
		})
	}

	return results, nil
}

// FuzzyMatch finds the top N closest matches for a given word in a list
func FuzzyMatch(target string, words []string, n int) []FuzzyMatchResult {
	var results []FuzzyMatchResult
	for _, word := range words {
		score := fuzzywuzzy.WRatio(target, word)
		if score > 0.7 {
			results = append(results, FuzzyMatchResult{Word: word, Score: score})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) > n {
		return results[:n]
	}
	return results
}

func main() {
	// Register the function with OpenFaaS
	http.HandleFunc("/", ExecuteFuzzySearch)

	// Start the server
	http.ListenAndServe(":8081", nil)
}
