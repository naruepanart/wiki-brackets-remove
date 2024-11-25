package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var version = "dev"

func main() {
	fmt.Println("App version:", version)

	// Use Glob to find all .txt files in the current directory
	fs, err := filepath.Glob("*.txt")
	if err != nil {
		// Log a fatal error and terminate the program if there's an error listing files
		log.Fatalf("Listing files error: %v", err)
	}

	// Check if no .txt files were found
	if len(fs) == 0 {
		log.Println("No .txt files found.")
		return
	}

	// Compile a regular expression to match patterns like [1], [2], etc.
	re := regexp.MustCompile(`\[\d+\]`)

	// Iterate over each found .txt file
	for _, f := range fs {
		// Read the contents of the current file
		d, err := os.ReadFile(f)
		if err != nil {
			// Log an error message if the file cannot be read and continue to the next file
			log.Printf("Failed to read file %s: %v", f, err)
			continue
		}

		// Use the regular expression to remove all matching patterns from the file content
		c := re.ReplaceAll(d, nil)

		// Check if the content has not changed after cleaning
		if len(c) == len(d) {
			// Continue to the next file if there are no changes
			continue
		}

		// Write the cleaned content back to the same file
		if err := os.WriteFile(f, c, 0644); err != nil {
			// Log an error message if the file cannot be written and continue to the next file
			log.Printf("Failed to write to file %s: %v", f, err)
			continue
		}

		// Print a message indicating that the file has been cleaned
		fmt.Printf("Cleaned: %s\n", f)
	}
}
