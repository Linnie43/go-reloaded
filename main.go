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
	args := os.Args[1:]
	input, output := args[0], args[1]
	content, err := os.ReadFile(input) // read file content into string
	if err != nil {
		println("Error reading file:", err)
		return
	}
	words := strings.Fields(string(content)) // split content into words based on spaces using fields (spaces gets removed as well)
	result := processMarker(words)           // (initialize result to) excecute functions based on markers

	finalOutput := finalizeOutput(strings.Join(result, " ")) // since result is a slice with individual words, it gets joined into a string instead with added spaces. This string then gets formated with finalize function
	err = os.WriteFile(output, []byte(finalOutput), 0644)    // finalized version to correct output (args[i])
	if err != nil {
		println("Error writing file:", err)
		return
	}
}

func processMarker(words []string) []string {
	apostropheCount := 0
	var result []string // to store correct order
	for i := len(words) - 1; i >= 0; i-- {
		switch words[i] {
		case "(cap)":
			words[i] = capitalize(words[i-1])              // capitalize the word before the marker
			result = append([]string{words[i]}, result...) // capitalized word as first element, so I don't have to loop again to reverse the result
			i--                                            // skip the word that was capitalized
		case "(cap,":
			// get the amount of words to capitalize (since we want the number after the marker and not the ")" we need len)
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil { // if the amount is not a number, return invalid
				return []string{"Invalid"}
			}
			result = result[1:]           // since I've already added the words that need to be capitalized, I make sure to remove the previous word in it's "uncapitalized" form
			for j := 0; j < amount; j++ { // j is the index of the word to capitalize and traverse based on the amount
				words[i-j] = capitalize(words[i-j-1]) // if amount is 2. i-0, i-1, i-2. Since it starts with 0 whe capitalize words -1
				result = append([]string{words[i-j]}, result...)
			}
			i -= amount // skip the multiple words changed
		case "(low)":
			words[i] = strings.ToLower((words[i-1]))
			result = append([]string{words[i]}, result...)
			i--
		case "(low,":
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				return []string{"Invalid"}
			}
			result = result[1:]
			for j := 0; j < amount; j++ {
				words[i-j] = strings.ToLower((words[i-j-1]))
				result = append([]string{words[i-j]}, result...)
			}
			i -= amount
		case "(up)":
			words[i] = strings.ToUpper((words[i-1]))
			result = append([]string{words[i]}, result...)
			i--
		case "(up,":
			amount, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				return []string{"Invalid"}
			}
			result = result[1:]
			for j := 0; j < amount; j++ {
				words[i-j] = strings.ToUpper((words[i-j-1]))
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
			if i+1 < len(words) && strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") { // Check if the next word starts with a vowel for 'an'
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
			if apostropheCount == 0 && i > 0 {
				words[i-1] += "'" // append apostrophe to the previous word
				apostropheCount++
			} else if apostropheCount > 0 && i+1 < len(words) {
				result[0] = "'" + result[0] // prepend apostrophe to the first word
				apostropheCount = 0
			}
		default:
			if strings.Contains(words[i], "'") { // Handle words without any specific marker
				apostropheCount++
			}
			result = append([]string{words[i]}, result...)
		}
	}
	return result
}

func finalizeOutput(input string) string {
	symbols := []string{",", ".", ":", ";", "!", "?"}
	trimWhite := regexp.MustCompile(`\s+`) // to remove if there are multiple spaces

	for _, v := range symbols {
		input = strings.ReplaceAll(input, " "+v, v)         // Remove space before symbol
		input = strings.ReplaceAll(input, v, v+" ")         // Add space after symbols
		input = strings.ReplaceAll(input, ". . . ", "... ") // if I fixed my dots, special case to right format
		input = strings.ReplaceAll(input, "! ? ", "!? ")
	}
	input = strings.ReplaceAll(input, " '", "'")   // needed for "something. '" -> "something.'"
	input = trimWhite.ReplaceAllString(input, " ") // removes multiple spaces if there are any

	return strings.TrimSpace(input)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
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
