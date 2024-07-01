package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// printDirectoryContents prints the directory structure starting from the root.
// It uses spaces for indentation to represent the depth of each file/directory.
func printDirectoryContents(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the relative path from the root directory.
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Calculate the depth of the current file/directory.
		depth := len(strings.Split(relPath, string(os.PathSeparator)))

		// Print the name of the file/directory with appropriate indentation.
		// The root directory itself should not be indented.
		if relPath == "." {
			fmt.Println(info.Name())
		} else {
			fmt.Printf("%s%s\n", strings.Repeat("  ", depth), info.Name())
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path")
		return
	}

	root := os.Args[1]
	err := printDirectoryContents(root)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
