package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`\[\d+\]`)
	for _, file := range files {
		input, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		cleaned := re.ReplaceAll(input, nil)
		err = os.WriteFile(file, cleaned, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Cleaned: %s\n", file)
	}
}
