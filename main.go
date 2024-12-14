package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Guard Clause: Ensure there's at least one .txt file to process
	files, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatalf("Error listing files: %v", err)
	}
	if len(files) == 0 {
		log.Println("No .txt files found")
		return
	}

	// Compile the regular expression pattern once, no need to recompile on each file
	re := regexp.MustCompile(`\[\d+\]`)

	// Iterate through files and process them
	for _, f := range files {
		// Read file into memory
		data, err := os.ReadFile(f)
		if err != nil {
			log.Printf("Error reading file %s: %v", f, err)
			continue
		}

		// Clean the data by replacing the pattern with nil
		cleaned := re.ReplaceAll(data, nil)

		// If the data has changed, write it back to the file
		if len(cleaned) != len(data) {
			err = os.WriteFile(f, cleaned, 0644)
			if err != nil {
				log.Printf("Error writing file %s: %v", f, err)
				continue
			}
			fmt.Printf("Cleaned: %s\n", f)
		}
	}
}
