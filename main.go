package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"log"
)

func main() {
	// Get all .txt files in the current directory
	files, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Err listing files: %v", err)
	}

	// Compile the regular expression pattern
	re := regexp.MustCompile(`\[\d+\]`)

	for _, f := range files {
		// Read the file into memory
		data, err := os.ReadFile(f)
		if err != nil {
			log.Printf("Err reading file %s: %v", f, err)
			continue
		}

		// Clean the data by replacing the pattern
		clean := re.ReplaceAll(data, nil)

		// Write back if the content has changed
		if len(clean) != len(data) {
			err = os.WriteFile(f, clean, 0644)
			if err != nil {
				log.Printf("Err writing file %s: %v", f, err)
				continue
			}
			fmt.Printf("Cleaned: %s\n", f)
		}
	}
}
