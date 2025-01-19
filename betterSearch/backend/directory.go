package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DirUtils struct{}

func (d *DirUtils) GetDirectoryMap(dir string) []string {
	return getDirectoryMap(dir)
}

func (d *DirUtils) openFileInOS(dir string) error {
	return openFileInOS(dir)
}

// Recursive reading of directories
func getDirectoryMap(dir string) []string {
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
func openFileInOS(path string) error {
	cmd := exec.Command("cmd", "/C", "start", "", path)
	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}

	return nil
}
