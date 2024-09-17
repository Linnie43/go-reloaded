package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Function to convert binary to decimal
func BintoDecimal(a string) string {
	decimalValue, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return "Invalid binary string"
	}
	return strconv.Itoa(int(decimalValue))
}

// Function to capitalize a word
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// Function to uppercase a word
func toUppercase(s string) string {
	return strings.ToUpper(s)
}

// Function to lowercase a word
func toLower(s string) string {
	return strings.ToLower(s)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Syntax error: go run main.go <input_file> <output_file>")
		return
	}
	inputFile := args[1]
	outputFile := args[2]

	// Read file content
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	// Convert file content to string
	content := string(fileContent)

	// Define regex pattern to find binary numbers followed by "(bin)"
	re := regexp.MustCompile(`([01]+)\s*\(bin\)`)

	// Replace all occurrences of binary numbers with their decimal equivalents
	updatedContent := re.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the binary number from the match
		submatches := re.FindStringSubmatch(match)
		binaryStr := submatches[1]

		// Convert the binary string to decimal
		decimalStr := BintoDecimal(binaryStr)

		// Return the decimal value (without the "(bin)" marker)
		return decimalStr
	})

	// Define a regex pattern for custom markers like (cap), (up), (low, x)
	customRe := regexp.MustCompile(`(\w+)\s*\((cap|up|low, \d+)\)`)

	// Replace all occurrences of custom markers with their manipulated versions
	updatedContent = customRe.ReplaceAllStringFunc(updatedContent, func(match string) string {
		// Extract the word and marker
		submatches := customRe.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		word := submatches[1]
		marker := submatches[2]

		// Process the word based on the marker
		switch {
		case marker == "cap":
			return Capitalize(word)
		case marker == "up":
			return toUppercase(word)
		case strings.HasPrefix(marker, "low"):
			var num int
			fmt.Sscanf(marker, "low, %d", &num)
			return toLower(word[:num]) + word[num:]
		default:
			return word
		}
	})

	// Write the updated content to the output file
	err = os.WriteFile(outputFile, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Printf("Content successfully written to %s\n", outputFile)
}
