package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run main.go <input> <output>")
		return
	}
	// get input and output file
	args := os.Args[1:]
	input, output := args[0], args[1]
	// read file content into string
	content, err := os.ReadFile(input)
	if err != nil {
		println("Error reading file:", err)
		return
	}
	// split content into words based on spaces using fields (spaces gets removed as well)
	words := strings.Fields(string(content))
	result := excecuteMarker(words) // (initialize result to) excecute functions based on markers

	// write result to output file
	finalOutput := finalizeOutput(strings.Join(result, " ")) // (join the result into a string with spaces between each word)
	err = os.WriteFile(output, []byte(finalOutput), 0644)    // (write the final output to the output file)
	if err != nil {
		println("Error writing file:", err)
		return
	}
}

func excecuteMarker(words []string) []string {
	apostropheCount := 0
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		switch words[i] {
		case "(cap)":
			// capitalize the word before the marker
			words[i] = capitalize(words[i-1])
			// add the capitalized word to the result
			// creating a new slice with the capitalized word as the first element, so I don't have to loop again to reverse the result and get the correct order
			result = append([]string{words[i]}, result...)
			i-- // decrease i to skip the word that was capitalized
		case "(cap,":
			// get the amount of words to capitalize (since we want the number after the marker and not the ")" we need len)
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			// if the amount is not a number, return invalid
			if err != nil {
				return []string{"Invalid"}
			}
			// remove the amount added to result in previous loop
			result = result[1:]
			// j is the index of the word to capitalize and only traverse based on the amount
			for j := 0; j < amount; j++ {
				words[i-j] = capitalize(words[i-j-1])
				result = append([]string{words[i-j]}, result...)
			}
			// decrease i by the amount of words capitalized
			i -= amount
		case "(low)":
			words[i] = toLower(words[i-1])
			result = append([]string{words[i]}, result...)
			i--
		case "(low,":
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				return []string{"Invalid"}
			}
			result = result[1:]
			for j := 0; j < amount; j++ {
				words[i-j] = toLower(words[i-j-1])
				result = append([]string{words[i-j]}, result...)
			}
			i -= amount
		case "(up)":
			words[i] = toUppercase(words[i-1])
			result = append([]string{words[i]}, result...)
			i--
		case "(up,":
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				return []string{"Invalid"}
			}
			result = result[1:]
			for j := 0; j < amount; j++ {
				words[i-j] = toUppercase(words[i-j-1])
				result = append([]string{words[i-j]}, result...)
			}
			i -= amount
		case "(bin)":
			words[i] = binToDecimal(words[i-1])
			result = append([]string{words[i]}, result...)
			i--
		case "(hex)":
			words[i] = hexToDecimal(words[i-1])
			result = append([]string{words[i]}, result...)
			i--
		case "a":
			// Check if the next word starts with a vowel for 'an'
			if i+1 < len(words) && strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
				result = append([]string{"an"}, result...)
			} else {
				result = append([]string{"a"}, result...)
			}
		case "A":
			if i+1 < len(words) && strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
				result = append([]string{"An"}, result...)
			} else {
				result = append([]string{"A"}, result...)
			}
		case "'":
			// Handle apostrophes for possessives or contractions
			if apostropheCount == 0 && i > 0 {
				// Append apostrophe to the previous word (i-1)
				words[i-1] += "'"
				apostropheCount++
			} else if apostropheCount > 0 && i+1 < len(words) {
				// Prepend apostrophe to the next word (i+1)
				result[0] = "'" + result[0]
				apostropheCount = 0
			}
		default:
			// Handle words without any specific marker
			if strings.Contains(words[i], "'") {
				apostropheCount++
			}
			result = append([]string{words[i]}, result...)
		}
	}
	return result
}

// need to add edgecase when "..." or "!?" is found. "I was thinking ..." -> "I was thinking..." and "I was thinking !?" -> "I was thinking!?"
func finalizeOutput(input string) string {
	symbols := []string{",", ".", ":", ";", "!", "?"}

	// Clear up whitespaces
	trimWhite := regexp.MustCompile(`\s+`)

	for _, v := range symbols {
		input = strings.ReplaceAll(input, " "+v, v) // Remove space before symbol
		input = strings.ReplaceAll(input, v, v+" ") // Add space after symbol
	}

	// Handle single quotes when used like "he's" etc
	input = strings.ReplaceAll(input, "' ", "'")
	input = strings.ReplaceAll(input, " '", "'")

	input = trimWhite.ReplaceAllString(input, " ")
	return input
}

func capitalize(word string) string {
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

func binToDecimal(a string) string {
	decimalValue, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return "Invalid"
	}
	return strconv.Itoa(int(decimalValue))
}

func hexToDecimal(a string) string {
	decimalValue, err := strconv.ParseInt(a, 16, 64)
	if err != nil {
		return "Invalid"
	}
	return strconv.Itoa(int(decimalValue))
}
