package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path"
)

func main() {
	input := generateRandomString(100)
	result := findWords(input)

	fmt.Printf("input: %v\n", input)
	fmt.Printf("result %d: %v\n", len(result), result)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)

	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}

func findWords(input string) (result []string) {
	// Open the file
	file, err := os.Open(path.Join("find_words_file", "words.txt"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > len(input) {
			continue
		}
		intersectRes := intersect(line, input)
		if len(line) == len(intersectRes) {
			result = append(result, intersectRes)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func intersect(in, target string) string {
	// Using map to count value occurrence
	countTarget := make(map[rune]int)
	for _, v := range target {
		countTarget[v] += 1
	}

	// Do a loop check, if value is available on map. Decrease the count value by 1
	result := []rune{}
	for _, v := range in {
		if countTarget[v] > 0 {
			result = append(result, v)
			countTarget[v] -= 1
		}
	}

	return string(result)
}
