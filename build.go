package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Specify the folder path
	folderPath := "sketches"

	// Get the list of folders in the specified path
	folderList, err := getFolderList(folderPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Create a slice to store the folder names with "sketch" property
	var data []map[string]string

	// Store each folder name in the slice with "sketch" property
	for _, folder := range folderList {
		data = append(data, map[string]string{"sketch": folder})
	}

	// Convert the slice to a JSON array
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Write the JSON data to db.json
	err = ioutil.WriteFile("db.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing to db.json: %v\n", err)
		return
	}

	fmt.Println("Folders have been stored in db.json.")
}

// getFolderList returns a list of folder names in the specified path
func getFolderList(path string) ([]string, error) {
	var folderList []string

	// Walk through the folder path
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a directory
		if info.IsDir() && currentPath != path {
			// Get the folder name
			folderName := filepath.Base(currentPath)
			folderList = append(folderList, folderName)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return folderList, nil
}
