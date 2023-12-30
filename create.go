package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get the list of archetypes
	archetypes, err := getArchetypes()
	if err != nil {
		fmt.Printf("Error getting archetypes: %v\n", err)
		return
	}

	// Prompt the user to select an archetype
	selectedArchetype, err := selectArchetype(archetypes)
	if err != nil {
		fmt.Printf("Error getting user input: %v\n", err)
		return
	}

	// Prompt the user to name their sketch
	sketchName, err := getSketchName()
	if err != nil {
		fmt.Printf("Error getting user input: %v\n", err)
		return
	}

	// Copy the selected archetype folder to "sketches" with the provided name
	err = copyFolder(filepath.Join("archetypes", selectedArchetype), filepath.Join("sketches", sketchName))
	if err != nil {
		fmt.Printf("Error copying folder: %v\n", err)
		return
	}

	fmt.Printf("Sketch created successfully: %s\n", filepath.Join("sketches", sketchName))
}

func getArchetypes() ([]string, error) {
	var archetypes []string

	// List directories in the "archetypes" folder
	files, err := os.ReadDir("archetypes")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			archetypes = append(archetypes, file.Name())
		}
	}

	return archetypes, nil
}

func selectArchetype(archetypes []string) (string, error) {
	fmt.Println("Select an archetype:")
	for i, arch := range archetypes {
		fmt.Printf("%d. %s\n", i+1, arch)
	}

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(archetypes) {
		return "", fmt.Errorf("Invalid selection")
	}

	return archetypes[choice-1], nil
}

func getSketchName() (string, error) {
	fmt.Print("Name your sketch: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimSpace(scanner.Text()), nil
}

func copyFolder(src, dest string) error {
	// Create the destination folder if it doesn't exist
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Walk through the source folder and copy files and subdirectories
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the destination path
		destPath := filepath.Join(dest, strings.TrimPrefix(path, src))

		if info.IsDir() {
			// Create the directory in the destination
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// Copy the file
		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		return err
	})
}
