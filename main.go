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

	// No .txt files found
	if len(fs) == 0 {
		log.Println("No .txt files")
		return
	}

	// Compile regex
	re := regexp.MustCompile(`\[\d+\]`)

	// Process each file
	for _, f := range fs {
		// Read file
		d, err := os.ReadFile(f)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}

		// Clean data
		c := re.ReplaceAll(d, nil)

		// If changed, write back
		if len(c) != len(d) {
			err = os.WriteFile(f, c, 0644)
			if err != nil {
				log.Printf("Write error: %v", err)
				continue
			}
			fmt.Printf("Cleaned: %s\n", f)
		}
	}
}
