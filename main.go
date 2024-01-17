package main

import (
	"bufio"
	"fmt"
	"os"
)

func findWords(letters string, words []string) []string {
	result := make([]string, 0)

	letterFreq := make(map[rune]int)
	for _, char := range letters {
		letterFreq[char]++
	}

	for _, word := range words {
		wordFreq := make(map[rune]int)
		for _, char := range word {
			wordFreq[char]++
		}

		canConstruct := true
		for char, count := range wordFreq {
			if letterCount, ok := letterFreq[char]; !ok || count > letterCount {
				canConstruct = false
				break
			}
		}

		if canConstruct {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	// Check if at least one command-line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <letters>")
		os.Exit(1)
	}
	letters := os.Args[1]

	// Read words from words.txt file
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
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
		fmt.Println("Error reading file:", scanner.Err())
		os.Exit(1)
	}

	matchingWords := findWords(letters, words)
	fmt.Println("Matching words:", matchingWords)
}
