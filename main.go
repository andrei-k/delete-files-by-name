package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path string
	fmt.Println("Enter the path to the folder you want to search and delete files in")
	fmt.Print("=>")
	fmt.Scan(&path)

	var search string
	fmt.Println("Enter the string that will be used to search and delete files")
	fmt.Print("=>")
	fmt.Scan(&search)

	var action string
	fmt.Println("Permanently delete files? (y/n)\nSelecting \"no\" will move files to a new folder")
	fmt.Print("=>")
	fmt.Scan(&action)

	var deletedFiles []string // Holds the names of the deleted files

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
			deletedFiles = append(deletedFiles, path)
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", path, err)
		return
	}

	deleteDir := path + "/_deleted"
	deleteLog := path + "/_deleted.txt"

	// If user doesn't want to delete files, create a new folder to move them to
	if action == "n" {
		err := os.MkdirAll(deleteDir, os.ModePerm)
		check(err)
		fmt.Println("Directory created:", deleteDir)
	}

	if len(deletedFiles) > 0 {
		for _, file := range deletedFiles {
			if action == "y" {
				// Delete file
				err := os.Remove(file)
				check(err)
			} else {
				// Move file to a folder
				err = os.Rename(file, deleteDir+"/"+filepath.Base(file))
				check(err)
			}
		}

		// Create a log of deleted files
		if action == "y" {
			log, err := os.Create(deleteLog)
			check(err)
			defer log.Close()
			for _, value := range deletedFiles {
				fmt.Fprintln(log, value)
			}
		}
	}

	if len(deletedFiles) == 0 {
		fmt.Println("No files found")
	} else {
		if action == "y" {
			fmt.Printf("%d files found and deleted\n", len(deletedFiles))
			fmt.Println("Log of deleted files:", deleteLog)
		} else {
			fmt.Printf("%d files found and moved to %s\n", len(deletedFiles), deleteDir)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
