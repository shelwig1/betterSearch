package backend

import (
	"fmt"
	"testing"
)

// func TestGetDirectoryMap(t *testing.T) {
// 	dirUtils := &DirUtils{}

// 	result := dirUtils.GetDirectoryMap(projectDir + "test_files")

// 	for _, v := range result {
// 		fmt.Println(v)
// 	}

// }

func TestOpenFileInOS(t *testing.T) {
	dirUtils := &DirUtils{}

	path := "C:\\Programming Projects\\betterSearch\\betterSearch\\test_files\\test.txt"

	dirUtils.OpenFileInOS(path)
}

func TestOpenFileExplorer(t *testing.T) {
	dirUtils := &DirUtils{}

	dir, err := dirUtils.ChooseDirectory()
	if err != nil {
		fmt.Println("Error retrieved: ", err)
	} else {
		fmt.Println("Directory: ", dir)
	}
}
