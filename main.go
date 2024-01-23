package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	// Read words from words.txt file
	words, err := readWordsFromFile("words.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	benchmarkRandomStrings(words)
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

// randomString generates a random string of given length
func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// benchmarkRandomStrings benchmarks the findWords function with random strings of different lengths
func benchmarkRandomStrings(words []string) {
	testLengths := []int{4, 8, 12, 50, 100, 200}

	for _, length := range testLengths {
		fmt.Printf("%d chars \n", length)
		startTime := time.Now()

		randomStr := randomString(length)
		fmt.Println("Genrated string: ", randomStr)
		matchingWords := findWords(randomStr, words)
		fmt.Println("Matching words:", matchingWords)

		elapsedTime := time.Since(startTime)
		fmt.Printf("Execution time: %v\n", elapsedTime)
	}
}
