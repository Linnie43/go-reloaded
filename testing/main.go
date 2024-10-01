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
	args := os.Args[1:]
	if len(args) == 2 {
		// Get file content
		content, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		// Split content to words and prepare relevant variables
		words := strings.Fields(string(content))
		result := []string{}
		apostropheCount := 0

		// Loop through words backwards to catch special cases first
		for i := len(words) - 1; i >= 0; i-- {
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
				result = append(result, toTitle(strings.ToLower(words[i-1])))
				i--
			case "(cap,":
				count := getNumberFromString(words[i+1])
				result = transformMultiple(toTitle, words, result, i, count)
				i -= count
			case "a":
				if strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
					result = append(result, "an")
				} else {
					result = append(result, "a")
				}
			case "A":
				if strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH") {
					result = append(result, "An")
				} else {
					result = append(result, "A")
				}
			case "'":
				if apostropheCount == 0 {
					words[i-1] = words[i-1] + "'"
					apostropheCount++
				} else {
					result = removeIndex(result, len(result)-1)
					result = append(result, "'"+words[i+1])
					apostropheCount = 0
				}
			default:
				if strings.Contains(words[i], "'") {
					apostropheCount++
				}
				result = append(result, words[i])
			}
		}

		// Reverse result to get correct order
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}

		// Finalize output and write to file
		err = os.WriteFile(args[1], []byte(finalizeOutput(strings.Join(result, " "))), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Usage: [input].txt [output].txt")
	}
}

func finalizeOutput(input string) string {
	symbols := []string{",", ".", ":", ";", "!", "?"}

	// Clear up whitespaces
	trimWhite := regexp.MustCompile(`\s+`)
	input = trimWhite.ReplaceAllString(input, " ")
	for _, v := range symbols {
		if !strings.Contains(input, v+"'") {
			input = strings.ReplaceAll(input, " "+v, v+" ")
		}
		input = strings.TrimSuffix(input, " ")
	}
	for _, v := range symbols {
		input = strings.ReplaceAll(input, " "+v, v)
	}
	input = trimWhite.ReplaceAllString(input, " ")
	return input
}

func transformMultiple(f func(string) string, words []string, result []string, index int, count int) []string {
	result = removeIndex(result, len(result)-1)
	for currentCount := 0; currentCount < count; currentCount++ {
		if index-1-currentCount >= 0 {
			words[index-1-currentCount] = strings.ToLower(words[index-1-currentCount])
			result = append(result, f(words[(index-1)-currentCount]))
		}
	}
	return result
}

func hexToDecimal(input string) int {
	input = strings.Replace(input, "0x", "", -1)
	input = strings.Replace(input, "0X", "", -1)
	hex, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		defer log.Fatal(err)
	}
	return int(hex)
}

func binToDecimal(input string) int {
	bin, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		defer log.Fatal(err)
	}
	return int(bin)
}

func toTitle(input string) string {
	if unicode.IsLower(rune(input[0])) {
		output := ""
		for i, v := range input {
			if i == 0 {
				output += strings.ToUpper(string(v))
			} else {
				output += string(v)
			}
		}
		return output
	}
	return input
}

func getNumberFromString(input string) int {
	temp := ""
	for _, v := range input {
		if strings.ContainsAny(string(v), "0123456789") {
			temp += string(v)
		}
	}
	result, err := strconv.Atoi(temp)
	if err != nil {
		defer log.Fatal(err)
	}
	return result
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
