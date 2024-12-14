package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"log"
)

func main() {
	// Use Glob to find all .txt files in the current directory.
	files, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	// Compile the regular expression once to avoid re-compiling in the loop.
	re := regexp.MustCompile(`\[\d+\]`)

	for _, file := range files {
		// Read file in one go to avoid repeated I/O calls
		input, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue // Log and continue on error
		}

		// Perform regex replacement
		cleaned := re.ReplaceAll(input, nil)

		// Only write to file if the content has changed
		if len(cleaned) != len(input) {
			err = os.WriteFile(file, cleaned, 0644)
			if err != nil {
				log.Printf("Error writing file %s: %v", file, err)
				continue // Log and continue on error
			}
			fmt.Printf("Cleaned: %s\n", file)
		}
	}
}
