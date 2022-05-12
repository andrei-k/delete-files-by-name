package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var path string
	fmt.Println("Enter the path to the folder you want to search and delete files in:")
	fmt.Scan(&path)

	var regex string
	fmt.Println("Enter the string that will be used to search and delete files:")
	fmt.Scan(&regex)

	fmt.Printf("\nPath: %s\nRegex: %s\n\n", path, regex)

	// List all files under the path recursively, skipping folders
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		fmt.Printf("File: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", path, err)
		return
	}
}
