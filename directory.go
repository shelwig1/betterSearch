package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Recursive reading of directories
func GetDirectoryMap(dir string) []string {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking filepath: ", err)
	}

	return files
}

// Opens passed file with default OS application
func OpenFileInOS(path string) error {
	cmd := exec.Command("cmd", "/C", "start", "", path)
	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}

	return nil
}
