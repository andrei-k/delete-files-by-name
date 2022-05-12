package main

import "fmt"

func main() {
	var path string
	fmt.Println("Enter the path to the folder you want to search and delete files in:")
	fmt.Scan(&path)

	var regex string
	fmt.Println("Enter the string that will be used to search and delete files:")
	fmt.Scan(&regex)

	fmt.Printf("\nPath: %s\nRegex: %s\n", path, regex)

}
