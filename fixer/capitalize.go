package fixer

import (
	"strings"
)

// Capitalize capitalizes the first letter of each word in the input string
func Capitalize(s string) string {
	var result strings.Builder
	words := strings.Fields(s) // Split the string into words

	for _, word := range words {
		if len(word) > 0 {
			// Capitalize the first letter and append the rest of the word
			result.WriteString(word)
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String())
}
