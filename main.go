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

	var action string
	fmt.Println("Permanently delete files? (y/n)\nSelecting \"no\" will move files to a new folder.")
	fmt.Scan(&action)

	fmt.Printf("\nPath: %s\nSearch: %s\nAction: %s\n\n", path, search, action)

	// If user doesn't want to delete files, create a new folder to move them to
	deleteDir := path + "/_deleted"
	if action == "n" {
		err := os.MkdirAll(deleteDir, os.ModePerm)
		check(err)
		fmt.Println("Directory created:", deleteDir)
	}

	var deletedFiles []string

	// Walk through all files under the path recursively, skipping folders
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Check if file matches the search string
		if strings.Contains(info.Name(), search) {
			fmt.Printf("File: %q\n", info.Name())

			if action == "y" {
				// Delete file
				err := os.Remove(path)
				check(err)
				deletedFiles = append(deletedFiles, path+info.Name())
			} else {
				// Move file to a new folder
				err = os.Rename(path, deleteDir+"/"+info.Name())
				check(err)
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

	// Create text file with all the filenames that were deleted
	if len(deletedFiles) > 0 {
		deleteLog := path + "/_deleted.txt"
		file, err := os.Create(deleteLog)
		check(err)
		defer file.Close()

		for _, value := range deletedFiles {
			fmt.Fprintln(file, value)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
