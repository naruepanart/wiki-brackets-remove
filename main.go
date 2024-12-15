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
	f, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Listing files error: %v", err)
	}

	// Guard clause for no .txt files found
	if len(f) == 0 {
		log.Println("No .txt files found.")
		return
	}

	// Compile regex
	r := regexp.MustCompile(`\[\d+\]`)

	// Process each file
	for _, p := range f {
		// Read file
		d, err := os.ReadFile(p)
		if err != nil {
			log.Printf("Failed to read file %s: %v", p, err)
			continue
		}

		// Clean data
		c := r.ReplaceAll(d, nil)

		// Guard clause for no changes in the file
		if len(c) == len(d) {
			continue
		}

		// Write cleaned data back to file
		if err := os.WriteFile(p, c, 0644); err != nil {
			log.Printf("Failed to write to file %s: %v", p, err)
			continue
		}
		fmt.Printf("Cleaned: %s\n", p)
	}
}
