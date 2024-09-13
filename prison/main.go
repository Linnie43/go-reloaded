package main

import (
	"fmt"
	"go-reloaded/fixer"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Syntax error: go run main.go <input_file> <output_file>")
		return
	}
	inputFile := args[1]
	outputFile := args[2]

	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	content := string(fileContent)
	updatedContent := content

	for {
		index := strings.Index(updatedContent, "(bin)")
		if index == -1 {
			break
		}

		start := index - 1
		for start >= 0 && (updatedContent[start] == '0' || updatedContent[start] == '1') {
			start--
		}

		binaryStr := updatedContent[start+1 : index]
		decimalStr := fixer.BintoDecimal(binaryStr)
		updatedContent = updatedContent[:start+1] + decimalStr + updatedContent[index+5:]
	}

	err = os.WriteFile(outputFile, fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	fmt.Printf("Content successfully written to %s\n", outputFile)
}
