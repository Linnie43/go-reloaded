package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func toUppercase(word string) string {
	return strings.ToUpper(word)
}

func toLower(word string) string {
	return strings.ToLower(word)
}

func BintoDecimal(a string) string {
	decimalValue, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return "Invalid"
	}
	return strconv.Itoa(int(decimalValue))
}

func processWords(words []string, count int, transformFunc func(string) string) []string {
	if count > len(words) {
		count = len(words)
	}
	for i := len(words) - count; i < len(words); i++ {
		if i >= 0 {
			words[i] = transformFunc(words[i])
		}
	}
	return words
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Syntax error: go run main.go <input_file> <output_file>")
		return
	}
	inputFile := args[1]
	outputFile := args[2]

	// Read input file
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	// Convert file content to string
	content := string(fileContent)

	// Regex pattern to find custom markers
	customRe := regexp.MustCompile(`((?:\w+\s*)+)\s*\((cap|up|low)(?:,\s*(\d+))?\)`)

	// Replace all occurrences of custom markers with their altered versions
	updatedContent := customRe.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the words, marker, and optional number
		submatches := customRe.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		// Extract the words, marker, and optional number
		wordsStr := submatches[1]
		marker := submatches[2]
		num := 0
		if len(submatches) == 4 {
			num, _ = strconv.Atoi(submatches[3])
		}

		// Split the words into a slice
		words := strings.Fields(wordsStr)
		var transformedWords []string

		// Process the words based on the marker
		switch marker {
		case "cap":
			transformedWords = processWords(words, num, Capitalize)
		case "up":
			transformedWords = processWords(words, num, toUppercase)
		case "low":
			transformedWords = processWords(words, num, toLower)
		default:
			transformedWords = words
		}

		return strings.Join(transformedWords, " ")
	})

	// Write the updated content to the output file
	err = os.WriteFile(outputFile, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Printf("Content successfully written to %s\n", outputFile)
}
