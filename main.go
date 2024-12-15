package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Get .txt files
	fs, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Listing files error: %v", err)
	}

	// Guard clause for no .txt files found
	if len(fs) == 0 {
		log.Println("No .txt files found.")
		return
	}

	// Compile regex
	re := regexp.MustCompile(`\[\d+\]`)

	// Process each file
	for _, f := range fs {
		// Read file
		d, err := os.ReadFile(f)
		if err != nil {
			log.Printf("Failed to read file %s: %v", f, err)
			continue
		}

		// Clean data
		c := re.ReplaceAll(d, nil)

		// Guard clause for no changes in the file
		if len(c) == len(d) {
			continue
		}

		// Write cleaned data back to file
		if err := os.WriteFile(f, c, 0644); err != nil {
			log.Printf("Failed to write to file %s: %v", f, err)
			continue
		}

		fmt.Printf("Cleaned: %s\n", f)

	}
}
