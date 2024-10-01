package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	args := os.Args
	//if its not 3 args, output what it needs
	if len(args) != 3 {
		log.Fatal("Usage: go run main.go [input].txt [output].txt")
		return
	}

	// read input file
	content, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal("Error reading file:", err)
		return
	}

	// split input into words
	words := strings.Fields(string(content))
	result := processMarkers(words)

	// finalize input and join to new output
	finalOutput := finalizeOutput(strings.Join(result, " "))
	err = os.WriteFile(args[2], []byte(finalOutput), 0644)
	if err != nil {
		log.Fatal("Error writing to output file:", err)
	}
}

// handles markers like (cap), (up), (low), (bin), and (hex)
func processMarkers(words []string) []string {
	result := []string{}
	for i := len(words) - 1; i >= 0; i-- { // backwards to detect special cases easier
		switch words[i] {
		case "(hex)":
			result = append(result, strconv.Itoa(hexToDecimal(words[i-1])))
			i--
		case "(bin)":
			result = append(result, strconv.Itoa(binToDecimal(words[i-1])))
			i--
		case "(up)":
			result = append(result, strings.ToUpper(words[i-1]))
			i--
		case "(up,":
			count := getNumberFromString(words[i+1])
			result = transformMultiple(strings.ToUpper, words, result, i, count)
			i -= count
		case "(low)":
			result = append(result, strings.ToLower(words[i-1]))
			i--
		case "(low,":
			count := getNumberFromString(words[i+1])
			result = transformMultiple(strings.ToLower, words, result, i, count)
			i -= count
		case "(cap)":
			result = append(result, toCap(words[i-1]))
			i--
		case "(cap,":
			count := getNumberFromString(words[i+1])
			result = transformMultiple(toCap, words, result, i, count)
			i -= count
		default:
			result = append(result, words[i])
		}
	}

	// reverse result into correct order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func finalizeOutput(input string) string {
	symbols := []string{",", ".", ":", ";", "!", "?"}

	// Clear up whitespaces
	trimWhite := regexp.MustCompile(`\s+`)
	input = trimWhite.ReplaceAllString(input, " ")
	for _, v := range symbols {
		input = strings.ReplaceAll(input, " "+v, v)
		input = strings.ReplaceAll(input, v, v+" ")
	}

	// Handle single quotes around contractions or possessives
	input = strings.ReplaceAll(input, " '", "'")
	input = strings.ReplaceAll(input, "' ", "'")

	input = trimWhite.ReplaceAllString(input, " ")
	return input
}

func transformMultiple(f func(string) string, words []string, result []string, index int, count int) []string {
	result = removeIndex(result, len(result)-1)
	for currentCount := 0; currentCount < count; currentCount++ {
		if index-1-currentCount >= 0 {
			result = append(result, f(words[(index-1)-currentCount]))
		}
	}
	return result
}

func hexToDecimal(input string) int {
	hex, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(hex)
}

func binToDecimal(input string) int {
	bin, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(bin)
}

func toCap(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}

func getNumberFromString(input string) int {
	temp := ""
	for _, v := range input {
		if unicode.IsDigit(v) {
			temp += string(v)
		}
	}
	result, err := strconv.Atoi(temp)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
