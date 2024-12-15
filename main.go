package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Guard clause: No .txt files found
	fs, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Listing files error: %v", err)
	}

	if len(fs) == 0 {
		log.Println("No .txt files found")
		return
	}

	// Compile the regex pattern once
	re := regexp.MustCompile(`\[\d+\]`)

	// Process each file
	for _, f := range fs {
		// Read file
		d, err := os.ReadFile(f)
		if err != nil {
			log.Printf("Error reading file %s: %v", f, err)
			continue
		}

		// Clean data
		c := re.ReplaceAll(d, nil)

		// Guard clause: Skip file if no change is needed
		if len(c) == len(d) {
			continue
		}

		// Write the cleaned data back to file
		err = os.WriteFile(f, c, 0644)
		if err != nil {
			log.Printf("Error writing file %s: %v", f, err)
			continue
		}

		fmt.Printf("Cleaned: %s\n", f)
	}
}
