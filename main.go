package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := getTxtFiles()
	if err != nil {
		log.Fatalf("Error listing files: %v", err)
	}

	if len(files) == 0 {
		log.Println("No .txt files found.")
		return
	}

	re, err := compileRegex(`\[\d+\]`)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	for _, file := range files {
		processFile(file, re)
	}
}

func getTxtFiles() ([]string, error) {
	return filepath.Glob("*.txt")
}

func compileRegex(pattern string) (*regexp.Regexp, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func processFile(file string, re *regexp.Regexp) {
	data, err := readFile(file)
	if err != nil {
		log.Printf("Failed to read file %s: %v", file, err)
		return
	}

	cleanedData := cleanData(data, re)

	if len(cleanedData) == len(data) {
		return // No changes made
	}

	if err := writeFile(file, cleanedData); err != nil {
		log.Printf("Failed to write to file %s: %v", file, err)
		return
	}

	fmt.Printf("Cleaned: %s\n", file)
}

func readFile(file string) ([]byte, error) {
	return os.ReadFile(file)
}

func cleanData(data []byte, re *regexp.Regexp) []byte {
	return re.ReplaceAll(data, nil)
}

func writeFile(file string, data []byte) error {
	return os.WriteFile(file, data, 0644)
}
