package main

import (
	"bufio"
	"fmt"
	"os"
)

// findWords returns a slice of words that can be constructed from a given set of letters.
func findWords(letters string, words []string) []string {
	result := make([]string, 0)

	letterFreq := make(map[rune]int)
	for _, char := range letters {
		letterFreq[char]++
	}

	for _, word := range words {
		if canConstruct(letterFreq, word) {
			result = append(result, word)
		}
	}

	return result
}

// canConstruct checks if a given word can be constructed from a frequency map of letters.
func canConstruct(letterFreq map[rune]int, word string) bool {
	wordFreq := make(map[rune]int)
	for _, char := range word {
		wordFreq[char]++
	}

	for char, count := range wordFreq {
		if letterCount, ok := letterFreq[char]; !ok || count > letterCount {
			return false
		}
	}

	return true
}

func main() {
	// Check if at least one command-line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <letters>")
		os.Exit(1)
	}
	letters := os.Args[1]

	// Read words from words.txt file
	words, err := readWordsFromFile("words.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	matchingWords := findWords(letters, words)
	fmt.Println("Matching words:", matchingWords)
}

// readWordsFromFile reads words from a file and returns a slice of strings.
func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return words, nil
}
