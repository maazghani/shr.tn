// execute_fuzzy_search.go implements the fuzzy search function.
// It is called by the fuzzy search handler in fuzzy_search.go.
// It returns a list of the top 10 most similar words to the input word.
// It uses the Levenshtein distance algorithm to calculate the similarity between words.

package backend

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

// ExecuteFuzzySearch takes in a word and returns a list of the top 10 most similar words to the input word.        
func ExecuteFuzzySearch(word string) []string {
    // 
    // Open the file containing the list of words.

    file, err := os.Open("backend/words.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    // 
    // Read the file line by line and store the words in a slice.

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }
    
    // 
    // Calculate the Levenshtein distance between the input word and each word in the slice.
    // Store the Levenshtein distance and the word in a map.

    var wordMap = make(map[int]string)
    for _, w := range words {
        wordMap[LevenshteinDistance(word, w)] = w
    }

    //
    // Sort the map by the Levenshtein distance in ascending order.

    var keys []int
    for k := range wordMap {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    //
    // Store the top 10 most similar words in a slice.
    
    var similarWords []string
    for i := 0; i < 10; i++ {
        similarWords = append(similarWords, wordMap[keys[i]])
    }

    //
    // Return the slice of the top 10 most similar words.
    
    return similarWords
}

// LevenshteinDistance calculates the Levenshtein distance between two words.
// It returns the Levenshtein distance between the two words.
func LevenshteinDistance(word1 string, word2 string) int {
    // 
    // Convert the words to lowercase.

    word1 = strings.ToLower(word1)
    word2 = strings.ToLower(word2)

    //
    // Create a 2D slice to store the Levenshtein distance between each prefix of the two words.

    var distance [][]int
    for i := 0; i <= len(word1); i++ {
        var row []int
        for j := 0; j <= len(word2); j++ {
            row = append(row, 0)
        }
        distance = append(distance, row)
    }

    //
    // Calculate the Levenshtein distance between each prefix of the two words.

    for i := 0; i <= len(word1); i++ {

        