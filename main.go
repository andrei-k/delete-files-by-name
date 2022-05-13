package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path string
	fmt.Println("Enter the path to the folder you want to search and delete files in:")
	fmt.Scan(&path)

	var search string
	fmt.Println("Enter the string that will be used to search and delete files:")
	fmt.Scan(&search)

	fmt.Printf("\nPath: %s\nSearch: %s\n\n", path, search)

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

		// Check if the file matches the search string
		if strings.Contains(info.Name(), search) {
			fmt.Printf("File: %q\n", info.Name())

			// Delete the file
			err := os.Remove(path)
			if err != nil {
				fmt.Println(err)
			}

			// TODO: Rather than permanently deleting the file, move it to a new folder called "deleted"
			// TODO: This should actually be an option; ask user if they want to permanently delete or not
			// TODO: If deleting, log the file names
			return nil
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", path, err)
		return
	}
}
