package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourceDir := "./source/app"
	initAllFile := "./source/initall/initall.go"

	// Ensure the initall directory exists
	if err := os.MkdirAll(filepath.Dir(initAllFile), 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Open initall.go file for writing
	file, err := os.Create(initAllFile)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Write package declaration
	if _, err := file.WriteString("package initall\n\nimport (\n"); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// Walk through the source directory
	err = filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the path is a directory and contains a Go file
		if d.IsDir() && path != sourceDir {
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			for _, fileInfo := range files {
				if strings.HasSuffix(fileInfo.Name(), ".go") {
					relativePath := strings.TrimPrefix(path, sourceDir+"/")
					importPath := fmt.Sprintf("_ \"regtool/%s\"\n", relativePath)
					if _, err := file.WriteString(importPath); err != nil {
						log.Fatalf("Failed to write to file: %v", err)
					}
					break
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", sourceDir, err)
	}

	// Write closing parenthesis
	if _, err := file.WriteString(")\n"); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Println("initall.go file generated successfully.")
}
