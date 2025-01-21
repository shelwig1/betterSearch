package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sqweek/dialog"
)

type DirUtils struct{}

var currentDirectoryMap []string

// Recursive reading of directories

// Change it so we only pull files we can read
func (d *DirUtils) GenerateDirectoryMap(dir string) []string {
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

	currentDirectoryMap = files
	return files
}

func GetCurrentDirectoryMap() []string {
	return currentDirectoryMap
}

// Opens passed file with default OS application
func (d *DirUtils) OpenFileInOS(path string) error {
	cmd := exec.Command("cmd", "/C", "start", "", path)
	err := cmd.Start()

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}

	return nil
}

func (d *DirUtils) ChooseDirectory() (string, error) {

	dir, err := dialog.Directory().Title("Select a folder").Browse()
	if err != nil {
		return "", err
	}

	// sanitizedDir := strings.ReplaceAll(dir, "\\", "//")

	return dir, err

}
