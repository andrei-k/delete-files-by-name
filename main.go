package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path, search, action string

	fmt.Println("Enter the path to the folder you want to search and delete files in:")
	fmt.Scan(&path)

	fmt.Println("Enter the string that will be used to search and delete files:")
	fmt.Scan(&search)

	fmt.Println("Permanently delete files? (y/n)\nSelecting \"no\" will move files to a new folder.")
	fmt.Scan(&action)

	fmt.Printf("\nPath: %s\nSearch: %s\nAction: %s\n\n", path, search, action)

	deleteDir := path + "/_deleted"

	// If user doesn't want to delete files, create a new folder to move them to
	if action == "n" {
		err := os.MkdirAll(deleteDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Directory created:", deleteDir)
		}
	}

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

		// Check if file matches the search string
		if strings.Contains(info.Name(), search) {
			fmt.Printf("File: %q\n", info.Name())

			if action == "y" {
				// Delete file
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
				// TODO: Log file names

			} else {
				// Move file to a new folder
				err = os.Rename(path, deleteDir+"/"+info.Name())
				if err != nil {
					fmt.Println(err)
				}
			}

			return nil
		} else {
			fmt.Println("No match found")
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", path, err)
		return
	}
}
