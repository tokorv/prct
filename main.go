package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	code := handleFile(os.Args[1])

	code = colourise(code, fileType(os.Args[1]))
	lineLength := displayContents(code)

	fmt.Printf(strings.Repeat("-", lineLength+4))
	fmt.Printf("\nFilename: \t%s\n", os.Args[1])
	fmt.Printf("LineCount: \t%d\n", len(code))
	fmt.Printf("Language: \t%s\n", fileType(os.Args[1]))
	fmt.Printf(strings.Repeat("-", lineLength+4))
	fmt.Println()
}

func colourise(code []string, language string) []string {
	config := handleConfig()
	for _, l := range config.Languages {
		if language == l.Language {
			for _, keywordColour := range l.Colours.Keywords {
				fmt.Println(keywordColour.Option)
			}
		}
	}
	return code
}

func displayContents(code []string) int {
	lineLength := 0
	for lineCount, line := range code {
		fmt.Printf("%d | %s\n", lineCount+1, line)
		if len(line) > lineLength {
			lineLength = len(line)
		}
	}
	return lineLength
}

func handleFile(filename string) []string {
	f, err := os.ReadFile(filename)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("Error: do not have permission to open file \"%s\"!", filename)
		} else if os.IsNotExist(err) {
			fmt.Printf("Error: file \"%s\" does not exist!", filename)
		}
		os.Exit(1)
	}
	return strings.Split(string(f), "\n")
}

func fileType(ending string) string {
	switch strings.Split(ending, ".")[1] {
	case "py":
		return "Python"
	case "js":
		return "JavaScript"
	case "c":
		return "C"
	case "cpp":
		return "C++"
	case "rct":
		return "ReCT"
	case "json":
		return "JSON"
	case "txt":
		return "None (Text file)"
	case "java":
		return "Java"
	case "cs":
		return "C#"
	case "php":
		return "PHP"
	default:
		return "Unknown"
	}
}

func handleConfig() Config {
	file, _ := os.ReadFile("config.json")
	var config Config
	json.Unmarshal(file, &config)
	return config
}
