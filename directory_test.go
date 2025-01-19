package main

import (
	"fmt"
	"testing"
)

func TestGetDirectoryMap(t *testing.T) {
	result := getDirectoryMap(projectDir + "test_files")

	for _, v := range result {
		fmt.Println(v)
	}

}

func TestOpenFileInOS(t *testing.T) {
	path := "C:\\Programming Projects\\betterSearch\\betterSearch\\test_files\\test.txt"

	openFileInOS(path)
}
